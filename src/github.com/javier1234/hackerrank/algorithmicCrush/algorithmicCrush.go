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
	var intervals *Tree
	fmt.Fscanf(stdin, "%v %v\n",&N,&M)
	var l,r,v, max uint64
	for i:=uint64(0);i<M;i++ {
		fmt.Fscanf(stdin, "%v %v %v\n",&l,&r,&v)
		intervals = appendInterval(l, r, v, intervals)
	}
	fmt.Printf("%v\n", max)
	return max
}


func appendInterval(t *Tree , l uint64, r uint64, value uint64 ) *Tree {
	if (t == nil) {
		return &Tree{l, r, value, nil, nil}
	}
	//evaluamos a que rama va
	if (l > t.left) {
		t.rightLeaf = append(t.rightLeaf, l, r, value)
	} else {
		t.leftLeaf = append(t.leftLeaf, l, r, value)
	}
	return t
}

func main_arbolbinario() uint64 {
	var N,M uint64
	fmt.Fscanf(stdin, "%v %v\n",&N,&M)
	var treeIntervalLeft *Tree
	var treeIntervalRight *Tree
	var arrIntervalLeft = make([]uint64, M)
	var arrIntervalRight = make([]uint64, M)
	var l,r,v, max uint64
	for i:=uint64(0);i<M;i++ {
		fmt.Fscanf(stdin, "%v %v %v\n",&l,&r,&v)
		treeIntervalLeft = append(treeIntervalLeft, l, r, v)
		treeIntervalRight = append(treeIntervalRight, r, l,  v)
		arrIntervalLeft[i] = l
		arrIntervalRight[i] = r
	}
	fmt.Printf("%v\n", max)

	for i:=uint64(0);i<M;i++ {
		temp := calculateMaxLeftSize(treeIntervalLeft, arrIntervalLeft[i])
		//fmt.Printf("calculado left:%v\n", temp)
		if (max < temp) {
			max = temp
		}
	}
	/*
	for i:=uint64(0);i<M;i++ {
		temp := calculateMaxLeftSize(treeIntervalRight, arrIntervalRight[i])
		//fmt.Printf("calculado right:%v\n", temp)
		if (max < temp) {
			max = temp
		}
	}*/
	return max
}



type Tree struct {
	left      uint64
	right     uint64
	value     uint64
	leftLeaf  *Tree
	rightLeaf *Tree
}



func append(t *Tree , key uint64, key2 uint64, value uint64 ) *Tree {
	if (t == nil) {
		return &Tree{key, key2, value, nil, nil}
	}
	//evaluamos a que rama va
	if (key > t.left) {
		t.rightLeaf = append(t.rightLeaf, key, key2, value)
	} else {
		t.leftLeaf = append(t.leftLeaf, key,key2, value)
	}
	return t
}

func calculateMaxLeftSize(leftTree *Tree, interval uint64) (max uint64) {
	if (leftTree == nil) {
		return uint64(0)
	}
	max = calculateMaxLeftSize(leftTree.leftLeaf, interval)
	if (leftTree.left <=interval) {
		if  (leftTree.right >= interval) {
			max += leftTree.value
		}
		max += calculateMaxLeftSize(leftTree.rightLeaf, interval)
	}
	return max
}

func calculateMaxRightSize(tree *Tree, interval uint64) (max uint64) {
	if (tree == nil) {
		return uint64(0)
	}
	max = calculateMaxRightSize(tree.rightLeaf, interval)
	if (tree.left >=interval) {
		if (tree.right <= interval) {
			max += tree.value
		}
		max += calculateMaxRightSize(tree.leftLeaf, interval)
	}
	return max
}

func printPreOrderTree(t *Tree) {
	if (t == nil) {
		return
	}
	printPreOrderTree(t.leftLeaf)
	fmt.Printf("key:%v value:%v\n ", t.left, t.value)
	printPreOrderTree(t.rightLeaf)
}


/* -----------------------
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
----------------------------------------------------------------------------------- */
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


