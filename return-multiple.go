package main

import "fmt"

func sayHello()(string, int){
	return "aji", 10;
}

func nameReturn()(firstName, middleName, lastName string){
	firstName = "Aji";
	middleName = "Muhammad";
	lastName = "Zapar";

	return
}

func main(){
	fmt.Println(sayHello());

	// menghiraukan return value
	firstName, _ := sayHello();
	fmt.Println(firstName);

	fmt.Println(nameReturn());
}