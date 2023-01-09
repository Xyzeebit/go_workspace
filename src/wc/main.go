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
	// Parsing the flags provided by the user
	flag.Parse();

	fmt.Println(count(os.Stdin, *lines));
}

// A counter function that counts the words read by the scanner
//
// Returns the number of words counted
func count(r io.Reader, countLines bool) int {
	// A Scanner is use to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r);

	// If the count lines flag is not set, we want to count words so we define
	// the Scanner split type to words (default is split by line)
	if !countLines {
		scanner.Split(bufio.ScanWords);
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