package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename);
	if err != nil {
		return 0, err;
	}
	h := crc32.NewIEEE();
	h.Write(bs);
	return h.Sum32(), nil;
}

func main() {

	file1, err := os.Create("test1.txt");
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer file1.Close();

	file2, err := os.Create("test2.txt");
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer file2.Close();

	h1, err := getHash("test1.txt");
	if err != nil {
		fmt.Println(err)
		return;
	}
	h2, err := getHash("test2.txt");
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println(h1, h2, h1 == h2);
}