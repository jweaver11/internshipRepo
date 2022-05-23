package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	checkFile()
	readFile()
	writeFile()
	SetGetEnv()
}

func checkFile() {
	_, err := os.Stat("OSTestFile.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist")
	} else {
		fmt.Println("file exists")
	}
}

func readFile() {
	_, err := os.ReadFile("OSTestFile.txt")

	if err != nil {
		log.Fatal(err)
	}

}

func writeFile() {
	err := os.WriteFile("OSTestFile.txt", []byte("Writing this message using the OS package"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func SetGetEnv() {
	os.Setenv("name", "Josh")
	os.Setenv("age", "21")

	fmt.Printf("This programmer's name is %s and is age is %s", os.Getenv("name"), os.Getenv("age"))
}
