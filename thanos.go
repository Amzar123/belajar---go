package main

import "fmt"

func thanosPlanet(days int )int {
	value := 1
	for i := 1; i <= days; i++ {
		fmt.Println("hari ke ",i,value)
		if i%3 == 0 {
			// thanos muncul
			value = value/2;
		} else {
			// Dr. strange muncul
			value = value*3;
		}
		
	}

	return value
}

func main() {
	result := thanosPlanet(6);
	fmt.Println("hasilnya" ,result);
}