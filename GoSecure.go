/*
A blinker example using go-rpio library.
Requires administrator rights to run
Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.
*/

package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
	"os/signal"
	"syscall"
)

var (
	pin = rpio.Pin(2)
	pin21 = rpio.Pin(21)
)

func main() {

	running := true

	// Handle Exit statements
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
			case os.Interrupt:
			running = false
			case syscall.SIGTERM:
			running = false
		}
	}()

	println("Attempting to open pin")

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	// Unmap gpio memory when done
	defer cleanup()

	// Set pin to output mode
	pin.Output()
	pin.PullDown()
	pin.Low()
	pin.Write(rpio.Low)

	pin.Input()
	pin.PullOff()
	pin.PullDown()

	//pin21.Input()
	//pin21.PullDown()

	for running {

		//println("changeing pin")
		//pin.Toggle()

		println("reading pin")
		state := pin.Read()
		if (state == rpio.High) {
			println("High")
		} else {
			println("low")
		}
		time.Sleep(time.Second)
	}
}


func cleanup() {
	println("Trying to close RPIO library")
	rpio.Close()
	println("Closed RPIO Library")
}