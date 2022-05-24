package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string

	boolArgPtr := flag.Bool("arg1", false, "This is a bool argument")
	//defines 'boolArgPtr':: names it as 'arg1', value is false, usage as a string

	numbPtr := flag.Int("numb", 42, "an int as a flag") //defines 'numbPtr' as a flag integer with value 42

	wordPtr := flag.String("word", "dank", "a string as a flag") //defines 'wordPtr' as a flag string as the work 'dank'

	flag.StringVar(&str, "str", "dope", "a string var") //uses existing declared variable instead of always creating a new one

	flag.Parse() //Parse command line into defined flags
	//Must be called after all flags are defined but before they are used by program

	fmt.Println("Bool arg:", *boolArgPtr)
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("str:", str)
	fmt.Println("tail: ", flag.Args())

}
