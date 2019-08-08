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