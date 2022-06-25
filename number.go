/*
	tipe data integer di go
	bersifat case sensitive (tidak bisa hanya menuliskan int)
	tipe data	| nilai minimum 		| nilai maksimum 
	int8		| -128					| 127
	int16		| -32768				| 32767
	int32		| -2147483648			| 2147483647
	int64		| -9223372036854775808	| 9223372036854775907

	ada lagi yang namanya uint ==> unsigned int 
	tipe data	| nilai minimum 		| nilai maksimum 
	uint8		| 0						| 255
	uint16		| 0						| 65535
	uint32		| 0						| 429467295
	uint64		| 0						| 18446744073709551615

	alias
	tipe data		| alias
	byte			| uint8
	rune			| int32
	int 			| minimal int32 (tergantung arsitektur OS, kalo 32 byte bakal pake 32)
	uint			| minimal uint32 (tergantung arsitektur OS, kalo 32 byte bakal pake 32)
*/

package main

import "fmt"

func main(){
	var aji int32 = 32;
	fmt.Println(aji);
}