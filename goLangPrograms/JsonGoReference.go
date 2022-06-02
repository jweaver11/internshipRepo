package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Exported fields must start with capital letters!!!!
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
} //jsonMarshal struct^^^^^^^^^^^^^^^^

type Animal struct {
	Name  string
	Order string
} //jsonUnmarshal struct^^^^^^^^^^^^^^

type Message struct {
	Name, Text string //sets 'Name' as a text string
}

func main() {
	group := ColorGroup{ //Sets 'group' as var type 'ColorGroup' and gives it values
		ID:     1,
		Name:   "Reds",
		Colors: []string{"crimson", "red", "ruby", "maroon"},
	}
	jsonMarshal(group) //Calls 'jsonMarshal' function
	//jsonMarshal Function^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	data := map[string]int{ //Sets 'data' as a map - which is a collection of unordered pairs
		"a": 1,
		"b": 2,
	}
	jsonMarshalIndent(data) //calls the 'jsonMarshalIndent' function
	//jsonMarshalIndent Function^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll", "Order": "Dasyuromophia"}
		]`) //sets 'jsonBlob' as a []byte equivalent value using the '`'
	var animals []Animal             //sets 'animals' as an 'animal' struct variable that uses reference
	jsonUnmarshal(jsonBlob, animals) //calls the 'jsonUnmarshal' function
	//jsonUnmarshal Function^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
` //the '`' states that the value is equivalent to, in this case, an untyped string
	jsonDecoder(jsonStream) //calls the 'jsonDecoder' function
	//jsonDecoder Function^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	jsonEncoder() //calls the 'jsonEncoder' function
}

func jsonMarshal(group ColorGroup) {
	fmt.Println("\njsonMarshal Function:")
	b, err := json.Marshal(group) //Sets 'b' as a byte variable for the 'group' variable. Marshal returns the JSON encoding of 'group'
	if err != nil {               //Returns 'err' as not nil if there is an error
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b) //writes the var 'b' into the default output (currently terminal)
	fmt.Println("")
}

func jsonMarshalIndent(data map[string]int) {
	fmt.Println("\njsonMarshalIndent Function:")
	b, err := json.MarshalIndent(data, "<prefix>", "<indent>") //Same as Marshal but indents the output for easier reading
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func jsonUnmarshal(jsonBlob []byte, animals []Animal) {
	fmt.Println("\njsonUnmarshal Function:")
	err := json.Unmarshal(jsonBlob, &animals) //Parses the encoded data of 'jsonBlob' and stores it in the 'animals' variable as a reference
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals) //'%v' means value, the '+' adds the field names
}

func jsonDecoder(jsonStream string) {
	fmt.Println("\njsonDecoder Function:")
	dec := json.NewDecoder(strings.NewReader(jsonStream)) //Sets 'dec' as a decoder that will read the strings from the 'jsonStream' untyped string
	for {
		var m Message                             //Sets 'm' as a variable type struct 'Message'
		if err := dec.Decode(&m); err == io.EOF { //Sets var 'err' to decode the 'm' var by reference. 'err == io.EOF' will return true when there is no more data left to read
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func jsonEncoder() {
	fmt.Println("\njsonEncoder Function:")
	enc := json.NewEncoder(os.Stdout) //Sets 'enc' as a NewEncoder that writes to the standard output (currently terminal)
	//Returns a new encoder that writes to w
	d := map[string]int{"apple": 5, "lettuce": 7} //Sets 'd' as a variable with a map (unordered pairs) variable that are string and then int
	enc.Encode(d)                                 //encodes the data of 'd' using the 'enc' variable, and writes out to the designated area.
}
