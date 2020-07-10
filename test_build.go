package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/elliotforbes/rs232-controller/rs232"
)

func sendRs232Command(com_port string, hex_command string, bit_rate int) {
	port, err := rs232.Open(com_port, rs232.Options{
		BitRate:  uint32(bit_rate),
		DataBits: 8,
		StopBits: 1,
		Parity:   rs232.PARITY_NONE,
		Timeout:  0,
	})
	if err != nil {
		fmt.Println(err)
	}

	data, err := hex.DecodeString(hex_command)
	if err != nil {
		panic(err)
	}

	n, err := port.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}

func main() {

	com_port := os.Args[1]
	bit_rate, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Second Argument needs to be a number")
	}
	hex_command := os.Args[3]

	fmt.Println("sending test rs232 command")
	sendRs232Command(com_port, hex_command, bit_rate)
	fmt.Println("Sent test rs232 command")

}
