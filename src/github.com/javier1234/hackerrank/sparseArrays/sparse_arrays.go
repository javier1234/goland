package main

import "fmt"

/**
There is a collection of  strings ( There can be multiple occurences of a particular string ). Each string's length is no more than  characters. There are also  queries. For each query, you are given a string, and you need to find out how many times this string occurs in the given collection of  strings.

Input Format

The first line contains , the number of strings.
The next  lines each contain a string.
The nd line contains , the number of queries.
The following  lines each contain a query string.

Constraints





Sample Input

4
aba
baba
aba
xzxb
3
aba
xzxb
ab
Sample Output

2
1
0
Explanation

Here, "aba" occurs twice, in the first and third string. The string "xzxb" occurs once in the fourth string, and "ab" does not occur at all.
 */
func main(){
	var nWords uint64
	fmt.Scanln(&nWords)
	//println(nWords)
	words := make(map[string]int, nWords)
	var word string;
	for i:=uint64(0);i<nWords;i++{
		fmt.Scanln(&word)
		words[word] += 1
		//println(words[word])
	}
	var nQueries uint64
	fmt.Scanln(&nQueries)
	var query string
	for i:=uint64(0);i<nQueries;i++{
		fmt.Scanln(&query)
		fmt.Println(words[query])
	}
}
