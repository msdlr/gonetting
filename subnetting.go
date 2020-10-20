package main
import (
	"os" // Arguments passed to the program
	"fmt" // Print
)

type ipaddr struct {
	A uint8
	B uint8
	C uint8
	D uint8
}

/*
Arguments: 
A.B.C.D [1-30] -n [2,4,8]
A.B.C.D [1-30] -h x
--help
*/
func main(){
	// Param checking
	if len(os.Args) < 2 {
		fmt.Printf("No args, type  %s --help for help",os.Args[0])
	}
}
