package main

import "fmt"

/**
Create a list, , of  empty sequences, where each sequence is indexed from  to . The elements within each of the  sequences also use -indexing.
Create an integer, , and initialize it to .
The  types of queries that can be performed on your list of sequences () are described below:
Query: 1 x y
Find the sequence, , at index  in .
Append integer  to sequence .
Query: 2 x y
Find the sequence, , at index  in .
Find the value of element  in  (where  is the size of ) and assign it to .
Print the new value of  on a new line
Task
Given , , and  queries, execute each query.

Note:  is the bitwise XOR operation, which corresponds to the ^ operator in most languages. Learn more about it on Wikipedia.

Input Format

The first line contains two space-separated integers,  (the number of sequences) and  (the number of queries), respectively.
Each of the  subsequent lines contains a query in the format defined above.

Constraints

It is guaranteed that query type  will never query an empty sequence or index.
Output Format

For each type  query, print the updated value of  on a new line.

Sample Input

2 5
1 0 5
1 1 7
1 0 3
2 1 0
2 1 1
Sample Output

7
3
Explanation

Initial Values:


 = [ ]
 = [ ]

Query 0: Append  to sequence .

 = [5]
 = [ ]

Query 1: Append  to sequence .
 = [5]
 = [7]

Query 2: Append  to sequence .

 = [5, 3]
 = [7]

Query 3: Assign the value at index  of sequence  to , print .

 = [5, 3]
 = [7]

7
Query 4: Assign the value at index  of sequence  to , print .

 = [5, 3]
 = [7]

3
 */
func main() {
	var N, Q int64
	fmt.Scanf("%v %v\n", &N, &Q)


	var xy = make([][]int64, Q)
	for q  := 0;int64(q)<Q;q++ {
		xy[q] = make([]int64, 3)
		fmt.Scanf("%v %v %v\n", &xy[q][0],&xy[q][1],&xy[q][2])
	}

	var seq = make([][]int64, N)
	for n:=0; int64(n)<N;n++{
		seq[n] = make([]int64, Q)
	}
	var lastAnswer = int64(0)
	for q := 0;int64(q)<Q;q++ {
		if (xy[q][0] == 1) {
			queryOne(xy[q][1], xy[q][2], seq, &lastAnswer, N)
		} else {
			queryTwo(xy[q][1], xy[q][2], seq, &lastAnswer, N)
		}
	}

}


func queryOne(s int64, y int64,  seq [][]int64, lastAnswer *int64, N int64) {
	var i = (s ^ *lastAnswer) % N
	seq[s][i] = y
}

func queryTwo(s int64, y int64,  seq [][]int64, lastAnswer *int64, N int64) {
	var i = (s ^ *lastAnswer) % N
	var j = i % int64(len(seq[i]))
	*lastAnswer = seq[i][j]
	fmt.Println(*lastAnswer)
}