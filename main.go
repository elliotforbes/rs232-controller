package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/elliotforbes/rs232-controller/rs232"
	"github.com/nathan-osman/go-sunrise"
)

func sendRs232Command(com_port string, hex_code string, bit_rate int) {
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

	data, err := hex.DecodeString(hex_code)
	if err != nil {
		panic(err)
	}

	n, err := port.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}

func get7am(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 7, 0, 0, 0, t.Location())
}

func main() {

	com_port := os.Args[1]
	bit_rate, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Second Argument needs to be a number")
	}
	half_brightness := os.Args[3]
	full_brightness := os.Args[4]

	for {

		today := time.Now()

		set_half_brightness := false
		set_full_brightness := false

		year, month, day := time.Now().Date()

		_, set := sunrise.SunriseSunset(
			51.5074, 0.1278, // London
			year, month, day,
		)

		// if time after sunset then send half_brightness
		if time.Now().Sub(set) > 0 {
			if set_half_brightness == false {
				fmt.Println("After Sunset, sending rs232 command")
				sendRs232Command(com_port, half_brightness, bit_rate)
				set_half_brightness = true
				set_full_brightness = false
			}
		}

		// if past 7am and less than sunset
		// send full brightness command
		if get7am(today).Sub(time.Now()) > 0 && time.Now().Sub(set) < 0 {
			if set_full_brightness == false {
				sendRs232Command(com_port, full_brightness, bit_rate)
				fmt.Println("Before Sunset, not sending")
				set_full_brightness = true
				set_half_brightness = false
			}
		}

		time.Sleep(1 * time.Minute)
	}

}
