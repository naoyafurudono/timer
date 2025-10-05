package main

import (
	"time"

	"github.com/naoyafurudono/timer"
)

func main() {
	// Create a new timer
	t := timer.New()

	// Simulate some processing
	t.Lap("Start processing")

	time.Sleep(100 * time.Millisecond)
	t.Lap("First step completed")

	// Simulate a loop with measurements
	for i := range 3 {
		time.Sleep(50 * time.Millisecond)
		t.Lap("Loop iteration %d", i)
	}

	time.Sleep(200 * time.Millisecond)
	t.Lap("Final step completed")

	// Print the results
	t.Print()

	// Print JSON format
	t.PrintJSON()
}
