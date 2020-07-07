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
	port, err := rs232.Open(COM_PORT, &rs232.Options{})
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
    // You can use the Parameters structure to set the parameters
    rise, set := sunrise.SunriseSunset(
		51.5074, 0.1278,          // London
		2020, time.July, 7,  // 2000-01-01
	)
	
	//51.5074° N, 0.1278° W
	
	fmt.Println(rise)
	fmt.Println(set)


	sendRs232Command()

}