package main
import (
	"os" // Arguments passed to the program
	"fmt" // Print
	//"net" // IP address type; type IP []byte
)

/*
Arguments: 
A.B.C.D [1-30] -n [2,4,8]
A.B.C.D [1-30] -h x
--help
*/
func main(){
	fmt.Println(os.Args[0])
	fmt.Printf("N params: %d", len(os.Args))
	// Param checking
	if len(os.Args) < 2 {
		fmt.Printf("No args, type  %s --help for help",os.Args[0])
	}
}
