package main

import(
	"io"
	"os"
)

func main()  {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1{
		myString = "Please give me one arguments!"
	}else{
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "Hallo \n")
// jika kita jalankan program secara normal maka akan menghasilkan 3 output
/*
	$ go run stdERR.go
	This is Standard output
	Please give me one arguments!Hallo
	sepintas ketika kita jalankan di terminal, stdError yang kita tampilkan (Hallo \n) akan terlihat seperti output pada umumnya ketika kita menggunakan stdout, namun sebenarnya ada sebuah trik yang dapat kita gunakan untuk memudahkan kita membedakan mana yang merupakan stdError:
	$ go run stdERR.go 2>tmp/stdError
*/
}