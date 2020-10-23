package main

import (
	"encoding/binary"
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
	case "super":
		//supernetting
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

	if argMask == 32 {
		os.Exit(1)
	}

	var newMask uint8 = argMask
	fmt.Printf("Subnetting %s ", argIP)
	var netw32 uint32 = IPstringToUint32(argIP)

	if argMode == 'n' {
		fmt.Printf("in %d subnets\n", PowUint(2, log2S(uint32(argN))))
		newMask += uint8(log2S(uint32(argN)))
		divideNetwork(netw32, argMask, newMask)
	} else if argMode == 'h' {
		fmt.Printf("in subnets for %d users\n", argN)
		newMask = uint8(log2S(PowUint(2, uint32(argN))))
	}
}

func divideNetwork(network uint32, oldmask uint8, newmask uint8) []uint32 {
	var num uint32 = 2 << (newmask-oldmask-1) 
	var offset uint32 = 32 - uint32(newmask)
	var i uint32 = 0
	var netwkSlice []uint32
	//netwkSlice = append(netwkSlice, network)
	convertUint32ToOctets(network)
	for i = 0; i < num; i = i + (1 << offset) {
		var newNetw uint32 = network + uint32(num << offset)
		netwkSlice = append(netwkSlice, newNetw)
		convertUint32ToOctets(newNetw)
	}
	return netwkSlice
}

// Parses an ip address in A.B.C.D format and converts it to a 32 bits unsigned int
func IPstringToUint32(netwStr string) uint32 {
	var IP uint32 = 0
	var dots uint8   // Number of dots to know which octet
	var mul byte = 1 // Weight of the parsed number in the string

	// Read the IP backwards
	for i := len(netwStr) - 1; i >= 0; i-- {
		if netwStr[i] == '.' {
			mul = 1
			dots++
		} else {
			IP += uint32(netwStr[i]-'0') * PowUint(256, uint32(dots))
			mul *= 10
		}
	}
	return IP
}

func mask2Uint32(mask uint8) uint32 {
	var mask32 uint32 = 0
	var i uint32
	for i = 31; i >= 0 && mask != 0; i-- {
		mask32 += 1 << i
		mask--
	}
	return mask32
}

func convertUint32ToOctets(address uint32) [4]uint8 {
	var octets [4]uint8
	h := address
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, h)
	for i := 0; i < 4; i++ {
		octets[i] = a[i]
	}
	return octets
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
	if num == 0 {
		return 1
	} else if num == 1 {
		return base
	} else {
		return base * PowUint(base, num-1)
	}
}
