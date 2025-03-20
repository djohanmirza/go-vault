package main

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"log"
)

func main() {
	//auth the vault
	config := vault.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken("ini_token")

	//write
	secretData := map[string]interface{}{
		"password":      "Keren-bro",
		"another_token": "gak-keren",
		"one_token":     2,
		"id_user":       "string",
		"token_array": []interface{}{
			"token1",
			"token2",
		},
	}

	_, err = client.KVv2("secret").Put(context.Background(), "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	fmt.Println("Secret written successfully.")

	//read
	//secret, err := client.KVv2("secret").Get(context.Background(), "my-secret-password")
	//if err != nil {
	//	log.Fatalf("unable to read secret: %v", err)
	//}
	//
	//value, ok := secret.Data["password"].(string)
	//if !ok {
	//	log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	//}
	//
	//if value != "Keren-bro" {
	//	log.Fatalf("unexpected password value %q retrieved from vault", value)
	//}
	//
	//fmt.Println("Access granted!")

}
