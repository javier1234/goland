package main

import "fmt"

/**
Given an array of integers, calculate which fraction of its elements are positive, which fraction of its elements are negative, and which fraction of its elements are zeroes, respectively. Print the decimal value of each fraction on a new line.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  are acceptable.

Input Format

The first line contains an integer, , denoting the size of the array.
The second line contains  space-separated integers describing an array of numbers .

Output Format

You must print the following  lines:

A decimal representing of the fraction of positive numbers in the array compared to its size.
A decimal representing of the fraction of negative numbers in the array compared to its size.
A decimal representing of the fraction of zeroes in the array compared to its size.
Sample Input

6
-4 3 -9 0 4 1
Sample Output

0.500000
0.333333
0.166667
Explanation

There are  positive numbers,  negative numbers, and  zero in the array.
The respective fractions of positive numbers, negative numbers and zeroes are , and , respectively.
 */
func main() {
	var N int
	var pos int
	var neg int
	var zeros int

	fmt.Scanf("%v", &N)
	x := make([]int, N)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	fmt.Scanln(y...)
	for _, val := range x {
		if val == 0 {
			zeros += 1
		} else if val > 0 {
			pos += 1
		} else {
			neg += 1
		}
	}
	fmt.Println(float64(pos) / float64(N))
	fmt.Println(float64(neg) / float64(N))
	fmt.Println(float64(zeros) / float64(N))
}