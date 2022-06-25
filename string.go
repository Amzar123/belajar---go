/*
	type data string digunakan diantara dua tanda petik
	function yang ada di type data string 
	len("string") --> untuk mengetahui panjang string
	"string"[number] --> number dimulai dari 0

*/

package main

import "fmt"

func main() {
	var (
		firstName = "Aji"
		lastName = "Zapar"
	)

	fmt.Println(firstName);
	fmt.Println(lastName);
}