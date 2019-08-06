package main

import(
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	//pada kasus ini kita akan mencoba untuk menggunakan "external package", dengan nama package simpleGitHub yang berada di https://github.com/mactsouk/go.
	fmt.Println(simpleGitHub.AddTwo(5,6))
	
	//jika kita langsung menjalankan program ini dengan command
	/** 
	go run downloadPackages.go, maka hasilnya adalah : 
	downloadPackages.go:5:2: cannot find package "github.com/mactsouk/go/simpleGitHub" in any of:
        C:\Go\src\github.com\mactsouk\go\simpleGitHub (from $GOROOT)
        C:\Users\NAMA_USER\go\src\github.com\mactsouk\go\simpleGitHub (from $GOPATH)
	*/

	/* 
	masalahnya adalah, kita harus memiliki package yang akan kita gunakan di komputer kita. untuk mendownload package tersebut kita jalankan perintah berikut:
	go get -v github.com/mactsouk/go/simpleGitHub

	$ go get -v github.com/mactsouk/go/simpleGitHub
	github.com/mactsouk/go (download)
	github.com/mactsouk/go/simpleGitHub

	silahkan anda cek folder
		C:\Go\src\github.com\mactsouk\go\simpleGitHub (from $GOROOT)
	atau 
		C:\Users\NAMA_USER\go\src\github.com\mactsouk\go\simpleGitHub (from $GOPATH)
	untuk memastikan bahwa anda telah mendownload packages-nya ke komputer Anda
	*/
	

	/*
	jika program ini kita run maka akan menghasilkan
		$ go run downloadPackages.go
		11
	*/
}
