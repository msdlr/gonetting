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
		// Reserve memory for the new addresses

		// Call the subnetting function
		var newSubnets [][4]uint8
		var newSubnetsPtr *[][4]uint8 = &newSubnets
		//getSubnets(ipArrayPtr, argMask, argN, newSubnetsPtr)
	} else if argMode == 'h' {
		fmt.Printf("in subnets for %d users\n", argN)
	}
}

func IPstringToArray(netwStrPtr *string, size uint8, netwArrayPtr *[4]uint8) {
	var dots uint    // Number of dots to know which octet
	var mul byte = 1 // Weight of the parsed number in the string
	var ipOctets [4]uint8
	var ipString = *netwStrPtr

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
	*netwArrayPtr = ipOctets
	fmt.Printf("IP: %d.%d.%d.%d\n", ipOctets[0], ipOctets[1], ipOctets[2], ipOctets[3])
}

func maskStringToArray(size uint8, maskArrayPointer *[4]uint8) {
	var maskOctets [4]byte
	var remaining uint8 = size
	for octet := 0; octet < 4; octet++ {
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

func maskOctet(n uint8) uint8 {
	switch n {
	case 0:
		return 0
	default:
		return uint8(math.Pow(float64(2), float64(8-n))) + maskOctet(n-1)
	}
}

//func getSubnets(networkPtr *[4]uint8, oldMask uint8, newNetworks uint8 ,newSubnetsPtr *[][4]uint8){
//	// De-reference and reserve memory
//	var network [4]uint8
//	var maskOffset uint8 = log2S(newNetworks)
//	var newMask uint8 = oldMask + maskOffset
//
//	// Get which octets are modified
//	var octet1 uint8 = oldMask / 8
//	var octet2 uint8 = newMask / 8
//
//	// Get bit of the octet
//	var oldMaskBit uint8 = oldMask % 8
//	var newMaskBit uint8 = oldMaskBit + maskOffset
//
//	//Variables for calculating the new networks
//
//	var newSubnets [][4]uint8
//
//	for i:=0 ; uint8(i) < newNetworks ; i++ {
//		// Copy the octets of the old network onto the new ones
//		append(newSubnets,network)
//		//newSubnets[i][0] = network[0]
//		//newSubnets[i][1] = network[1]
//		//newSubnets[i][2] = network[2]
//		//newSubnets[i][3] = network[3]
//	}
//
//	if octet1 == octet2 {
//		// Same octet, simple math
//		// Octet in the new networks: maskOctet(1-oldMaskBit) + ( {0..((2^maskOffset)-1) * 2^(7-newMaskBit) }}
//		for newNetworkNum := 0 ; uint8(newNetworkNum) < ((PowUint(2,maskOffset ))-1) ; newNetworkNum++ {
//			newSubnets[newNetworkNum][octet1] = maskOctet(1-oldMaskBit) + (PowUint(2, uint8(newNetworkNum)) - 1)*PowUint(2, 7 - newMaskBit)
//		}
//	}
//	*newSubnetsPtr = newSubnets
//}

func log2S(n uint8) (log uint8) {
	// Get the real log2
	var res float64 = math.Log2(float64(n))

	// If the real log2 is not an integer we add 1 more
	if res-float64(int64(res)) == 0.0 {
		return uint8(res)
	}
	return uint8(res + 1)
}

func PowUint(base uint8, num uint8) (result uint8) {
	if num == 1 {
		return base
	} else {
		return base * PowUint(base, num-1)
	}
}
