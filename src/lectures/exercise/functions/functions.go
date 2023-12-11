//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet (name string) {
	fmt.Println("Hello ", name);
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func linda () string {
	return "Á, chết chị rồi em ơi!";
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add(a, b, c int) int {
	return a + b + c;
}
//* Write a function that returns any number
func returnNum () {
	return 1;
}
//* Write a function that returns any two numbers
func returnTwoNum () {
	return 1, 2;
}

func main() {
	fmt.Println(linda());




//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once
}
