package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	// "encoding/json"
	// "encoding/xml"
	// "encoding/csv"
	// "os"
)

func main() {
	message := "hello, go (*w3hu%#"

	encodeBase64(message)
	hexEncoding(message)
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

func hexEncoding(message string) {
	fmt.Println("------------ Hex Encoding example ----------------")
	fmt.Println("Plaintext")
	fmt.Println(message)

	encoded := hex.EncodeToString([]byte(message))
	fmt.Println("Encoding message:")
	fmt.Println(encoded)

	decoded, _ := hex.DecodeString(encoded)
	fmt.Println("Decoding message:")
	fmt.Println(string(decoded))
}