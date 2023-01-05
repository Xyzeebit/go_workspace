package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "hello";
	str2 := "world";
	str3 := str1 + str2;

	fmt.Println(str3)

	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)

	// .. string to numeric

	str_val1 := "5"
	str_val2 := "5.5232"

	var err error
	var int_val1 int64
	int_val1, err = strconv.ParseInt(str_val1, 10, 64);
	if err == nil {
		fmt.Println(int_val1)
	} else {
		fmt.Println(err)
	}

	var float_val2 float64
	float_val2, err = strconv.ParseFloat(str_val2, 64);
	if err == nil {
		fmt.Println(float_val2)
	} else {
		fmt.Println(err)
	}

	// ... num to string
	num1 := 9
	num2 := 12.012

	var str_num1 string
	str_num1 = fmt.Sprintf("%d", num1);
	var str_num2 string
	str_num2 = fmt.Sprintf("%f", num2);

	fmt.Printf("%s %s\n", str_num1, str_num2)

	fmt.Println("----demo String Parser----")
	data := "Berlin;Amsterdam;London;Tokyo";
	fmt.Println(data)
	cities := strings.Split(data, ";")
	for _, city := range cities {
		fmt.Println(city)
	}

	fmt.Printf("len=%d", len("welcome to go"))

	fmt.Println("----demo copy data----")
	sample := "hello world, go!"
	fmt.Println(sample)
	fmt.Println(sample[0:len(sample)]) // copy all
	fmt.Println(sample[:5]) // copy 5 characters
	fmt.Println(sample[3:8]) // copy characters from index 3 until 8
	fmt.Println(sample[len(sample)-5:len(sample)]) // 5 copy characters
	fmt.Println(strings.ToUpper(sample))
	fmt.Println(strings.ToLower(sample))
	fmt.Println(strings.Contains(sample, "go"))
	fmt.Println(strings.Count(sample, "o"))
	fmt.Println(strings.Index(sample, "go"))
	fmt.Println(strings.HasPrefix(sample, "hello"))
	fmt.Println(strings.Replace(sample, "ll", "ke", 1))
	fmt.Println(strings.Split(sample, " "))
	arr := []byte(sample)
	str := string([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', ',', ' ', 'g', 'o', '!'});
	fmt.Println(arr, ":", str)

}