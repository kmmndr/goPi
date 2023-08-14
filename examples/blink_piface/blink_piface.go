package main

import (
	"fmt"
	"github.com/kmmndr/goPi/piface"
	"github.com/kmmndr/goPi/spi"
	"time"
)

func main() {

	// creates a new pifacedigital instance
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// initializes pifacedigital board
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}

	// blink time!!
	fmt.Println("Blinking led 7 each second")
	for {
		pfd.Leds[7].Toggle()
		time.Sleep(time.Second)
	}
}
