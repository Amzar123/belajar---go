package main 

import "fmt"

func greetings(firstName string, lastName string){
	fmt.Println("hello", firstName, lastName)
}

func main(){
	greetings("Aji", "zapar");
}