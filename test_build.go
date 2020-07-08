package main

import (
	"fmt"
	"os"

	"github.com/elliotforbes/rs232-controller/rs232"
)

func sendRs232Command(com_port string, hex_command string) {
	port, err := rs232.Open(com_port, rs232.Options{
		BitRate:  115200,
		DataBits: 8,
		StopBits: 1,
		Parity:   rs232.PARITY_NONE,
		Timeout:  0,
	})
	if err != nil {
		fmt.Println(err)
	}

	n, err := port.Write([]byte(hex_command))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}

func main() {

	com_port := os.Args[1]
	hex_command := os.Args[2]

	fmt.Println("sending test rs232 command")
	sendRs232Command(com_port, hex_command)
	fmt.Println("Sent test rs232 command")

}
