package main


import (
	"fmt"
	"strings"
)

var print = fmt.Println

func main() {


	print("eee")
	phrase := "A description" // the original phrase
	replacer := strings.NewReplacer("A", "Another") // this is a replacer variable which changes any "A" to another
	nphrase := replacer.Replace(phrase) // This is the new phrase after using the replacer on the original phrase.
	print(nphrase)
	print(len(phrase)) // len is buiilt into Golang without the need for an import
	print(strings.Contains(nphrase, "Another"))

}