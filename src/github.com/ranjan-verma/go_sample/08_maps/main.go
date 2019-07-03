package main
import (
	"fmt"
)

func main () {
	
	// Define map
	//emails := make(map[string]string)
	
	//Assign key values
	//emails["Bob"] = "bob@gmail.com"
	//emails["John"] = "john@gmail.com"
	//emails["Raj"] = "raj@gmail.com"
	//emails["Sam"] = "sam@gmail.com"
	
	emails := map[string]string{"Bob":"bob@gmail.com", "John": "john@gmail.com", "Raj":"raj@gmail.com", "Sam":"sam@gmail.com"}
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["John"])
	fmt.Println(emails["john"])
	
}