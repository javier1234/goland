package main

import "fmt"

/**
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

Input Format

A single line of five space-separated integers.

Constraints

Each integer is in the inclusive range .
Output Format

Print two space-separated long integers denoting the respective minimum and maximum values that can be calculated by summing exactly four of the five integers. (The output can be greater than 32 bit integer.)

Sample Input

1 2 3 4 5
Sample Output

10 14
Explanation

Our initial numbers are , , , , and . We can calculate the following sums using four of the five integers:

If we sum everything except , our sum is .
If we sum everything except , our sum is .
If we sum everything except , our sum is .
If we sum everything except , our sum is .
If we sum everything except , our sum is .
As you can see, the minimal sum is  and the maximal sum is . Thus, we print these minimal and maximal sums as two space-separated integers on a new line.

Hints: Beware of integer overflow! Use 64-bit Integer.
 */
func main() {
	var max, min, allSum uint64
	x := make([]uint64, 5)
	fmt.Scanf("%v %v %v %v %v\n", &x[0], &x[1], &x[2], &x[3], &x[4])
	allSum = x[0] + x[1] + x[2] + x[3] + x[4]
	min = allSum
	max = 0

	for i:=0;i<5;i++{
		nsum := allSum - x[i]
		if max < nsum {
			max = nsum
		}
		if min > nsum {
			min = nsum
		}
	}
	fmt.Printf("%v %v", min, max)
}
