package main

import "fmt"

func recursiveFactorial(value int)int {
	if value == 1 {
		return 1
	}
	return value * recursiveFactorial(value - 1)
}

func main(){
	result := recursiveFactorial(5)
	fmt.Println(result);
}