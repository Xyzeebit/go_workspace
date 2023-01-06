package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("records.csv");
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close();

	r := csv.NewReader(file);
	r.Comma = ',';
	r.FieldsPerRecord = -1;
	records, err := r.ReadAll();
	if err != nil {
		log.Fatal(err);
	}
	type Country struct {
		Code, Name string
		Population int
	}
	var c []Country;

	for _, v := range records {
		n, _ := strconv.Atoi(v[2]);
		c = append(c, Country {Code: v[0], Name: v[1], Population: n });
		
	}
	for i, v := range c {
		fmt.Println(i, ":", v.Code, v.Name, v.Population);
	}
	
}