package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
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

	// asymmetric crypto
	generateRSAKeys()
	plainText := "Hello world, go!"
	fmt.Println("plainText:")
	fmt.Println(plainText)
	rsa_encrypted := encrypt_asymmetric_crypto(plainText)
	fmt.Println("encrypted:")
	fmt.Println(rsa_encrypted)
	rsa_decrypted := decrypt_asymmetric_crypto(rsa_encrypted)
	fmt.Println("decrypted:")
	fmt.Println(rsa_decrypted)
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

func generateRSAKeys() {
	fmt.Println("Generating RSA keys...")

	pubKeyFile := "pub_rsa.key"
	privKeyFile := "priv_rsa.key"

	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	// extract private and public keys from RSA keys
	privASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	// store the private and public keys into files
	privBytes := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privASN1,
	})
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	ioutil.WriteFile(privKeyFile, privBytes, 0644)
	ioutil.WriteFile(pubKeyFile, pubBytes, 0644)

	fmt.Println("Done")
}

func encrypt_asymmetric_crypto(message string) string {
	fmt.Println("--------- Encrypt asymmetric_crypto ---------------")

	pubKeyFile := "pub_rsa.key"

	pubBytes, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		panic(err)
	}
	pubBlock, _ := pem.Decode(pubBytes)
	if pubBlock == nil {
		fmt.Println("Failed to load public key file")
		return ""	
	}
	
	// Decode the RSA public key
	publicKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		fmt.Printf("bad public key %s", err)
		return ""
	}

	// encrypt message
	msg := []byte(message)
	encryptedMsg, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), msg)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(encryptedMsg)
}

func decrypt_asymmetric_crypto(message string) string {
	fmt.Println("--------- Decrypt asymmetric_crypto ---------------")

	privKeyFile := "priv_rsa.key"

	privBytes, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		panic(err)
	}

	privBlock, _ := pem.Decode(privBytes)
	if privBlock == nil {
		fmt.Println("Failed to load private key file")
		return ""
	}

	// Decode the RSA private key
	privateKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		fmt.Printf("Bad private key %s", err)
		return ""
	}

	// decrypt message
	encrypted, _ := base64.StdEncoding.DecodeString(message)
	decryptedMsg, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
	if err != nil {
		panic(err)
	}

	return string(decryptedMsg)
}