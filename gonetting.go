package main

import (
	"fmt"     // Print
	"os"      // Arguments passed to the program
	"strconv" // Atoi
)

/*
Arguments:
-sub A.B.C.D [1-30] -n [2,4,8]
-sub A.B.C.D [1-30] -h x
-super n A.B.C.D A.B.C.D mask
--help
*/
func main() {
	// Param checking
	if len(os.Args) < 2 {
		fmt.Printf("No args, type  %s --help for help\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--help":
		help()
		break
	case "-super":
		callSupernetting()
		break
	case "-sub":
		callSubnetting()
		break
	default:
		if len(os.Args) < 5 {
			fmt.Println("Not enough arguments. See --help")
			os.Exit(1)
		}
	}
}

func callSubnetting() {
	argNetmask, _ := strconv.ParseInt(os.Args[3], 10, 0) // ParseInt returns an error code too
	argN, _ := strconv.ParseInt(os.Args[5], 10, 0)
	subnetting(os.Args[2], uint8(argNetmask), byte(os.Args[4][1]), uint8(argN))
}

func callSupernetting() {
	rgN, _ := strconv.ParseInt(os.Args[2], 10, 0)
	// Check that we have 2^n subnets
	if len(os.Args) == 4+int(rgN) {
		// Slice for storing uint32 networks
		var networks []uint32
		for i := 0; i < int(rgN); i++ {
			networks = append(networks, IPstringToUint32(os.Args[i+3]))
		}
		argNetmask, _ := strconv.ParseInt(os.Args[len(os.Args)-1], 10, 0)
		supernetting(networks, uint8(argNetmask))
	} else {
		fmt.Println("Not enough arguments. See --help")
		os.Exit(1)
	}
}

func help() {
	fmt.Printf("[%s] Help menu: \n", os.Args[0])
}
