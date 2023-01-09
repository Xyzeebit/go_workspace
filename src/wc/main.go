package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin));
}

// A counter function that counts the words read by the scanner
//
// Returns the number of words counted
func count(r io.Reader) int {
	// A Scanner is use to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r);

	// Define the Scanner split type to words (default is split by line)
	scanner.Split(bufio.ScanWords);

	// Define a counter
	wc := 0;

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++;
	}

	// Return the total
	return wc;
}