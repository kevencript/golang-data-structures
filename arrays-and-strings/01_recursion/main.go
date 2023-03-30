package main

import "fmt"

/*
 * RECURSION is:
 * "A function calling itself repeatedly until it stafisfies a base case"
 */

func main() {
	// We call the getFactorial function passing the number 5
	// and print the result
	fmt.Printf("\nFactorial of 5: %d", getFactorial(5))
}

// Algorithm notes for factorial
// 1. Factorial of 0 = 1
// 2. Factorial of n = n * (n-1)!
func getFactorial(n int) int {
	// If the number equals 0, we return 1 (factorial rule)
	if n == 0 || n == 1 {
		return 1
	}

	/* How this works?
	 *
	 * 1. We will cal the function getFactorial(n-1) until n equals 1 or 0
	 * 2. When n equals one, we will return 1 (NOT calling the function again)
	 * 3. When this happens, we will define the (n-1) = 1
	 * 4. This means that the function will start to replace the (n-1) function by function
	 * 5. So this way, it will be something like (when it reaches n=1):
	 * Chronology after reach n=1:
	 *       n * (n - 1) -> we returned 1, so that means that (n - 1) can be replaced by 1
	 * 	  -> 1 * (1 - 1) -> Equals to 0, so we return 1 (factorial of 1)
	 * 	  -> 2 * (2 - 1) -> Equals to 2 (factorial of 2)
	 * 	  -> 3 * (3 - 1) -> Equals to 6 (factorial of 3)
	 * 	  -> 4 * (4 - 1) -> Equals to 24 (factorial of 4)
	 * 	  -> 5 * (5 - 1) -> Equals to 120 (factorial of 5) <- FINAL RESULT
	 * 
	 * Final obs: We can compare it to a stack. After reach the n = 1
	 *  		  we start to replace the (n-1) by 1, and then we start to
	 * 			  multiply the value in a reverse way, passing the value
	 * 			  to the next function call untill the number in which we want.			   		
	 */
	return n * getFactorial(n-1)
}
