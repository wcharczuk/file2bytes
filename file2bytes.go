package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func usage() {
	fmt.Println("USAGE:")
	fmt.Println("> file2bytes <filename>")
}

func fileBytes(filePath string) ([]byte, error) {
	var err error
	if f, err := os.Open(filePath); err == nil {
		defer f.Close()
		return ioutil.ReadAll(f)
	}
	return nil, err
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	file := os.Args[1]
	fmt.Printf("Reading %s\n", file)
	fileBytes, err := fileBytes(file)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	if len(fileBytes) == 0 {
		fmt.Printf("File was empty.\n")
		os.Exit(1)
	}

	var fileByteStr []string
	for _, b := range fileBytes {
		fileByteStr = append(fileByteStr, fmt.Sprintf("0x%X", b))
	}

	fmt.Printf("[]byte{%s}\n", strings.Join(fileByteStr, ", "))
}
