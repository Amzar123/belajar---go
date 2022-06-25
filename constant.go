/*
	constant nilainya tetap dan tidak bisa di ubah lagi
	dalam constant pendeklarasiannya harus langsung dengan valure nya 
	constant jika dibuat dan tidak digunakan juga nggak apa apa, tidak akan kena error  
*/

package main

import "fmt"

func main() {
	// deklarasi multiple constant 
	const(
		firstName = "Aji"
		lastName = "Zapar"	
	)

	fmt.Println(firstName)
}