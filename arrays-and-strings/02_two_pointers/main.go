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
	* Example 1: Return the sum of the numbers in an array
	* 1. We must to return the multiplication of the divisible numebers by 3 & 5 of an array
	* 2. This array can be a range from 1 to 100
	* 3. We must to return a string with the result 
	* 4. Example with the result: "3+5+6+9+10=33"
	*/
	arrayOfRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(returnSum(arrayOfRange))

	/*
	 * Example 2: Return true if a given string is a palindrome, false otherwise.
	 * 1. A string is a palindrome if it reads the same forwards as backwards.
	 * 2. That means, after reversing it, it is still the same string. 
	 * 3. For example: "aba", or "101".	
	 */
	fmt.Printf("\nIs Palindrome: %v", isPalindrome("csyyysc"))

	/*
	* Example 3: Given a sorted array of unique integers and a target integer, return true 
	* if there exists a pair of numbers that sum to target, false otherwise. 
	* This problem is similar to Two Sum.
	*
	* For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true 
	* because 4 + 9 = 13.
	*/
	fmt.Printf("\nPair Sum Array: %v\n", pairSumArray([]int{1, 2, 4, 6, 8, 9, 14, 15}, 13))

	/*
	 * Example 3: Given two sorted integer arrays, 
	 * return an array that combines both of them and is also sorted.
	 * 
	 *  Ex: Array1: [1, 3, 4, 5] Array2: [2, 6, 7, 8]
	 */
	fmt.Printf("\nMerge Sorted Arrays: %v", mergeAndSortTwoArrays([]int{1, 3, 4, 5}, []int{2, 6, 7, 8}))
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


func isPalindrome(s string) bool {
	// We define a start and end pointer
	start := 0
	end := len(s) - 1

	for start <= end {

		/*
		 * Since that the first and last character must be the same
		 * for a string being a palindrome, we use the start and end
		 * pointers to compare the characters.
		 *
		 */
		if s[start] != s[end] { // Middlechar both Start: x, End: x are equals
			return false
		}
		fmt.Printf("\nStart: %v, End: %v", string(s[start]), string(s[end]))
		start++
		end--
	}

	fmt.Printf("\nMiddle char: %v", string(s[(len(s) - 1) / 2]))

	return true
}

func pairSumArray(arr []int, target int) bool {
	start := 0
	end := len(arr) - 1
	
	/*
	 * Here we want to find a pair of numbers that sum to target.
	 *
	 * 1. We want to check if the Sum is equals to the target
	 * 2. If it is, we will return true
	 * 3. If not, we will check if the Sum is greater than the target
	 * 	3.1 Sum > Target: we will decrease the END pointer
	 * 	3.2 Sum < Target: we will increase the start
	 * 4. Visual example:	
	 *		[*1,  2,  4, 6, 8,  9,  14, *15] (1+15=16, greater than 13. Decrease the end pointer)
	 *		[*1,  2,  4, 6, 8,  9, *14, XX]  (1+14=15, greater than 13. Decrease the end pointer)
	 *		[*1,  2,  4, 6, 8, *9,  XX, XX]  (1+9=10, minor than 13. Increase the start pointer)
	 *		[XX, *2,  4, 6, 8, *9,  XX, XX]  (1+9=10, minor than 13. Increase the start pointer)
	 *		[XX, *2,  4, 6, 8, *9,  XX, XX]  (2+9=11, minor than 13. Increase the start pointer)
	 *		[XX, XX, *4, 6, 8, *9,  XX, XX]  (4+9=13, equals to 13. Return true)
	 */
	for start <= end {
		// Sum of the pair
		pairSum := arr[start] + arr[end]

		// Checking if the sum is equals to the target
		if pairSum == target {
			fmt.Printf("\n%v + %v = %v (Target: %v)", arr[start], arr[end], pairSum, target)
			return true
		}

		// Checking if the sum is greater than the target
		if pairSum > target {
			end-- // If yes, we will move the end pointer to the left
		} else {
			start++ // If not, we will move the start pointer to the right
		}

		fmt.Printf("\n%v + %v = %v (Target: %v)\n", arr[start], arr[end], pairSum, target)
	}

	return false
}

func mergeAndSortTwoArrays(firstArr []int, secondArr []int) []int {
	i,j := 0,0
	sortedArr := []int{}

	/*
	 * For this example we will use the TwoPointers to iterate
	 * through the arrays. This means that one pointer will be for firstArray, 
	 * and the other for secondArray.
	 *
	 * 1. We will iterate i && j compared to the length of the arrays
	 * 2. If the first array value on i positio is less than the second array value on j position:
	 *	  we will increase the i pointer and add the value to the sorted array
	 * 3. If the second array value on j position is less than the first array value on i position:
	 *	  We will increase the j pointer and add the value to the sorted array
	 * 4. When i or j satisfy the for loop, usually we have remaining values into the other array.
	 * 	  For this reason, we must to validate the length and add the remaining values to the array.
	 *    obs: we do not need to compare, cause one of the arrays is done already. What remain is just the other numbers
	 *
	 */
	for i < len(firstArr) && j < len(secondArr)	{
		// If the first array value on i position is less than the second array value on j position
		if firstArr[i] < secondArr[j] {
			sortedArr = append(sortedArr, firstArr[i])
			i++
		} else {
			// If the second array have the minor value, we will append it to the array
			// and increase the j poiter
			sortedArr = append(sortedArr, secondArr[j])
			j++
		}

		for i < len(firstArr) {
			sortedArr = append(sortedArr, firstArr[i])
			i++
		}

		for j < len(secondArr) {
			sortedArr = append(sortedArr, secondArr[j])
			j++
		}

	
		// Returning a visual representation of how the values are changing
		fmt.Println(sortedArr, firstArr[i:], secondArr[j:])
	}

	return sortedArr
}