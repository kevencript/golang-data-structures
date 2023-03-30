package main

import "fmt"

func main() {
	println(reverseString("gabriel"))
}

func reverseString(s string) string {
	// Creating the two pointers
	head := 0
	tail := len(s) - 1

	// Var to store the reversed rune array string
	reversed := []rune(s)

	for head < tail {
		reversed[head] = rune(s[tail]) // Define the head (index 0) to the tail (index 4)
		reversed[tail] = rune(s[head]) // Define the tail (index 4) to the head (index 0)
		// This way we will always replace the first and last index of the string


		fmt.Println(string(reversed))
		head++
		tail--
	}
	
	return string(reversed)
}