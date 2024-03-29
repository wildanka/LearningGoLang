package main

import(
	"fmt"
	"os"
	"strconv"
)

func main()  {
	//mengambil inputan dari user
	arguments := os.Args

	if len(arguments)==1 {
		fmt.Println("Silahkan masukan satu atau lebih bilangan desimal!")
		os.Exit(1);
	}

	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}


/*
$ go run forloop.go -10 0.75 3.5 12.345
Min:  -10
Max:  12.345

$ go run forloop.go -10
Min:  -10
Max:  -10

seperti yang sudah kita bicarakan sebelumnya, bahwa program yang kita tulis tidak menyaring input yang dimasukkan oleh user, maka yang seharusnya input a b c menghasilkan error tetapi malah tidak
$ go run forloop.go a b c 10
Min:  0
Max:  10

*/