// function-demo
package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a+b
}

func plusplus(a, b, c int) int {
	return a+b+c
}


func main() {
	
	res := plus(1,2)
	fmt.Println("res: ", res)
	
	res = plusplus(12,23,4)
	fmt.Println("res: ", res)
	
}
