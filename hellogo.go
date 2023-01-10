package main

import (
	"fmt" // this is format
	// 	"bufio"
	// 	"os"
	"strings"
) // this is how we import packages into go

var print = fmt.Println // we can instantiate functions as variables to make it easier to use them
var scan = fmt.Scanln   // this is the scanner

func main() {
	// 	print("What is your name")
	// 	value := bufio.NewReader(os.Stdin)
	// 	name, err := value.ReadString('\n')
	// 	if err != nil { // if the error has a value, print error. Otherwise, it should print the name. Nil is zero values for many types.
	// 		print("error")
	// 	} else { // has to be on the same line as the other curly brace or it will show an error
	// 		print("Your name is " + name)
	// 	}
	// 	if name == "Christian" {
	// 		print("Hey, that's my name")
	// 	} else if (string(name[0]) == "C") && (strings.Compare(name, "Christian") != 1){
	// 		print("We do not have the same name but we share the same first letter")
	// 	}else {
	// 		print("Your name is not my name")
	// 	}
	print("What is your name")
	var name string // this will make the variable name a string
	scan(&name)
	name = strings.ToLower(name)
	if name == "christian" {
		print("We have the same name")
	} else if (string(name[0]) == "c") && (name != "christian") {
		print("We have the same first letter")
	} else {
		print("We do not have the same name")
	}

}
