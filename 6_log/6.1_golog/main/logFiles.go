package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)


func main()  {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err!=nil{
		log.Fatal(err)
	}else{
		log.SetOutput(sysLog)
	}
	log.Println(" LOG_INFO + LOG_LOCAL7 : logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some Program!")
	if err != nil{
		log.Fatal(err)
	}else{
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
	/*
	The last part shows that you can change the logging configuration in your programs as
	many times as you want, and that you can still use fmt.Println() for printing output on
	the screen.
	*/
}

/* 
The first parameter of the syslog.New() function is the priority, which is a combination of
the logging facility and the logging level. Therefore, a priority of LOG_NOTICE |
LOG_MAIL, which is mentioned as an example, will send notice logging-level messages to
the MAIL logging facility.

As a result, the preceding code sets the default logging to the local7 logging facility using
the info logging level. The second parameter of the syslog.New() function is the name of
the process that will appear on the logs as the sender of the message. Generally speaking, it
is considered a good practice to use the real name of the executable in order to be able to
find the information you want easily in the log files at another time.

After the call to syslog.New(), you will have to check the error variable that it returns so
that you can make sure that everything is fine. If everything is OK, which means that the
value of the error variable is equal to nil, you call the log.SetOutput() function. This
sets the output destination of the default logger, which in this case is the logger you created
earlier on (sysLog). Then you can use log.Println() to send information to the log
server.
 */


 /* 
 https://golang.org/pkg/log/syslog/
 if we run the script in windows it will show something like this
	$ go run logFiles.go
	# command-line-arguments
	.\logFiles.go:14:17: undefined: syslog.New
	.\logFiles.go:14:28: undefined: syslog.LOG_INFO
	.\logFiles.go:14:44: undefined: syslog.LOG_LOCAL7
	.\logFiles.go:23:16: undefined: syslog.New
	.\logFiles.go:23:27: undefined: syslog.LOG_MAIL

	this is caused by a bugs in GO, because "syslog" package is frozen in windows:
	* This package is not implemented on Windows. As the syslog package is frozen, Windows users are encouraged to use a package outside of the standard library. For background, see https://golang.org/issue/1108. 
	* This package is not implemented on Plan 9. 
	* This package is not implemented on NaCl (Native Client). 
 */