package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func calculator(left float64, right float64, op string) (float64, error) {
	switch op {
	case "+":
		total := left + right
		return total, nil
	case "-":
		total := left - right
		return total, nil
	case "*":
		total := left * right
		return total, nil
	case "/":
		var total float64
		if right == 0 {
			return total, errors.New("error: division by zero");
		}
		total = left / right;
		return total, nil;
	default: 
		fmt.Println("Something went wrong, try again");
		var total float64;
		return total, nil;
	}
}

type Operation struct {
	left float64
	right float64
	op string
}

func run(s Operation, cal func(i float64, j float64, op string) (float64, error)) {
	fmt.Print(s.left, s.op, s.right);
	result, err := cal(s.left, s.right, s.op);
	if err != nil {
		fmt.Println(" = ", err)
	} else {
		fmt.Println(" = ", result);
	}
}

func main() {
	left := []float64 { 3, 4, 5, 6, 7, 9, 0 };
	right := []float64 { 0.3, 0, 2.1, 4, 6, 8, 12 };
	op := []string { "+", "-", "*", "/" };
	
	for _, i := range left {
		for _, j := range right {
			k := rand.Int63n(int64(len(op)));
			s := Operation { i, j, op[k]};
			run(s, calculator);
		}
	}
}