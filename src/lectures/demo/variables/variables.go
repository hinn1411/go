package main

import "fmt"

func main() {

	var myName = "Hien";
	fmt.Println("my name is", myName, myName);

	var name string = "Tuan Hien";
	fmt.Println("my =", name);

	userName := "Giang Tuan Hien"; // decl & assign syntax
	fmt.Println("username =", userName);

	var sum int; // init val = 0, "" for string
	fmt.Println("The sum is", sum);

	part1, other := 1, 5;
	fmt.Println("part 1 is", part1, "other is", other);

	part2, other := 2, 0; // Commna OK
	fmt.Println("part 2 is", part2, "other is", other);

	sum = part1 + part2;
	fmt.Println("sum =", sum);

	var (
		lessonName = "Variables";
		lessonType = "Demo";
	)

	fmt.Println("lessonName =", lessonName);
	fmt.Println("lessonType =", lessonType);

	word1, word2, _ := "hello", "word", "!";
	fmt.Println(word1, word2);
}
