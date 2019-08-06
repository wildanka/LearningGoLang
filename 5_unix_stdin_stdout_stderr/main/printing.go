package main

import(
	"fmt" // https://golang.org/pkg/fmt/ fmt are formatter actually, while the format verbs are derived from C, but simpler *they said
)

func main(){
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"
	// v5 := 123

	fmt.Print(v1, v2, v3, v4)
	fmt.Println()
	fmt.Println(v1, v2, v3, v4)
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	//akan menghasilkan 
	/*
	$ go run printing.go
 	  123123Have a nice day
	  abc
	  123 123 Have a nice day
	   abc
	  123 123 Have a nice day
	   abc
	  123123 Have a nice day
	   abc
	*/
}