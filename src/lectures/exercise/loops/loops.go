//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//--Requirements:
	for i := 1; i <= 50; i++ {
		div3 := i%3 == 0
		div5 := i%5 == 0
		if div3 && div5 {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FizzBuzz")
		} else if div3 {
			//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println("Fizz")
		} else if div5 {
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("Fizz")
		} else {
			//* Print integers 1 to 50, except:
			fmt.Println(i)
		}
	}

}
