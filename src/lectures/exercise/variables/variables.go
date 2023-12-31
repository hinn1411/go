//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
//* Store your favorite color in a variable using the `var` keyword
	var favColor string = "blue";
	fmt.Println("my favorite color is", favColor);
//* Store your birth year and age (in years) in two variables using
//  compound assignment
	age, birthYear := 21, "2002";	
	fmt.Println("Born in", birthYear, "aged", age, "years");
//* Store your first & last initials in two variables using block assignment
	var (
		first = "G";
		last = "H";
	)
	fmt.Println("Initials =", first, last);
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
	var declAge int;
	declAge = 365 * age;
	fmt.Println("I am", declAge, "days old");
}
