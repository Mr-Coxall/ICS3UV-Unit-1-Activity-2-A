/*
 * Author Mr Coxall
 * Version 1.0.0
 * Date 2025-01-01
 * This program shows mixing types.
 */

package main

import "fmt"

func main() {
	// initialize variables
	var someInt int = 25
	var someFloat float64 = 85.75


	// this will hold the sum of the two values
	var sum int

	// add an int and a float64, casting the float64 to an int
	sum = someInt + int(someFloat)
	// the int() conversion truncates the decimal part of someFloat

	// output the result
	fmt.Println("The sum of a type casted float and an int is:", sum)

	fmt.Println("\nDone.")
}
