package main
/* 
since windows and GO-LANG syslog have some issues, we need to use external libraries (package) help, we can use this : 
https://github.com/hashicorp/go-syslog/blob/master/unix.go
so, let's download it by run =>  go get -v github.com/hashicorp/go-syslog

"This repository provides a very simple gsyslog package. The point of this package is to allow safe importing of syslog without introducing cross-compilation issues. The stdlib log/syslog cannot be imported on Windows systems, and without conditional compilation this adds complications." << that's what they said
 */
import(
	"fmt"
	"os"
	"log"
	"path/filepath"
	"github.com/hashicorp/go-syslog"
)

func main()  {
	programName := filepath.Base(os.Args[0])
	sysLog, err := gsyslog.NewLogger(gsyslog.LOG_INFO|gsyslog.facilityPriority("LOCAL7"), programName)

	if err!=nil{
		log.Fatal(err)
	}else{
		log.SetOutput(syslog)
	}
	log.Println(" LOG_INFO + LOG_LOCAL7 : logging in Go!")

	sysLog, err = gsyslog.New(gsyslog.LOG_MAIL, "Some Program!")
	if err != nil{
		log.Fatal(err)
	}else{
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}