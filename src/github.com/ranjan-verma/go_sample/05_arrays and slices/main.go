package main
import (
	"fmt"
)

func main () {
	var  fruitArr[2] string 
	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Grape"
	fmt.Println(fruitArr) // Output: [Apple Grape]
	fmt.Println(fruitArr[0]) // Output: [Apple]
	
	// Declare and assign
	fruitArr1 := [2]string{"Banana", "Orange"}
	fmt.Println(fruitArr1) // Output: [Banana Orange]
	fmt.Println(fruitArr1[0]) // Output: [Banana]
	
	//fruitArr2 := [2]string{"Banana", "Orange", "Chiku"}
	//fmt.Println(fruitArr2) // Output: Index 2 out of Bound exception
	
	fruitArr3 := []string{"Banana", "Orange", "Chiku"}
	fmt.Println(len(fruitArr3)) // Output: 3
	
	fruitArr4 := []string{"Banana", "Orange", "Chiku", "Grape"}
	fmt.Println(len(fruitArr4[1:3])) // Output: [Orange Chiku]

}