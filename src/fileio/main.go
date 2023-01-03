package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// write data into a file
	fmt.Println("writing data into a file")
	writeFile("welcome to go")
	readFile()
	
	// read data from a file
	fmt.Println("reading data from a file")
	readFile()

	fileReader("main.go")
}

func writeFile(message string) {
	bytes := []byte(message);
	ioutil.WriteFile("testgo.txt", bytes, 0644);
	fmt.Println("created a file")
}

func readFile() {
	data, _ := ioutil.ReadFile("testgo.txt")
	fmt.Println("file content:")
	fmt.Println(string(data))
}

func process(data []byte) {
	fmt.Println(string(data))
}

func fileReader(fileName string) error {
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close();
	data := make([]byte, 100);
	
	for {
		count, err := file.Read(data);
		if err != nil {
			return err;
		}
		if count == 0 {
			return nil;
		}
		process(data[:count]);
	}
}