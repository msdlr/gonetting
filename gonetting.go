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

	// Parse the ip and mask
	IPstringToUint32(ipArgPtr, uint8(len(argIP)))

	fmt.Printf("Subnetting %s ", argIP)
	if argMode == 'n' {
		fmt.Printf("in %d subnets\n", argN)

		//getSubnets(ipArrayPtr, argMask, argN, newSubnetsPtr)
	} else if argMode == 'h' {
		fmt.Printf("in subnets for %d users\n", argN)
	}
}

func IPstringToUint32(netwStrPtr *string, size uint8) uint32 {
	//var IP uint32 = 0
	//var dots uint    // Number of dots to know which octet
	//var mul byte = 1 // Weight of the parsed number in the string
	//var ipOctets [4]uint32
	//var ipString = *netwStrPtr
	//
	//// Read the IP backwards
	//for i := size - 1; i >= 0; i-- {
	//	if ipString[i] == '.' {
	//		dots++ // We change octet
	//		mul = 1
	//	} else {
	//		ipOctets[3-dots] += uint32(mul * (ipString[i-1] - '0'))
	//		mul *= 10
	//	}
	//	// Which ip octet to parse depends on the dots we encountered
	//}
	//fmt.Printf("IP: %d.%d.%d.%d\n", ipOctets[0], ipOctets[1], ipOctets[2], ipOctets[3])
	//
	//IP = ipOctets[0]
	//for i := 1; i < 3; i++ {
	//	IP += ipOctets[i] << (8 * i)
	//}
	//return IP
}

func log2S(n uint32) (log uint32) {
	// Get the real log2
	var res float64 = math.Log2(float64(n))

	// If the real log2 is not an integer we add 1 more
	if res-float64(int64(res)) == 0.0 {
		return uint32(res)
	}
	return uint32(res + 1)
}

func PowUint(base uint32, num uint32) (result uint32) {
	if num == 1 {
		return base
	} else {
		return base * PowUint(base, num-1)
	}
}
