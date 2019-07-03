package main
import (
	"fmt"
)
	
func main() {  
      number := 10  
      squareNum := func() (int){  
      number *= number  
      return number  
   }  
   fmt.Println(squareNum())  // Output: 100
   fmt.Println(squareNum())  // Output: 10000
}  