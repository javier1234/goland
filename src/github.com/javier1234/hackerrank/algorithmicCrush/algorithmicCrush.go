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


