package main


// in Go, array sizes do not change
import (
	"fmt"
)

var print = fmt.Println


func main() {
	// var arr1 [5]int // this instantiates a new array called arr1 with the size of 5
	arr2 := [5]int{1, 2, 3, 4, 5}
	
	for i := 0; i < len(arr2); i++ {
		print(arr2[i])
	}

	// i is the index and v is the value
	for i, v := range arr2 {
		print(i, v)
	}

	arr3 := [2][2]int {
		{1,2},
		{3,4},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			print(arr3[i][j])
		}
	}

}