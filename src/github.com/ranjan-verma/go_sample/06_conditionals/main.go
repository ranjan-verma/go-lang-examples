package main
import (
	"fmt"
)

func main () {
	a := 10 
	b := 20
	
	// If else
	if a < b {
		fmt.Printf("%d is less than %d\n", a, b)
	} else {
		fmt.Printf("%d is greater than %d\n", a, b)
	}
	
	// If Else If
	color := "b"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}
	
	// Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}