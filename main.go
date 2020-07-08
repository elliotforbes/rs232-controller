package main

import (
	"fmt"
	"os"
	"time"

	"github.com/elliotforbes/rs232-controller/rs232"
	"github.com/nathan-osman/go-sunrise"
)

const (
	HALFBRIGHTNESS = "HEXCODE"
)

func sendRs232Command(com_port string) {
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

	n, err := port.Write([]byte(HALFBRIGHTNESS))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}

func main() {

	com_port := os.Args[1]

	for {
		year, month, day := time.Now().Date()

		_, set := sunrise.SunriseSunset(
			51.5074, 0.1278, // London
			year, month, day,
		)

		if time.Now().Sub(set) > 0 {
			fmt.Println("After Sunset, sending rs232 command")
			sendRs232Command(com_port)
		} else {
			fmt.Println("Before Sunset, not sending")
		}
		time.Sleep(1 * time.Minute)
	}

}
