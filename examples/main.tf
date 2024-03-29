// Sample key for testing
resource "zitadeltools_jwt" "jwt" {
  key = "{\"type\":\"serviceaccount\",\"keyId\":\"250142936681996292\",\"key\":\"-----BEGIN RSA PRIVATE KEY-----\\nMIICXQIBAAKBgQCLF60USA3gHMI+uiL+6F2Mwylke2pPeLffW9w3svZb69uYNzkA\\nG2YlK8WDnQPQKDXcIU/ryOH7Lhz4pOl9FEuTVxBiph6sAOyGqtokUjX+kpa6kWjY\\na75e0pID5bnNZB8T50HxJS0kM06tVbuFL342+nalzzDph6t75UQGjLy9hQIDAQAB\\nAoGANTC6qXTciYW16pL36w8lJz1CyQBBKV76wdm5Hzu65nANq+oCTD4uE6znfWtA\\nGAwofKPokjxh3tJFlfhKqeBm/cGOEwZg1AV7EDjJueP7Sh9b0qlYboTSGf8XYo9X\\nKbpzFNWHZk2xigrk6NSpdzzpD2XlHcgeB+SJZABvdFOLbgkCQQDZ4titlohIBW8m\\nWw3zS91L5V5FVQkJnj54QAvrGUBXVrfFdo9zhXY4Hl77oiRW404K11GOQohyGUAy\\nk3jKMABnAkEAo2xfqDSEqARZhrlInAbvLUJD/qjSs1NlSiV/YJKGiXRh2JLGDVft\\nGnoJ0N9MeWbAK3Rj1CBepbpZ1UL7wwBvMwJBAI6d1jghTQjfTbasaQA0SyCPfNoi\\n4+yAwOETAvoaqCvC3j0I8rKpzAzFjGRm6CRbWkzsTTyxvf/5GTVBpBGrw8cCQQCh\\npn9waJKXh8Xup7QU8h7/y75qViAk1ecpUrIOmqGNsZtfmL9jT4fvWqv++gIxS9vm\\nn/hnSaWVlSFq/BkAqJtVAkA8IuF4SjVwNZBmPrQHCrKts2s9JhumVIs/MOBvYtZ6\\ngCieidjAg3Pmja1N3Jd7TpJYk8+r1ojsLfEnxIfeDs7f\\n-----END RSA PRIVATE KEY-----\\n\",\"userId\":\"250142936463892484\"}"
  audience = "https://zitadel.dev.socrate.local"
}

output "zitadeltools_jwt_jwt" {
  value = zitadeltools_jwt.jwt.jwt
  sensitive = true
}