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
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin7 = rpio.Pin(11)
	pin3 = rpio.Pin(5)
)

func main() {

	println("Attempting to open pin")

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	// Unmap gpio memory when done
	defer cleanup()

	// Set pin to output mode
	pin3.Input()
	pin7.Output()
	pin7.PullDown()
	pin7.Low()
	// Toggle pin 20 times
	for x := 0; x < 200; x++ {

		print("Pin Read: ")
		state := pin3.Read()
		if (state == rpio.High) {
			println("High")
		} else {
			println("Low")
		}

		println("changeing pin 7")
		pin7.Toggle()
		time.Sleep(time.Second)
	}



}


func cleanup() {
	rpio.Close()
}