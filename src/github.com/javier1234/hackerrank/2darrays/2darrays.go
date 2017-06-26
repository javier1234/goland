package main

import "fmt"

/**
Context
Given a  2D Array, :

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
We define an hourglass in  to be a subset of values with indices falling in this pattern in 's graphical representation:

a b c
  d
e f g
There are  hourglasses in , and an hourglass sum is the sum of an hourglass' values.

Task
Calculate the hourglass sum for every hourglass in , then print the maximum hourglass sum.

Note: If you have already solved the Java domain's Java 2D Array challenge, you may wish to skip this challenge.

Input Format

There are  lines of input, where each line contains  space-separated integers describing 2D Array ; every value in  will be in the inclusive range of  to .

Constraints

Output Format

Print the largest (maximum) hourglass sum found in .

Sample Input

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
Sample Output

19
Explanation

 contains the following hourglasses:

1 1 1   1 1 0   1 0 0   0 0 0
  1       0       0       0
1 1 1   1 1 0   1 0 0   0 0 0

0 1 0   1 0 0   0 0 0   0 0 0
  1       1       0       0
0 0 2   0 2 4   2 4 4   4 4 0

1 1 1   1 1 0   1 0 0   0 0 0
  0       2       4       4
0 0 0   0 0 2   0 2 0   2 0 0

0 0 2   0 2 4   2 4 4   4 4 0
  0       0       2       0
0 0 1   0 1 2   1 2 4   2 4 0
The hourglass with the maximum sum () is:

2 4 4
  2
1 2 4
 */
func main() {
	var xy = [6][6]int{}
	var p = make([][]interface{}, 6)

	for  y := 0;y < 6;y++ {
		p[y] = make([]interface{}, 6)
		for x := 0;x < 6;x++ {
			p[y][x] = &xy[y][x]
		}
	}
	for  y := 0;y < 6;y++ {
		fmt.Scanln(p[y]...)
	}


	var max int
	for  y := 0;y < 4;y++ {
		for  x := 0;x < 4;x++ {
			var sum = xy[y][x]+ xy[y][x+1]+ xy[y][x+2]+ xy[y+1][x+1]+ xy[y+2][x]+ xy[y+2][x+1]+ xy[y+2][x+2]
			if (x==0 && y == 0) {
				max = sum
			} else if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}

