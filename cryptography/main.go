package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	hash_md5()
	hash_sha256()
	hash_hmac("secrt", "Hello world, go!")

	// symmetric crypto
	key := "this is key 1234"
	message := "Hello world, go!"
	encrypted := encrypt_symmetric_crypto(key,message)
	fmt.Println("message:")
	fmt.Println(message)
	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("encrypted:")
	fmt.Println(encrypted)
	fmt.Println()
	decrypted := decrypt_symmetric_crypto(key,encrypted)
	fmt.Println("encrypted message:")
	fmt.Println(encrypted)
	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("decrypted:")
	fmt.Println(decrypted)
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

func encrypt_symmetric_crypto(key, message string) string {
	fmt.Println("-------- Encrypt with symmetric_crypto ------------")

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		fmt.Println("Key must be 16, 24 or 32 byte length")
		return ""
	}
	bc, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	text := []byte(message)

	cipherText := make([]byte, aes.BlockSize + len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(bc, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], text)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func decrypt_symmetric_crypto(key, message string) string {
	fmt.Println("-------- Decrypt with symmetric_crypto ------------")

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		fmt.Println("Key must be 16, 24 or 32 byte length")
		return ""
	}

	encrypted, _ := base64.StdEncoding.DecodeString(message)
	bc, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	if len(encrypted) < aes.BlockSize {
		panic("cipher text too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(bc, iv)
	cfb.XORKeyStream(encrypted, encrypted)

	return string(encrypted)
}