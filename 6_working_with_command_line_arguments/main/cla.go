package main

import(
	"fmt"
	"os" //untuk melakukan operasi command-line tentunya kita perlu package os untuk berkomunikasi dengan OS
	"strconv" //strconv akan mengkonversi "command-line arguments" yang berupa string, ke dalam tipe data aritmetik
)

func main()  {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1);
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < count; i++ {
		n, _ := strconv.ParseFloat(arguments[1], 64)
		if n < min{
			min = n
		}
		if n > max{
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

//cla.go akan memeriksa apakah kita memiliki "command-line arguments" atau tidak dengan menguji panjang dari os.args, karena program memerlukan setidaknya satu command-line argument untuk bisa beroperasi.
//perlu diingat bahwa os.Args akan "Go Slice" yang bertipe data string, 
//element pertama di dalam "slice" adalah nama program (executable program), maka dari itu untuk menginisialisasi variable min dan max kita memerlukan element kedua dari String os.Args yang memiliki index 1.
/* Mulustrasi:
                 0								1							2
os.Args = ['nama_program', 'isi_os.Args_yang_diinput_user', 'error']
*/


/*
n, _ := strconv.ParseFloat(arguments[i], 64)
contoh di atas mengasumsikan bahwa semua input pada command-line telah sesuai dengan yang diinginkan oleh program dan mengabaikan error value yang mungkin di-return oleh function strconv.ParseFloat() (pada slice ke-2).
statement di atas memberitahukan kepada Go bahwa kita hanya ingin mengambil value pertama yang di-return oleh strconv.ParseFloat() dan kita tidak tertarik pada value ke2 (yang pada kasus ini adalah "Error Variable") dan kita masukan ke dalam underscore.
# Underscore character pada GO disebut sebagai "blank identifier"
blank identifier adalah cara GO men-discard sebuah nilai (tidak dipedulikan lagi, hiks)

WARNING : melakukan ignore terhadap semua atau beberapa nilai return dari sebuah Go function apalagi error variable kedalam blank identifier sangatlah berbahaya (bad practices) dan sangat tidak disarankan untuk digunakan pada fase produksi. 

*/