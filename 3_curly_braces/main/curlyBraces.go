package main

import(
	"fmt"
)
//sekilas, koding disini terlihat normal, tapi jika kita jalankan akan menghasilkan error, kenapa?
//karena Go sebenarnya memerlukan semicolon/titik-koma (;) di setiap akhir statement (tentunya tergantung konteks). 

func main() 
{
	fmt.Println("Hello brother!")
}

/**
yang dilakukan oleh Go adalah

func main(); 
{
	fmt.Println("Hello brother!");
}

dan tentu saja, ini akan menghasilkan error :)
*/