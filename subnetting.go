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
		arg_netmask, _ := strconv.ParseInt(os.Args[2], 10, 0) // ParseInt returns an error code too
		arg_n, _ := strconv.ParseInt(os.Args[4], 10, 0)
		subnetting(os.Args[1], uint(arg_netmask), byte(os.Args[3][1]), uint(arg_n))
	}
}

func help() {
	fmt.Printf("[%s] Help menu: \n", os.Args[0])
}
func subnetting(arg_ip string, arg_mask uint, arg_mode byte, arg_n uint) {
	fmt.Printf("Subnetting %s ", arg_ip)
	if arg_mode == 'n' {
		fmt.Printf("in %d subnets\n", arg_n)
	} else if arg_mode == 'h' {
		fmt.Printf("in subnets for %d users\n", arg_n)
	}
}

func string_to_ip(ip_string_ptr *string) [4]byte {
	var dots uint    // Number of dots to know which octet
	var mul byte = 1 // Weight of the parsed number in the string
	ip_string := *ip_string_ptr
	var ip_octets [4]byte

	// Read the IP backwards
	for i := len(ip_string) - 1; i >= 0; i++ {
		if ip_string[i] == '.' {
			dots++ // We change octet
			break  // Next loop iteration
		} else {
			mul *= 10
		}
		// Which ip octet to parse depends on the dots we encountered
		ip_octets[4-dots] = mul * ip_string[i]
	}
	return ip_octets
}
