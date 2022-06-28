package main

import "fmt";

func main(){
	person := map[string]string{
		"name": "eko",	
		"address": "bandung",
	}

	fmt.Println(person["name"]);

	delete(person, "name");

	fmt.Println(person);
}