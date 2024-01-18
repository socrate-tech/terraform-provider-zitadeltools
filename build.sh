go build .
rm -rf examples/.terraform
rm -rf examples/terraform.tfstate
rm -rf examples/terraform.tfstate.backup
rm -rf examples/.terraform.lock.hcl
mkdir -p $HOME/.terraform.d/plugins/registry.socrate.fr/socrate-tech/zitadeltools/1.0.0/darwin_arm64/
cp zitadel-tools-provider $HOME/.terraform.d/plugins/registry.socrate.fr/socrate-tech/zitadeltools/1.0.0/darwin_arm64/terraform-provider-zitadeltools_v1.0.0
cd $PWD/examples
# export TF_LOG=info
terraform init && terraform apply -auto-approve
