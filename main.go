package main // import "hello"

import (
	"machine"
	"time"
)

func main() {
	// led := machine.LED // raspberry pi pico
	led := machine.GP18 // pimoroni tiny2040
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
