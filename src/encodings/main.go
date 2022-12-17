package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"encoding/json"
	"encoding/xml"
	"encoding/csv"
	"os"
)

func main() {
	message := "hello, go (*w3hu%#"

	encodeBase64(message)
	hexEncoding(message)
	jsonEncoding()
	xmlEncoding()
	csvEncoding()
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

func xmlEncoding() {
	fmt.Println("------------ XML Encoding example ----------------")
	type EmployeeCountry struct {
		CountryCode string `xml: "code, attr"` // XML attribute code
		CountryName string `xml: ",chardata"` // XML value
	}

	type Employee struct {
		XMLName xml.Name `xml:"employee"`
		Id string `xml:"id"`
		Name string `xml:"name"`
		Email string `xml:"email"`
		Country EmployeeCountry `xml:"country"`
	}

	// struct to xml
	fmt.Println(">>>>struct to xml….")
	emp := &Employee{
		Id:"12345",
		Name:"Michael",
		Email:"michael@email.com",
		Country: EmployeeCountry{
			CountryCode:"DE",
			CountryName: "DE",
		},
	}

	encoded, err := xml.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(encoded))

	// xml string to struct
	fmt.Println(">>>>xml to struct….")
	var newEmp Employee
	// str := `<employee><id>555</id><name>Lee</name><email>lee@email.com</email><country code="CN">China</country></employee>`
	xml.Unmarshal([]byte(encoded), &newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)
	fmt.Printf("CountryCode: %s\n", newEmp.Country.CountryCode)
	fmt.Printf("CountryName: %s\n", newEmp.Country.CountryName)
}

func csvEncoding() {
	fmt.Println("------------ CSV Encoding example ----------------")

	type Employee struct {
		Name string
		Email string
		Country string
	}

	fmt.Println("Read a csv file and load to an array of struct...")
	file, err := os.Open("demo.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var employee Employee
	var employees []Employee

	for _, item := range csvData {
		employee.Name = item[0]
		employee.Email = item[1]
		employee.Country = item[2]
		employees = append(employees, employee)
		fmt.Printf("Name: %s\nEmail: %s\nCountry: %s", 
		employee.Name, employee.Email, employee.Country)
	}
	fmt.Println("Show all employees ----------------")
	fmt.Println(employees)

	fmt.Println("Write data to a csv file ----------------")
	csvFile, err := os.Create("encode.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	for _, itemEmp := range employees {
		records := []string {itemEmp.Name, itemEmp.Email, itemEmp.Country}
		err := writer.Write(records)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
	writer.Flush()
	fmt.Println(">>>>Done")
}