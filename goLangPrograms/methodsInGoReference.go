package main

import (
	"fmt"
)

type human struct { //Declares human as struct with name as a string and age as an int
	name string
	age  int
}

func main() {
	var person human //declares variable person of data type human

	fmt.Print("Enter your name and age: ") //prep to enter name and age

	fmt.Scanln(&person.name) //accepts input of name
	fmt.Scanln(&person.age)  //accepts input of age

	person.printPerson() //calls printPerson function using a method
} //end main

func (person human) printPerson() { //declares function printPerson, uses variable person of datatype human
	fmt.Printf("The person's name is %s and their age is %d", person.name, person.age) //print persons name and age
}
