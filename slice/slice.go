package main


import (
	"fmt"
)

var print = fmt.Println

// a slice is just an array with a size that can be changed
func main() {

sl1 := make([]string, 1)

sl1[0] = "We"
// we can append to slice like this
sl1 = append(sl1, "are")

for _, v := range sl1 {
	print(v)
}

}