package main

import (
	"machine"
	"time"
)

// https://tinygo.org/docs/reference/microcontrollers/nucleo-f103rb/
// tinygo flash -target=nucleo-f103rb code/embedded/nucleo/main.go
func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
