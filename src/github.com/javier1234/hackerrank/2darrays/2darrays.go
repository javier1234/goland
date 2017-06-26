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
	var p = [6][6]interface{}{}

	for  y := 0;y < 6;y++ {
		for x := 0;x < 6;x++ {
			p[y][x] = &xy[y][x]
		}
	}
	for  y := 0;y < 6;y++ {
		fmt.Scanf("%v %v %v %v %v %v\n",p[y][0], p[y][1], p[y][2], p[y][3], p[y][4], p[y][5])
	}
	var max = 0
	for  y := 0;y < 4;y++ {
		for  x := 0;x < 4;x++ {
			fmt.Printf("%d %d\n", y, x)
			var xxyy = xy[y:][x:]
			var sum = calcSumHourglasses(xxyy)
			fmt.Printf("sum :%v\n", sum)
			if sum > max {
				max = sum
			}
		}
	}

}

func calcSumHourglasses(xy [][6]int) int {
	fmt.Printf("estoy aca\n")
	return xy[0][0]+xy[0][1]+xy[0][2]+xy[1][1]+xy[2][0]+xy[2][1]+xy[2][2]
}