package main

import (
	"fmt"
)

var print = fmt.Println


// get quotient receives a float of x and y and returns a float. If y is 0, we will pass an error because you cannot do that sh*t
func getQuotient(x float64, y float64) (answer float64, err error) {
if y != 0{
	return x/y, nil
} else {
	return 0, fmt.Errorf("You cant divide by zero silly")
}

}

func addOne(ptr *int) int {
	*ptr += 1
	return *ptr
}


func main() {
print(getQuotient(10, 3))
print(getQuotient(32, 0))

var num int
num = 23

print(addOne(&num))
}