package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines");

	// Defining a boolean flag -b to count the bytes instead of words
	bytes := flag.Bool("b", false, "Count bytes");

	// Parsing the flags provided by the user
	flag.Parse();

	fmt.Println(count(os.Stdin, *lines, *bytes));
}

// A counter function that counts the words read by the scanner
//
// Returns the number of words counted
func count(r io.Reader, countLines bool, countBytes bool) int {
	// A Scanner is use to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r);

	// If the count lines flag is not set or the count bytes flag is not set,
	// we want to count words so we define
	// the Scanner split type to words (default is split by line)
	if !countLines {
		if !countBytes {
			scanner.Split(bufio.ScanWords);
		} else {
			scanner.Split(bufio.ScanBytes);
		}
	}

	// Define a counter
	wc := 0;

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++;
	}

	// Return the total
	return wc;
}