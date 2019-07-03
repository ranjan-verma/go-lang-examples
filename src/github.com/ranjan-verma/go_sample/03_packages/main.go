package main

import (
	"fmt"
	"math"
	
	"github.com/ranjanv/go_sample/03_packages/strutils"
)

func main () {
	
	fmt.Println(math.Floor(2.7)) // Output: 2
	// fmt.Println(math.Ceil(2.7)) // Output: 3
	// fmt.Println(math.Sqrt(16)) // Output: 4
	fmt.Println(strutils.Reverse("Floor")) // Output: rlooF // User defined Package
	
}