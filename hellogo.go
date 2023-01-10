package main

import (
	"fmt" // this is format

	"strings" // there is a seperate strings package
) // this is how we import packages into go

var print = fmt.Println // we can instantiate functions as variables to make it easier to use them
var scan = fmt.Scanln   // this is the scanner

func main() {
	print("What is your name")
	var name string // this will make the variable name a string
	n, err := scan(&name)
	if err != nil || n == 0 { // sees if there is an error or if there are no read in values 
		print("Invalid name")
	}else {
	name = strings.ToLower(name) // makes the name lowercase before conditionals
	if name == "christian" {
		print("We have the same name")
	} else if (string(name[0]) == "c") && (name != "christian") { // we have to cast name[0] because the value is a byte. A string is really just an array of bytes...
		print("We have the same first letter")
	} else {
		print("We do not have the same name")
	}

}

}
