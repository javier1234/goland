package main

import "fmt"

/**
Russian

You are given a list of size , initialized with zeroes. You have to perform  operations on the list and output the maximum of final values of all the  elements in the list. For every operation, you are given three integers ,  and  and you have to add value  to all the elements ranging from index  to (both inclusive).

Input Format

First line will contain two integers  and  separated by a single space.
Next  lines will contain three integers ,  and  separated by a single space.
Numbers in list are numbered from  to .

Constraints






Output Format

A single line containing maximum value in the updated list.

Sample Input

5 3
1 2 100
2 5 100
3 4 100
Sample Output

200
Explanation

After first update list will be 100 100 0 0 0.
After second update list will be 100 200 100 100 100.
After third update list will be 100 200 200 200 100.
So the required answer will be 200.
 */
func main(){
	var N,M uint64
	fmt.Scanf("%v %v\n",&N,&M)
	var arr = make([]uint64, N)
	var x,y,v,max uint64
	for i:=uint64(0);i<M;i++ {
		fmt.Scanf("%v %v %v\n",&x,&y,&v)
		if (x == y) {
			arr[x-1] += v
			if (max < arr[x-1]) {
				max = arr[x-1]
			}
		} else {
			for j:=x-1;j<y;j++ {
				arr[j] += v
				if (max < arr[j]) {
					max = arr[j]
				}
			}
		}
		//fmt.Println(arr)
	}
	fmt.Println(max)
}
