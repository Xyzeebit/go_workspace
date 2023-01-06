package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const million = 1000000

	type Country struct {
		Code, Name string
		Population int
	}

	records := []Country{
		{Code: "IT", Name: "Italy", Population: 60 * million},
		{Code: "ES", Name: "Spain", Population: 46 * million},
		{Code: "JP", Name: "Japan", Population: 126 * million},
		{Code: "US", Name: "United State of America", Population: 327 * million},
	}

	file, err := os.Create("records.csv");
	if err != nil {
		fmt.Println("IO error:", err);
		return;
	}
	defer file.Close();

	w := csv.NewWriter(file);
	// w := csv.NewWriter(os.Stdout);
	defer w.Flush();

	for _, r := range records {
		if err := w.Write([]string{r.Code, r.Name, strconv.Itoa(r.Population)}); err != nil {
			fmt.Println("error:", err);
			os.Exit(1);
		}
	}
}