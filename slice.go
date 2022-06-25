/*
	tipe data slice adalah potongan dari array
	slice mirip array, yang membedakan adalah ukuran dari slice bisa berubah
	slice dan array selalu terkoneksi

	tipe data slice ada 3 
	1. pointer --> data pertama pad slice
	2. length --> panjang data yang dapat ditampung
	3. capacity --> panjang max array dikurangi index pada pointer

	length tidak boleh lebih dari capacity
	
	function dalam slice
	len(slice) --> panjang slice
	cap(slice) --> capacity dari slice
	append(slice, data) --> membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas penuh maka akan membuat array baru
	make([]TypeData, length, capacity) --> membuat slice baru
	copy(destination, source) --> menyalin slice dari source ke destination

*/

package main 

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[10:11]
	// var slice2 = months[6:9]

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}