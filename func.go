package main

import (
	"encoding/binary"
	"fmt" // Print
	"os"  // Arguments passed to the program
)

func subnetting(argIP string, argMask uint8, argMode uint8, argN uint8) {
	// up to /32 mask
	if argMask == 32 {
		os.Exit(1)
	}

	var newMask uint8 = argMask
	fmt.Printf("Subnetting [%s/%d] ", argIP, argMask)

	// Convert ip as string to uint32
	var netw32 uint32 = IPstringToUint32(argIP)
	convertUint32ToOctets(netw32)

	// Get the slice of networks as uint32
	var netwSlice []uint32

	if argMode == 'n' {
		// New mask = mask + log2S(n)
		newMask += uint8(log2S(uint32(argN)))
		fmt.Printf("in %d /%d subnets\n", 1<<log2S(uint32(argN)), newMask)
		netwSlice = calculateSubnets(netw32, argMask, newMask)
	} else if argMode == 'h' {
		// New mask = 32 - log2( 2 ^ h )
		newMask = 32 - uint8(log2S(uint32(argN+2)))
		fmt.Printf("in /%d subnets for %d users (size %d)\n", newMask, argN, 1<<log2S(uint32(argN+2)))
		netwSlice = calculateSubnets(netw32, argMask, newMask)
	}
	printNetwSlice(netwSlice, newMask)
}

func supernetting(networks []uint32, mask uint8) {

}

func calculateSubnets(network uint32, oldmask uint8, newmask uint8) []uint32 {
	// Number of networks
	var num uint32 = 2 << (newmask - oldmask - 1)
	// Numbers of the new mask that are set to 0
	var offset uint32 = 32 - uint32(newmask)
	var i uint32 = 0
	var netwkSlice []uint32
	// Calculate the new network addresses
	for i = 0; i < num; i++ {
		var newNetw uint32 = network + uint32(i<<offset)
		netwkSlice = append(netwkSlice, newNetw)
		convertUint32ToOctets(newNetw)
	}
	return netwkSlice
}

func printNetwSlice(slice []uint32, mask uint8) {
	var octets [4]uint8

	for i := 0; i < len(slice); i++ {
		octets = convertUint32ToOctets(slice[i])
		fmt.Printf("%d:\t[%d.%d.%d.%d/%d]\n", i+1, octets[0], octets[1], octets[2], octets[3], mask)
	}
}

// IPstringToUint32 parses an ip address in A.B.C.D format and converts it to a 32 bits unsigned int
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
			//IP += uint32(netwStr[i]-'0') * PowUint(256, uint32(dots))
			IP += uint32(netwStr[i]-'0') * uint32(mul) << (uint8(8) * dots)
			mul *= 10
		}
	}
	return IP
}

func mask2Uint32(mask uint8) uint32 {
	var mask32 uint32 = 0
	var i uint32
	// Add one ^ (32-n) for every bit == 1
	for i = 31; i >= 0 && mask != 0; i-- {
		mask32 += 1 << i
		mask--
	}
	return mask32
}

func convertUint32ToOctets(address uint32) (octets [4]uint8) {
	// We just copy the contents of the uint32 into 4 uint8s
	binary.BigEndian.PutUint32(octets[0:4], uint32(address))
	return
}

// Log2 of a number but increments 1 if not exact
func log2S(n uint32) (log uint32) {
	for i := 1; i < 31; i++ {
		// Find which 2^i is equal or bigger than n
		if n <= 1<<i {
			log = uint32(i)
			break
		}
	}
	return log
}
