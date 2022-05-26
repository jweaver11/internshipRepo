package main

import (
	"fmt"
)

type human struct { //Declares human as struct with name as a string and age as an int
	name string
	age  int
}

type speak interface {
	intro()
	favColor()
}

func main() {
	var person human //declares variable person of data type human
	color := "blue"

	fmt.Print("Enter your name and age: ") //prep to enter name and age

	fmt.Scanln(&person.name) //accepts input of name
	fmt.Scanln(&person.age)  //accepts input of age

	person.intro()         //calls intro function using a method for person
	person.favColor(color) //calls favColor function using a method for person
} //end main

func (person human) intro() { //declares function intro, uses variable person of datatype human to declare this as a usable method
	fmt.Printf("The person's name is %s and their age is %d . ", person.name, person.age) //print persons name and age
}

func (person human) favColor(color string) {
	fmt.Printf("%s's favorit color is %s", person.name, color)
}
