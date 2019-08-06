package main
/*
standard output kurang lebih bekerja dengan menampilkan sesuatu di layar, akan tetapui 'standard output' mungkin saja memerlukan penggunaan function yang tidak ada di package fmt, dan sekarang akan kita bahas
*/
import(
	"io"
	"os"
)


func main(){
	/* Tips   
		myString := ""
		sama artinya dengan 
		var myString string = ""
		jadi menggunakan := sama dengan membuat var (variable) baru
	*/
	myString := "" //variable ini akan menampung text yang akan dicetak di layar, entah itu dari arguments[1] ataupun String hardcode  ("Please give me one argument!")

	arguments := os.Args //membaca argument yang dimasukkan, dan menyimpannya dalam var arguments
	
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
		// io.WriteString(os.Stdout, myString)
		// io.WriteString(os.Stdout, "\n")
	} else {
		myString = arguments[1]
		// io.WriteString(os.Stdout, myString)
		// io.WriteString(os.Stdout, "\n")
	}
	
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

	/*
		dalam kasus ini function io.WriteString bekerja dengan cara yang sama seperti function fmt.Print(), hanya saja dia memerlukan 2 parameter, yang pertama adalah ke file apa kita mau melakukan write, parameter kedua adalah variable Stringnya.

		NOTE: Strictly speaking, the type of the first parameter of the
		io.WriteString() function should be an io.Writer interface, which
		requires a slice of bytes as the second parameter. However, in this case, a
		string does the job just fine. You will learn more about slices in Chapter
		3, Working with Basic Go Data Types.

	*/
}