/*
 * Author Mr Coxall
 * Version 1.0.0
 * Date 2025-01-01
 * This program This program sums 3 constants.
 */

package main

import "fmt"

func main() {
	// initialize numbers as constants
	const NUMBER1 int = 10
	const NUMBER2 int = -20
	const NUMBER3 int = 85

	// INPUT - none

	// PROCESS
	// calculate the sum
	answer := NUMBER1 + NUMBER2 + NUMBER3

	// OUTPUT
	// display the result
	fmt.Println("The sum of", NUMBER1, "&", NUMBER2, "&", NUMBER3, "is:", answer)

	fmt.Println("\nDone.")
}
