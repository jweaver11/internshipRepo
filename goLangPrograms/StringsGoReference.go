package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"foo", "bar", "bat"}

	strContain()
	strCount()
	strHasPrefix()
	strHasSuffix()
	strIndex()
	strJoin(s)
	strRepeat()
	strReplace()
	strSplit()
	strToLower()
	strToUpper()
}

func strContain() {
	fmt.Println(strings.Contains("seafood", "foo")) //Checks if 'foo' is within string 'seafood', which it is - returns true
	fmt.Println(strings.Contains("seafood", "bar")) //Checks if 'bar' is within string 'seafood', whit it isn't - returns false
}

func strCount() {
	fmt.Println(strings.Count("cheese", "e")) //Checks how many 'e' there are in string - returns 3
	fmt.Println(strings.Count("five", ""))    //Checks how many characters there are - returns 5
}

func strHasPrefix() {
	fmt.Println(strings.HasPrefix("Gopher", "Go")) //Checks if string 'Gopher' starts with 'Go' - returns true
	fmt.Println(strings.HasPrefix("Gopher", "c"))  //Checks if string 'Gopher' starts with 'c' - returns false
	fmt.Println(strings.HasPrefix("Gopher", ""))   //Always return true
}

func strHasSuffix() {
	fmt.Println(strings.HasSuffix("Amigo", "go")) //Checks if string 'Amigo' has suffix 'go' - returns true
	fmt.Println(strings.HasSuffix("Amigo", "O"))  //Checks if string 'Amigo' has suffix 'O' - returns false
}

func strIndex() {
	fmt.Println(strings.Index("chicken", "ken")) //Checks where 'ken' appears in string 'chicken' - returns 4 (k starts in position 4)
	fmt.Println(strings.Index("chicken", "dmr")) //Checks where "dmr" appears in string 'chicken' - returns -1 (it doesn't appear)
}

func strJoin(s []string) {
	fmt.Println(strings.Join(s, ", ")) //Adds the elements of 's', and uses the separating ', ' between elements - prints foo, bar, bat
}

func strRepeat() {
	fmt.Println("ba" + strings.Repeat("na", 2)) //Prints 'ba', and adds 'na' twice - prints banana
}

func strReplace() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //Takes string 'oink oink oink' and replaces the first two 'k' with 'ky'
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //Takes string 'oink oink oink' and replaces all 'oink' with 'moo'
}

func strSplit() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //Splits string 'a,b,c' into characters at all positions where ',' is present - prints ["a" "b" "c"]
	//Uses %q (character)
	fmt.Printf("%q\n", strings.Split(" xyz ", "")) //Splits every character in string ' xyz ' with a space at all positions - prints [" " "x" "y" "z"]
}

func strToLower() {
	fmt.Println(strings.ToLower("GoPhEr")) //Prints string 'GoPhEr' as all lowercase - 'gopher'
}

func strToUpper() {
	fmt.Println(strings.ToUpper("GoPhEr")) //Prints string 'GoPhEr' as all uppercase - 'GOPHER'
}
