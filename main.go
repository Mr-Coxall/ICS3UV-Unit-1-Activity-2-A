/*
 * Author Mr Coxall
 * Version 1.0.0
 * Date 2025-01-01
 * This is the This program show variable declarations.
 */

package main

import "fmt"

func main() {
	// int data type count
	var count int = 20

	// int64 data type big_number
	var big_number int64 = 14500000000

	// float32 data type temperature
	var temperature float32 = -32.4

	// float64 data type molecules
	var molecules float64 = 158799625.35

	// string data type
	var courseCode string = "ICS3UV"

	// bool data type play_again
	var play_again bool = true

	// display each data type to the screen
	fmt.Println("Can you count up to", count, "?")
	fmt.Println("This is a really big number", big_number)
	fmt.Println("The outside temperature currently is", temperature)
	fmt.Println("A drop of water might have", molecules, "molecules")
	fmt.Println("Do you like", courseCode, "?")
	fmt.Println("True/False Computer programming is cool:", play_again)

	fmt.Println("\nDone.")
}

