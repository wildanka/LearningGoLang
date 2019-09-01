package main

import(
	"os"
	"log"
)

func main()  {
	file, e := os.OpenFile("sample.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	//0666 is permission to create a new file

	if e != nil{
		log.Fatalln("failed")
	}

	log.SetOutput(file)
	log.Println("Success")
}