# Recursion

### Overview: What is recursion?

* In recursion, a function calls itself as a subroutine to solve a problem. 

* Recursion can be used to break down complex problems into smaller, more manageable subproblems. 

* Recursive functions typically have a base case, which is a condition where the function stops calling itself, and one or more recursive cases that lead to the base case.

### Example: Factorial

```python
factorial = n * factorial(n-1)
```

Factorial uses recursion to calculate the factorial of a given number. The factorial of a number n is the product of all positive integers from 1 to n.