package main


import(
	"fmt" // this is format
	"bufio"
	"os" 
) // this is how we import packages into go

var print = fmt.Println // we can instantiate functions as variables to make it easier to use them


func main() {
	print("What is your name")
	value := bufio.NewReader(os.Stdin)
	name, err := value.ReadString('\n')
	if err != nil { // if the error has a value, print error. Otherwise, it should print the name
		print("error")
	} else { // has to be on the same line as the other curly brace or it will show an error
		print("Your name is " + name)
	}
}
