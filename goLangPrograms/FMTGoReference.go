//Go lang program
package main //Sets the program as a main package

import ( //imports the fmt library for all the fmt commands
	"fmt"
)

func main() { //Start of main function
	num := 10                        //Declares variable num, and lets the compiler auto decide its type (in this case integer)
	var str string = "I am a string" //Declares str as a string and gives its value

	var scanNum int //Sets scanNum as an integer and scanStr as a string
	var scanStr string

	var str1 string = "String1 and " //Declares variables str1 and str2 as strings
	var str2 string = "String2"
	myStr := fmt.Sprintf("%s %s", str1, str2) //Declares and Defines myStr as an auto variable (string) that is str1 and str2 added together

	printLine(num, str)        //Calls printLine function
	scanLine(scanNum, scanStr) //Calls scanLine function
	printFunction(num, str)    //Calls printFunction
	SPrintFunction(myStr)      //Calls SPrintFunction
} //end of main function

func printLine(num int, str string) {
	fmt.Print("\n", num, "\n", str, "\n") //uses command 'fmt.Print' to print everything in parentheseis, including '\n'
} //end of printLine function

func scanLine(scanNum int, scanStr string) {
	fmt.Println("Enter a number, then enter a word") //uses command 'fmt.Println' to print the line of code in parentheses. Does not include '\n'
	fmt.Scanln(&scanNum)                             //uses Scanln for user input of value scanNum and scanStr. == cin or getline
	fmt.Scanln(&scanStr)                             //^
	fmt.Print("\nNumber:", scanNum, "\nWord: ", scanStr, "\n")
} //end of scanLine function

func printFunction(num int, str string) {
	fmt.Print("\n Printing using 'Printf'")
	fmt.Printf("\n %d \n %s \n", num, str) //Uses the fmt.Printf command to print the verb types, and then defines the verbs after the quotations
} //end of printFunction

func SPrintFunction(myStr string) {
	fmt.Print("\n Printing your number and word using 'Sprintf'\n")
	fmt.Print("\n", myStr)
} //end of SPrintFunction
