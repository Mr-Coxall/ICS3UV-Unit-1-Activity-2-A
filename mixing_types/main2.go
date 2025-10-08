/*
 * Author Mr Coxall
 * Version 1.0.0
 * Date 2025-01-01
 * This program shows mixing types.
 */

package main

import "fmt"

func main() {
	// initialize numbers as constants
	var someNumber int = 25
	var someFloat float32 = 85.75

	// INPUT - none

	// PROCESS
	// calculate the sum
	answer := someNumber + someFloat

	// OUTPUT
	// display the result
	fmt.Println("The sum of", someNumber, "&", someFloat, "is:", answer)

	fmt.Println("\nDone.")
}

