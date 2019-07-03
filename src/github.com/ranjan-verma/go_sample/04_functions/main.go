package main
import (
	"fmt"
)

func Sum (a int, b int) int {
	return a+b
}

func Subs (a, b int) int {
	return a-b
}

func Greeting (name string) string {
	return "Hello " + name
}

// Multiple returns
func SumSub (a int, b int)(int, int) {
	return (a+b), (a-b)
}

func main () {
	fmt.Println(Sum(10, 30)) // Output: 40
	fmt.Println(Subs(20, 30)) // Output: -10
	fmt.Println(Greeting("Ranjan")) // Output: Hello Ranjan
	fmt.Println(SumSub(10,5)) // Output: 15, 5
}