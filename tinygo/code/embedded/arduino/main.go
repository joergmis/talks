package main

import (
	"machine"
	"time"
)

// https://tinygo.org/docs/reference/microcontrollers/arduino/
// tinygo flash -target=arduino code/embedded/arduino/main.go
func main() {
	pin := machine.D4
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if pin.Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 500)
	}

}
