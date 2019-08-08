package main
/*
The official description of the os package tells us
that it offers functions that perform operating system operations. This includes functions for
creating, deleting, and renaming files and directories, as well as functions for learning the
Unix permissions and other characteristics of files and directories
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(">", scanner.Text())
	}
}