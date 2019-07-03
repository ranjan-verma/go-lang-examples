package main
import "fmt"

name := "Abhay" // Error : It should be declared inside the function

func main () {
	//  string
	// bool
	// int: int, int8, int16, int32, int64
	// uint: uint8, uint16, uint32, uint64, uintptr
	// byte: aliase for uint8
	// rune: aliase for int32
	// float32, float64
	
	// using vars
	
	// String
	// var name string = "Ranjan"
	// fmt.Println("Hello " + name + "...") // Output: Hello Ranjan...
	
	// Int
	// var age = 37
	// fmt.Println(age) // Output: 37
	// fmt.Printf("%T\n", age) // Output: int
	
	// var age1 int32 = 40
	// fmt.Println(age1) // Output: 40
	// fmt.Printf("%T\n", age1) // Output: int32
	
	// Const
	// const isCool = true
	// fmt.Println(isCool) // Output: true
	// fmt.Printf("%T\n", isCool) // Output: bool
	// isCool = false // error as its a const
	
	// Shorthand
	// name := "Abhay" // Should in declared inside the function
	// name, email := "Abhay", "xyz@gamil.com"
	
}