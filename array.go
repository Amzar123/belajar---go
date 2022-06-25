package main

import "fmt"

func main() {
	var names [10]string

	names[0] = "Aji Muhammad Zapar";
	names[1] = "Aji Muhammad Zapar";
	names[2] = "Aji Muhammad Zapar";

	fmt.Println(names)

	// membuat array secara langsung
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)

	// fungsi yang ada di array
	// mengembalikan panjang si array bukan mengembalikan jumlah data yang mengisi array
	fmt.Println(len(names))
}