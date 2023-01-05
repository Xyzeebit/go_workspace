package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("testgo.txt");
	if err != nil {
		// handle the error here
		return;
	}
	defer file.Close();

	// get the file size
	stat, err := file.Stat();
	if err != nil {
		return;
	}

	// read the file
	bs := make([]byte, stat.Size());
	_, err = file.Read(bs);
	if err != nil {
		return;
	}

	str := string(bs);
	fmt.Println(str);
	fmt.Println();

	dir, err := os.Open(".");
	if err != nil {
		return
	}
	defer dir.Close();

	fileInfos, err := dir.ReadDir(-1);
	if err != nil {
		return;
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name());
	}

	fmt.Println();

	filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path);
		return nil;
	});
}