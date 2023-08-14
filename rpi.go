
// Package rpi is a library for interfacing with Raspberry Pi IO
package rpi

// This is the wrong url here ??? 
import (
	_ "github.com/kmmndr/goPi/MCP23S17"
	_ "github.com/kmmndr/goPi/ioctl"
	_ "github.com/kmmndr/goPi/spi"
	_ "github.com/kmmndr/goPi/piface"
)
