package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("records.csv");
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(strings.NewReader(string(bs)));
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
		n, _ := strconv.ParseInt(v[2], 0, 64);
		c = append(c, Country {Code: v[0], Name: v[1], Population: int(n) });
		
	}
	for i, v := range c {
		fmt.Println(i, ":", v.Code, v.Name);
	}
	
}