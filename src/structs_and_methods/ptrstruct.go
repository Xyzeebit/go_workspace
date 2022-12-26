package main

import "fmt"

type Country struct {
	name            string
	population      uint64
	gdp             float64
	politicalSystem string
}

func (c Country) details() {
	fmt.Printf("Country name: %s\nPopulation: %d\nGDP per capita: %.2f\nPolitical system: %s\n", 
	c.name, c.population, c.gdp, c.politicalSystem);
}

func (c *Country) changePopulation(n uint64) {
	c.population = n;
}

func main() {
	c := Country {
		population: 1000000,
		name: "Comoros",
		politicalSystem: "Parliamentary Democracy",
		gdp: 500.46,
	};

	c.details();
	c.changePopulation(1100000);
	c.details();
}