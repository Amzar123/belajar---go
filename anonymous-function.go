package main

import "fmt"

func main(){
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	if blacklist("Aji") == true {
		fmt.Println("Betull")
	}
}