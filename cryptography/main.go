package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hash_md5()
	hash_sha256()
	hash_hmac("secrt", "Hello world, go!")
}

func hash_md5() {
	fmt.Println("------------ Hashing with md5 --------------")

	message := "Hello world, go!"
	fmt.Println("Plaintext")
	fmt.Println(message)

	h := md5.New()
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)
	fmt.Println()
}

func hash_sha256() {
	fmt.Println("------------ Hashing with sha256 --------------")

	message := "Hello world, go!"
	fmt.Println("Plaintext")
	fmt.Println(message)

	h := sha256.New()
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)
	fmt.Println()
}

func hash_hmac(key, message string) {
	fmt.Println("------------ Hashing with hmac --------------")

	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("plaintext:")
	fmt.Println(message)

	hmacKey := []byte(key)
	h := hmac.New(sha256.New, hmacKey)
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)
	fmt.Println()
}