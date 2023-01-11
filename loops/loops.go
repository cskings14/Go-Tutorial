package main

import (
	"fmt"
)

var scan = fmt.Scanln

func main() {

	for x := 5; x >=1; x-- {
		fmt.Println(x)
	}

	fmt.Println("Guess a number from 1-10")
	for (true) { // there are no while loops. There are only for loops. We can say for true though unlike other languages.
		
		var num int
		scan(&num)
		
		if num > 5 {
			fmt.Println("Too high")
		} else if num < 5 {
			fmt.Println("Too low")
		} else {
			fmt.Println("You guessed right")
			break
		}
	}
}
 