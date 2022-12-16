package main

import (
	"fmt"
	"encoding/base64"
	// "encoding/hex"
	// "encoding/json"
	// "encoding/xml"
	// "encoding/csv"
	// "os"
)

func main() {
	message := "hello, go (*w3hu%#"
	encodeBase64(message)
}

func encodeBase64(message string) {
	fmt.Println("------------ Base 64 Encoding example ----------------")
	fmt.Println("Plaintext")
	fmt.Println(message)

	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("Encoding message:")
	fmt.Println(encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoding message:")
	fmt.Println(string(decoded))
}