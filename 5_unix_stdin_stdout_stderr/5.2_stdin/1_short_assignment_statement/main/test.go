package main

import(
	"fmt"
)

//uncomment statement(baris) dengan comment menggunakan // (double slash) untuk melakukan percobaan
func main()  {
	/*PERCOBAAN 1
	ini akan mencetak sebuah variable baru bernama "m" dengan tipe data "int" dengan nilai 123
	*/
	// m:=123
	// fmt.Println(m)
	/* HASIL : 
		$ go run test.go
		123
	*/

	/*PERCOBAAN 2
	namun jika kita mencoba untuk memasukkan nilai baru kepada variable yang di deklarasikan menggunakan Short Assignment Statement terhadap variable Short Assignment Statement yang sudah kita buat sebelumnya (m), maka akan menghasilkan error
	*/
	// m:=123
	// m:=456
	// fmt.Println(m)
	/* HASIL :
		$ go run test.go
		# command-line-arguments
		.\test.go:20:3: no new variables on left side of :=
	*/

	/*PERCOBAAN 3
	sekarang bagaimana jika kita mendapatkan dua atau lebih value dari function, dan kita ingin menggunakan  variable yang sudah ada, untuk menampuk salah satu nilainya. Mana yang harus kita gunakan? := atau =?
	jawabannya adalah := silahkan lihat contoh dibawah
	*/
	i, k := 3, 4
	fmt.Println(i, " ", k)
	j, k := 1, 2

	fmt.Println(i, " ", k)
	fmt.Println(j, " ", k)
	/* HASIL
		$ go run test.go
		3   4
		3   2
		1   2

		Penjelasan
		karena vaiable J di baru saja deklarasikan pada saat statement kedua, maka kita boleh menggunakan := kepada j untuk melakukan inisialisasi, meskipun k sudah di definisikan pada statement sebelumnya
	*/
}