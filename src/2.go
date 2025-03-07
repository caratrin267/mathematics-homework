package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x, y int
	x = rand.Intn(10)
	y = rand.Intn(10)
	result := multiply(x, y)
	fmt.Println("The result of multiplying", x, "by", y, "is", result)
}
func multiply(a int, b int) (product int) {
	product = a * b
	return
}