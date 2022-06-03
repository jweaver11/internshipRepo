package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	//Sets 'tpl' as an untyped string, very useful for html
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	t, err := template.New("webpage").Parse(tpl) //sets 't' as a template, that will return 'err' as an error. Parses 'tpl', the untyped string
	check(err)                                   //calls 'check' (check error) function

	data := struct { //Sets data as a struct with 'Title' as a string and 'Items' as an array of strings
		Title string
		Items []string
	}{
		Title: "My page", //Defines 'Title'
		Items: []string{ //Defines 'Items'
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data) //Executes the template with standard output (terminal), using the 'data' struct
	check(err)                       //calls 'check' (check error) function

	noItems := struct { //Declares a new struct 'noItems'
		Title string
		Items []string
	}{
		Title: "My other page",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems) //Executes the template with standard output (teminal), using the 'noItems' struct
	check(err)                          //calls 'check' (check error) function

	template.HTMLEscape(os.Stdout, []byte(tpl)) //Replaces the '<' and '>' characters
	//Escape means to replace one character(s) with others
	//Primarily used on '<' '>' '&' '"'
	//==
	fmt.Println(template.HTMLEscapeString(tpl))
}

func check(err error) { //Defines 'check' function to catch errors
	if err != nil {
		log.Fatal(err)
	}
}
