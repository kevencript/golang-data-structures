package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
 * The Two Pointers is a technique in which we use two pointers to solve a problem.
 * 1. We can use a "start" and "end", or even "left" and "right" pointers.
 * 2. Depeding on the problem, we can iterate the string or array and do our logic
 * 3. This way we can iterate from left to right, or right to left, or even from both sides.
 */
func main() {
	
	/*
	* ARRAY EXAMPLE 1: Return the sum of the numbers in an array
	* 1. We must to return the multiplication of the divisible numebers by 3 & 5 of an array
	* 2. This array can be a range from 1 to 100
	* 3. We must to return a string with the result 
	* 4. Example with the result: "3+5+6+9+10=33"
	*/
	arrayOfRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(returnSum(arrayOfRange))
}

func returnSum(arr []int) string {
	start := 0
	end := len(arr) - 1

	finalResult := 0
	// We can use  Buffer to manipulate strings (WriteString, Write, etc.)
	var buf bytes.Buffer

	// Here we are iterating while the start doesn't reach the end
	for start <= end {
		// Cheching if the number is divisible by 3
		if arr[start]%3 == 0 || arr[start]%5 == 0 {
			// If it is not the first number, we must to add a "+" to the string
			if arr[start] != 3 {
				buf.WriteString("+")
			}

			// Adding the number to the final result
			buf.WriteString(strconv.Itoa(arr[start]))
			finalResult += arr[start]
		}
		start++
	}

	buf.WriteString("=")
	buf.WriteString(strconv.Itoa(finalResult))

	return buf.String()
}
