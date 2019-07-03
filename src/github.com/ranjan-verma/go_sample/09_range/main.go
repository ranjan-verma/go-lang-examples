package main
import (
	"fmt"
)

func main () {
	
	ids :=[]int{22,54,78,12,99,13}
	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	
	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	
	// Range with Maps
	emails := map[string]string{"Bob":"bob@gmail.com", "John": "john@gmail.com", "Raj":"raj@gmail.com", "Sam":"sam@gmail.com"}
	for k, v:= range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k:= range emails {
		fmt.Printf("Name: %s\n", k)
	}
}