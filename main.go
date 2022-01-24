package main

import "fmt"

func main() {
	arr := [5]int{4,3,6,7,1}
	var temp int;

	//sebelum sorting
	fmt.Println("Sebelum Sorting")
	for i := 0; i<len(arr);i++{
		fmt.Print(arr[i]," ")
	}
	fmt.Print("\n")

	//sorting dengan menggunakan sequencial sort
	for i := 0; i<len(arr);i++{
		for j := i; j<len(arr);j++{
			if (arr[i] > arr[j]){
				temp = arr[i];
				arr[i] = arr[j];
				arr[j] = temp;
			}
		}
	}
	//Menampilkan hasil sorting
	fmt.Println("Setelah Sorting")
	for i := 0; i<len(arr);i++{
		fmt.Print(arr[i]," ")
	}
	fmt.Print("\n")
}