package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileContent, err := ioutil.ReadFile("TestFile.txt")
	if err != nil {
		fmt.Println("Read file failed.")
		return
	}
	contentString := string(fileContent)
	fmt.Println(contentString)

	osFile, err := os.Open("TestFile.txt")
	if err != nil {
		return
	}
	defer osFile.Close() // Close file connection when the function is finished.
	stat, err := osFile.Stat()
	if err != nil {
		return
	}

	fmt.Printf("The file size is %d\n", stat.Size())
	byteContent := make([]byte, stat.Size())
	_, osErr := osFile.Read(byteContent)
	if osErr != nil {
		return
	}
	readAsString := string(byteContent)
	fmt.Println(readAsString)
}
