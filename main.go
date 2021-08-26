package main // import "hello"

import (
	"machine"
	"time"
)

func main() {
	// led := machine.LED
	led := machine.GP18
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
