package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"encoding/json"
	// "encoding/xml"
	// "encoding/csv"
	// "os"
)

func main() {
	message := "hello, go (*w3hu%#"

	encodeBase64(message)
	hexEncoding(message)
	jsonEncoding()
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

func jsonEncoding() {
	fmt.Println("------------ JSON Encoding example ----------------")
	type Employee struct {
		Id int `json: "id"`
		Name string `json: "name"`
		Email string `json: "email"`
	}

	fmt.Println("------- Struct to JSON---------")
	emp := &Employee{ Id: 12345, Name: "Michael", Email: "michael@email.com" }
	b, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("------ JSON string to struct ------------")
	var newEmp Employee
	// str := `{"Id":"4566","Name":"Brown","Email":"brown@email.com"}`
	json.Unmarshal([]byte(b), &newEmp)
	fmt.Printf("Id: %d\n",newEmp.Id)
	fmt.Printf("Name: %s\n",newEmp.Name)
	fmt.Printf("Email: %s\n",newEmp.Email)
}