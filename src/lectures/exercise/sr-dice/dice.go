//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(side int) int {
	return rand.Intn(side) + 1
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	numOfDice, numOfRoll, numOfSide := 2, 2, 6		

	for r := 1; r <= numOfRoll; r++ {
		sum := 0
		for d := 1; d <= numOfDice; d++ {
			rolled := roll(numOfSide)
			sum += rolled
			fmt.Println("Rolls #", r, " Die #", d, ":", rolled)
		}
		fmt.Println("Total rolled:", sum)
		switch sum := sum; {
		case sum == 2 && numOfDice == 2:
			//   - "Snake eyes": when the total roll is 2, and total dice is 2
			fmt.Println("Snake eyes")
		case sum == 7:
			//   - "Lucky 7": when the total roll is 7
			fmt.Println("Lucky 7")
		case sum%2 == 0:
			//   - "Even": when the total roll is even
			fmt.Println("Even")
			//   - "Odd": when the total roll is odd
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}

	// --Requirements:
	// * Print the sum of the dice roll
	// * Print additional information in these cirsumstances:

	//
	// * The program must handle any number of dice, rolls, and sides
	//
	// --Notes:
	// * Use packages from the standard library to complete the project
}
