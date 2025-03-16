
package main
import "fmt"
func main() {
	const pi = 3.14 // Define constant
	var radius int = 5 // define variable
	area := pi * radius * radius // calculate area
	perimeter := 2 * pi * radius // calculate perimeter
	fmt.Println("The area of the circle is:", area)
	fmt.Println("The perimeter of the circle is:", perimeter)
}
