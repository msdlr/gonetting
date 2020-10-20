package main

import (
	"fmt"     // Print
	"os"      // Arguments passed to the program
	"strconv" // Atoi
)

/*
Arguments:
A.B.C.D [1-30] -n [2,4,8]
A.B.C.D [1-30] -h x
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
	default:
		if len(os.Args) < 5 {
			fmt.Println("Not enough arguments. See --help")
			os.Exit(1)
		}
		argNetmask, _ := strconv.ParseInt(os.Args[2], 10, 0) // ParseInt returns an error code too
		argN, _ := strconv.ParseInt(os.Args[4], 10, 0)
		subnetting(os.Args[1], uint(argNetmask), byte(os.Args[3][1]), uint(argN))
	}
}

func help() {
	fmt.Printf("[%s] Help menu: \n", os.Args[0])
}
func subnetting(argIP string, argMask uint, argMode byte, argN uint) {
	fmt.Printf("Subnetting %s ", argIP)
	if argMode == 'n' {
		fmt.Printf("in %d subnets\n", argN)
	} else if argMode == 'h' {
		fmt.Printf("in subnets for %d users\n", argN)
	}
}

func stringToIP(ipString string, ipArrayPointer *[4]byte) {
	var dots uint    // Number of dots to know which octet
	var mul byte = 1 // Weight of the parsed number in the string
	var ipOctets [4]byte

	// Read the IP backwards
	for i := len(ipString) - 1; i >= 0; i++ {
		if ipString[i] == '.' {
			dots++ // We change octet
			break  // Next loop iteration
		} else {
			mul *= 10
		}
		// Which ip octet to parse depends on the dots we encountered
		ipOctets[4-dots] = mul * ipString[i]
	}
	fmt.Printf("\n%d.%d.%d.%d\n", ipOctets[0], ipOctets[1], ipOctets[2], ipOctets[3])
	//return ipOctets
}
