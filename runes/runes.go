package main


// a rune is simply a character's unicode
import (
	"fmt"
	"unicode/utf8"
)

var println = fmt.Println

func main() {
	rStr := "abcdefghijklmnopqrstuvwxyz"
	println("Rune Count: ", utf8.RuneCountInString(rStr)) // this prints the amount of runes (characters) in the string

	for i, runeVal := range rStr {
		fmt.Printf("%d: %#U: %c\n", i, runeVal, runeVal)
	}
}