package main

import (
    "fmt"
    "time"
	"github.com/nathan-osman/go-sunrise"
	"github.com/chimera/go-inside/rs232"
)

const (
	HALFBRIGHTNESS = "HEXCODE"
)

// func sendRs232Command() {
// 	port, err := rs232.Open("Controll Dimness", &rs232.Options{})
// 	if err != nil {
// 		fmt.Println(err)
// 	}	

// 	n, err := port.Write([]byte(HALFBRIGHTNESS))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(n)

}

func main() {
    // You can use the Parameters structure to set the parameters
    rise, set := sunrise.SunriseSunset(
		43.65, -79.38,          // Toronto, CA
		2000, time.January, 1,  // 2000-01-01
	)


	// sendRs232Command()

}