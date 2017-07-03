package main

import (
	"fmt"
	"os"
	"log"
)


// So stdin can be mocked during testing.
var stdin *os.File

func setStdinFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		panic("no se encontro el archivo en el path")
	}
	stdin = file
}



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

NO PERFORMA, NO FUNCIONA CON ENTRADAS GRANDER

 */
func main_old(){
	var N,M uint64
	fmt.Fscanf(stdin, "%v %v\n",&N,&M)
	var arr = make([]uint64, N)
	var x,y,v,max uint64
	for i:=uint64(0);i<M;i++ {
		fmt.Fscanf(stdin, "%v %v %v\n",&x,&y,&v)
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
	}
	fmt.Println(max)
}


func main_local() uint64 {
	var N,M uint64
	fmt.Fscanf(stdin, "%v %v\n",&N,&M)
	var queqe *Node
	queqe = &Node{nil,nil,uint64(0),uint64(0)}

	var l,r,v, max uint64
	for i:=uint64(0);i<M;i++ {
		fmt.Fscanf(stdin, "%v %v %v\n",&l,&r,&v)
		//fmt.Printf("entrada: %v %v %v\n", l,r,v)
		leftNode := addNode(queqe, queqe.next, l, v)
		rightNode := addNode(leftNode, leftNode.next, r, v)
		//fmt.Printf("izq:%v der:%v\n",leftNode.key, rightNode.key)
		max = calculateMax(leftNode, rightNode, v, max)
	}
	for ;queqe != nil;{
		fmt.Printf("key:%v value:%v\n", queqe.key, queqe.value)
		queqe = queqe.next
	}
	fmt.Printf("%v\n", max)
	return max
}

type Node struct {
	before *Node
	next *Node
	key uint64
	value uint64
}

func main(){
	var N,M uint64
	fmt.Scanf("%v %v\n",&N,&M)
	var queqe *Node
	queqe = &Node{nil,nil,uint64(0),uint64(0)}

	var l,r,v, max uint64
	for i:=uint64(0);i<M;i++ {
		fmt.Scanf("%v %v %v\n",&l,&r,&v)
		leftNode := addNode(queqe, queqe.next, l, v)
		rightNode := addNode(leftNode, leftNode.next, r, v)
		max = calculateMax(leftNode, rightNode, v, max)
	}
	fmt.Printf("%v\n", max)
}

func calculateMax(left *Node, right *Node, value uint64, max uint64) uint64{
	queqe := left.next
	for ;queqe!=right.before && queqe != nil;{
		queqe.value += value
		if (max < queqe.value) {
			max = queqe.value
			//fmt.Printf("nuevo max:%v\n", *max)
		}
		queqe = queqe.next
	}
	return max
}






func addNode(nodo *Node, nodoNext *Node, k uint64, v uint64) (*Node) {
	if (nodoNext == nil) {
		if (nodo.key < k) {
			var newNode = &Node{nodo, nil, k, v}
			if (nodo.before != nil) {
				nodo.before.next = nodo
			}
			nodo.next = newNode
			return newNode
		} else {
			var newNode = &Node{nodo.before, nodo, k, v}
			if (nodo.before != nil) {
				nodo.before.next = newNode
			}
			nodo.before = newNode
			return newNode
		}
	} else {
		if (nodo.key < k) {
			var newNode = addNode(nodo.next, nodo.next.next, k, v)
			return newNode
		} else {
			var newNode = &Node{nodo.before, nodo, k, v}
			if (nodo.before != nil) {
				nodo.before.next = newNode
			}
			return newNode
		}
	}
}


