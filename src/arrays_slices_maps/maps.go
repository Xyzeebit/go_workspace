package main

import (
	"fmt"
	"math/rand"
)

func main() {
	products := make(map[string] int);
	customers := make(map[string] int);

	products["product1"] = rand.Intn(100);
	products["product2"] = rand.Intn(100);

	customers["cust1"] = rand.Intn(100);
	customers["cust2"] = rand.Intn(100);

	fmt.Println(products["product1"]);
	fmt.Println(products["product2"]);
	fmt.Println(customers["cust1"]);
	fmt.Println(customers["cust2"]);

}