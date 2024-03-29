package main
//# di dalam Go, ada 2 hal yang harus kita tanamkan dalam mindset kita:
// 1. Kita akan menggunakan Go Package
// 2. Jangan menambahkannya sama sekali

//Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
//source : https://golang.org/pkg/fmt/
import(
	"fmt" 
	_"os" // sebenarnya kita juga bisa untuk melakukan "bypass" terhadap aturan import Go, yaitu dengan menambahkan karakter _ (undescore) di depan package yang tidak digunakan. mengapa Go memiliki fitur ini? nanti akan dibahas, sabar ya :D
)

// di kasus ini jika file packageNotUsed.go kita execute (dengan go run packageNotUsed.go / bukan compile) maka kita akan mendapatkan pesan error sebagai berikut:
// .\packageNotUsed.go:10:2: imported and not used: "os"
// error tersebut disebabkan karena kita mengimport Go package (os) tapi tidak menggunakan pakcage tersebut

func main(){
	fmt.Println("Hello brother!")
}