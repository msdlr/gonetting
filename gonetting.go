package main

import (
	"fmt" // Print
	"math"
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
		subnetting(os.Args[1], uint8(argNetmask), byte(os.Args[3][1]), uint8(argN))
	}
}

func help() {
	fmt.Printf("[%s] Help menu: \n", os.Args[0])
}
func subnetting(argIP string, argMask uint8, argMode uint8, argN uint8) {
	// De-reference parameter
	var ipArgPtr *string = &argIP

	// Create array for 
	ipArray := [4]uint8{0, 0, 0, 0}
	var ipArrayPtr *[4]uint8 = &ipArray
	var maskArray [4]uint8
	var maskArrayPtr *[4]uint8 = &maskArray

	// Parse the ip and mask
	IPstringToArray(ipArgPtr, uint8(len(argIP)), ipArrayPtr)
	maskStringToArray(argMask, maskArrayPtr)


	fmt.Printf("Subnetting %s ", argIP)
	if argMode == 'n' {
		fmt.Printf("in %d subnets\n", argN)
	} else if argMode == 'h' {
		fmt.Printf("in subnets for %d users\n", argN)
	}
}

func IPstringToArray (ipStringPointer *string, size uint8, ipArrayPointer *[4]uint8) {
	var dots uint    // Number of dots to know which octet
	var mul byte = 1 // Weight of the parsed number in the string
	var ipOctets [4]uint8
	var ipString = *ipStringPointer

	// Read the IP backwards
	for i := size - 1; i > 0; i-- {
		if ipString[i-1] == '.' {
			dots++ // We change octet
			mul = 1
		} else {
			ipOctets[3-dots] += mul * (ipString[i-1] - '0')
			mul *= 10
		}
		// Which ip octet to parse depends on the dots we encountered
	}
	*ipArrayPointer = ipOctets 
	fmt.Printf("IP: %d.%d.%d.%d\n", ipOctets[0], ipOctets[1], ipOctets[2], ipOctets[3])
}

func maskStringToArray (size uint8, maskArrayPointer *[4]uint8) {
	var maskOctets [4]byte
	var remaining uint8 = size
	for octet := 0 ; octet < 4; octet++ {
		if remaining >= 8 {
			maskOctets[octet] = maskOctet(8)
			remaining -= 8
		} else {
			maskOctets[octet] = maskOctet(remaining)
		}
	} 
	*maskArrayPointer = maskOctets
	fmt.Printf("Mask: %d.%d.%d.%d\n", maskOctets[0], maskOctets[1], maskOctets[2], maskOctets[3])
}

func maskOctet(n uint8) (uint8) {
	switch (n) {
		case 0:
			return 0
		default:
			return uint8(math.Pow(float64(2),float64(8-n))) + maskOctet(n-1)
	}
}