package main

import (
	"fmt"
	"math"
)

func main() {
	mathCeil()
	mathFloor()
	mathMax()
	mathMin()
	mathMod()
}

func mathCeil() {
	c := math.Ceil(1.49)    //Ceil returns the lowest integer greater than the value given
	fmt.Printf("%.1f\n", c) //Decimal point with only one decimal. '%f' means decimal
}

func mathFloor() {
	f := math.Floor(1.51)   //Floor returns the highest integer lower than the value
	fmt.Printf("%.1f\n", f) //Decimal point with only one decimal. '%f' means decimal
}

func mathMax() {
	higher := math.Max(3.69, 5.53) //Returns the higher value (float) of the two given values
	fmt.Printf("%.1f\n", higher)   //Prints the variable 'higher' to one decimal
}

func mathMin() {
	lower := math.Min(4.38, 6.45) //Returns the lower value (float) of the two given values
	fmt.Printf("%.1f\n", lower)   //Prints the variable 'lower' to one decimal
}

func mathMod() {
	mod1 := math.Mod(7, 3) //Returns the remainder of value 1 (7) divided by value 2 (3) = 1.0
	mod2 := math.Mod(7, 4) //Returns the remainder of value 1 (7) divided by value 2 (4) = 3.0

	fmt.Printf("%.1f", mod1)
	fmt.Printf("%.1f", mod2)
}
