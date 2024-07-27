package main

import (
	"encoding/base64"
	"log"
	"os"
)

func main() {
	var token string
	token, ok := os.LookupEnv("ENCODE_TARGET")
	if !ok {
		log.Fatal("ENCODE_TARGET is not set")
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(token))

	f, err := os.Create("encoded.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write([]byte(encoded))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Token encoded and saved to encoded_token.txt")
}
