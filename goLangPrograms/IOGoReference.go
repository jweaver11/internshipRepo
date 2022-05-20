package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Writer interface { //interface that wraps the basic write method
	write(p []byte) (n int, err error)
	//Writes len(p) bytes from p tot eh underlying data stream.
	//This returns the number of bytes written as int 'n' (0 <= n <= len(p)), and any errors that stopped writing early
	//Write must return a non-nil error value
}

type Reader interface { //interface that wraps the basic read method
	Read(p []byte) (n int, err error)
	//Reader will read up to len(p) bytes into p
	//This returns the number of bytes read as int 'n' (0 <= n <= len(p))
	//Returns successful error after reading 'n' bytes, sometimes as 'non-nil' value
}

type ReadWriter interface {
	io.Reader
	io.Writer
}

func main() {
	var err error
	file, err := os.Create("IOTestFile.txt")

	if err != nil { //Error handling
		log.Fatal(err)
	}

	writeFile(file, err)
	readFile(file, err)
	readWriteFile(file, err)

	defer file.Close()

}

func writeFile(file *os.File, err error) {
	fmt.Println("\n************START OF THE 'writeFile' FUNCTION*************") //visibility message

	len, err := file.WriteString("This is a message written into this file") //Writes string text into the file, returns the length as len,
	//And the finished writing error value as err

	fmt.Printf("File Name: %s", file.Name()) //Writes the File name to the console
	fmt.Printf("\nLength: %d bytes\n", len)  //Writes the file length of bytes/characters as len to the console

	fmt.Print("************END OF THE 'writeFile' FUNCTION**********\n\n") //visibility message

	file.Close() //Closes the file
}

func readFile(file *os.File, err error) {
	fmt.Println("**********START OF THE 'readFile' FUNCTION**********") //visibility message

	fmt.Printf("Reading from file %s", file.Name()) //writes message to console with file name

	data, err := os.ReadFile("IOTestFile.txt") //Reads the file name 'IOTestFile.txt', returns length of file in characters/bytes as data
	//Also returns error value as err

	if err != nil { //prints a panic message and crashes the function is there is an error reading the file
		log.Panicf("failed reading data from file: %s", err)
	}

	fmt.Printf("\nSize: %d characters", len(data)) //prints length of data to console
	fmt.Printf("\nFile data: '%s'", data)          //prints data to console

	fmt.Println("\n***********END OF THE 'readFile' FUNCTION************") //visibility message

	file.Close() //closes the file
}

func readWriteFile(file *os.File, err error) {

}
