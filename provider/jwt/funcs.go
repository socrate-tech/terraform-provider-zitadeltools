package jwt

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zitadel/oidc/pkg/client"
	"github.com/zitadel/oidc/pkg/oidc"
)

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started create")

	key := d.Get("key").(string)

	keyId, err := getKeyId([]byte(key))

	if err != nil {
		return diag.Errorf("error getting keyId: %v", err.Error())
	}

	hashedKeyId := hash(keyId)

	d.SetId(hashedKeyId)

	return get(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started delete")

	d.SetId("")
	return nil
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started update")

	return get(ctx, d, m)
}

func get(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Info(ctx, "started get")

	key := d.Get("key").(string)
	tflog.Info(ctx, fmt.Sprintf("key: %s", key))
	audience := d.Get("audience").(string)

	jwt, err := key2JWT(key, audience)

	if err != nil {
		return diag.Errorf("error generating jwt: %v", err.Error())
	}

	d.Set("jwt", jwt)

	return nil
}

func getKeyId(data []byte) (string, error) {
	keyData := new(struct {
		KeyId string `json:"keyId"` // serviceaccount or application
	})
	err := json.Unmarshal(data, keyData)
	if err != nil {
		return "", err
	}
	return keyData.KeyId, nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

func key2JWT(key string, audience string) (string, error) {
	if key == "" || audience == "" {
		log.Println("Please provide at least an audience and key param:")
		return "", fmt.Errorf("Please provide at least an audience and key param")
	}

	jwt, err := generateJWTFromJSON([]byte(key), audience)

	if err != nil {
		log.Fatalf("error generating jwt: %v", err.Error())
		return "", err
	}

	return jwt, nil
}

func generateJWTFromJSON(key []byte, audience string) (string, error) {
	keyType, err := getType(key)
	if err != nil {
		return "", err
	}
	switch keyType {
	case "application":
		keyData, err := client.ConfigFromKeyFileData(key)
		if err != nil {
			return "", err
		}
		signer, err := client.NewSignerFromPrivateKeyByte([]byte(keyData.Key), keyData.KeyID)
		if err != nil {
			return "", err
		}
		return client.SignedJWTProfileAssertion(keyData.ClientID, []string{audience}, time.Hour, signer)
	case "serviceaccount":
		jwta, err := oidc.NewJWTProfileAssertionFromFileData(key, []string{audience})
		if err != nil {
			return "", err
		}
		return oidc.GenerateJWTProfileToken(jwta)
	default:
		return "", fmt.Errorf("unsupported key type")
	}
}

func getType(data []byte) (string, error) {
	keyData := new(struct {
		Type string `json:"type"` // serviceaccount or application
	})
	err := json.Unmarshal(data, keyData)
	if err != nil {
		return "", err
	}
	return keyData.Type, nil
}
