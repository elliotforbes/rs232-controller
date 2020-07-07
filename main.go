package main

import (
    "fmt"
    "time"
	"github.com/nathan-osman/go-sunrise"
	"github.com/chimera/go-inside/rs232"
)

const (
	HALFBRIGHTNESS = "HEXCODE"
	COM_PORT = "COM1"
)

func sendRs232Command() {
	port, err := rs232.Open(COM_PORT, rs232.Options{
		BitRate: 115200,
		DataBits: 8,
		StopBits: 1,
		Parity: rs232.PARITY_NONE,
		Timeout: 0,
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

	for {
		year, month, day := time.Now().Date()
	
		_, set := sunrise.SunriseSunset(
			51.5074, 0.1278,          // London
			year, month, day, 
		)
		
		if time.Now().Sub(set) > 0 {
			fmt.Println("After Sunset, sending rs232 command")
			sendRs232Command()
		} else {
			fmt.Println("Before Sunset, not sending")
		}
	}

}