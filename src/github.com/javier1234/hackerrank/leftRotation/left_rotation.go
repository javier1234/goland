package main

import (
	"fmt"
)

/**
A left rotation operation on an array of size  shifts each of the array's elements  unit to the left. For example, if left rotations are performed on array , then the array would become .

Given an array of  integers and a number, , perform  left rotations on the array. Then print the updated array as a single line of space-separated integers.

Input Format

The first line contains two space-separated integers denoting the respective values of  (the number of integers) and  (the number of left rotations you must perform).
The second line contains  space-separated integers describing the respective elements of the array's initial state.

Constraints

Output Format

Print a single line of  space-separated integers denoting the final state of the array after performing  left rotations.

Sample Input

5 4
1 2 3 4 5
Sample Output

5 1 2 3 4
Explanation

When we perform  left rotations, the array undergoes the following sequence of changes:

Thus, we print the array's final state as a single line of space-separated values, which is 5 1 2 3 4.
 */
func main()  {
	var N, LR uint64
	fmt.Scanf("%v %v\n", &N, &LR)

	x := make([]uint64, N)
	y := make([]interface{}, N)
	for i := range x {
		y[i] = &x[i]
	}
	fmt.Scanln(y...)

	var arrLR = make([]uint64, N)
	for i:=uint64(0);i<N;i++ {
		if (i<LR) {
			arrLR[uint64(N-LR+i)] = x[i]
		} else {
			arrLR[uint64(i-LR)] = x[i]
		}
	}
	for i:=uint64(0);i<N;i++ {
		if (i != 0) {
			fmt.Printf(" %v", arrLR[i])
		} else {
			fmt.Printf("%v", arrLR[i])
		}
	}

}
