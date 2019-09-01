package main

import(
	"errors"
	"fmt"
)

func returnError(a, b int) error{
	if a==b {
		err := errors.New("Error in returnError() function!") //string di dalam New nantinya dapat di akses oleh method err.Error()  (tidak harus err sih, tergantung variable errornya di method pemanggilnya juga, contoh: lihat line 18, 25, 32)
		return err
	}else{
		return nil
	}
}

func main()  {
	err := returnError(2,3)
	if err == nil{
		fmt.Println("returnError() ended normally! (which means no error returned, it returns nil instead")
	}else{
		fmt.Println(err)
	}

	err = returnError(10,10)
	if err == nil{
		fmt.Println("returnError() ended normally! (which means no error returned, it returns nil instead")
	}else{
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() function!"{ // lihat line 10
		fmt.Println("!!")
	}
	/* 
		OUTPUT : 
		$ go run newError.go
		returnError() ended normally! (which means no error returned, it returns nil instead
		Error in returnError() function!
		!!

		jika kita tidak mencoba langsung membandngkan antara [err] dengan string "Error in returnError() function!" (contoh kode sebagai berikut) : 
		if err == "Error in returnError() function!"{ // lihat line 10
			fmt.Println("!!")
		}
		
		maka akan terjadi error seperti berikut: 
		$ go run newError.go
		# command-line-arguments
		.\newError.go:32:9: invalid operation: err == "Error in returnError() function!" (mismatched types error and string)
	*/
}