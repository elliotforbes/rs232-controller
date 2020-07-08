package main

import (
	"fmt"

	"github.com/chimera/go-inside/rs232"
)

const (
	HALFBRIGHTNESS = "HEXCODE"
	COM_PORT       = "COM1"
)

func sendRs232Command() {
	port, err := rs232.Open(COM_PORT, rs232.Options{
		BitRate:  115200,
		DataBits: 8,
		StopBits: 1,
		Parity:   rs232.PARITY_NONE,
		Timeout:  0,
	})
	if err != nil {
		fmt.Println(err)
	}

	n, err := port.Write([]byte(HALFBRIGHTNESS))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}

func main() {

	fmt.Println("sending test rs232 command")
	sendRs232Command()
	fmt.Println("Sent test rs232 command")

}
