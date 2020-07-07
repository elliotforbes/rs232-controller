package main

import (
    "fmt"
    "time"
	"github.com/kelvins/sunrisesunset"
	"github.com/chimera/go-inside/rs232"
)

const (
	HALFBRIGHTNESS = "HEXCODE"
)

func sendRs232Command() {
	port, err := rs232.Open("Controll Dimness", &rs232.Options{})
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
    p := sunrisesunset.Parameters{
      Latitude:  -23.545570,
      Longitude: -46.704082,
      UtcOffset: -3.0,
      Date:      time.Date(2017, 3, 23, 0, 0, 0, 0, time.UTC),
    }

    // Calculate the sunrise and sunset times
    sunrise, sunset, err := p.GetSunriseSunset()

    // If no error has occurred, print the results
    if err == nil {
        fmt.Println("Sunrise:", sunrise.Format("15:04:05")) // Sunrise: 06:11:44
        fmt.Println("Sunset:", sunset.Format("15:04:05")) // Sunset: 18:14:27
    } else {
        fmt.Println(err)
	}
	
	sendRs232Command()

}