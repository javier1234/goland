package main

import (
	"os"
	"log"
	"fmt"
)

var stdin *os.File

func setStdinFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		panic("no se encontro el archivo en el path")
	}
	stdin = file
}

func main_read_from_file() uint64 {
	var N,M uint64
	fmt.Fscanf(stdin, "%v %v\n",&N,&M)
	var l,r,v, max uint64
	var tree = &Tree{nil, nil, nil, N/2}
	for i:=uint64(0);i<M;i++ {
		fmt.Fscanf(stdin, "%v %v %v\n",&l,&r,&v)
		var interval = &Interval{l, r, v}
		tree = appendInterval(interval, tree)
	}
	//fmt.Println(tree.rightTree)
	queryIntersection(Interval{uint64(5), uint64(5), uint64(200)}, tree)
	//fmt.Printf("%v\n", max)
	return max
}


type Interval struct {
	left uint64
	right uint64
	value uint64
}
func (i Interval) isInternal(x uint64) bool {
	return i.left <= x && x <= i.right
}
func (i Interval) isIntersec(x Interval) bool {
	return x.right >= i.left && x.left <= i.right
}

type Tree struct {
	leftTree, rightTree *Tree
	//TODO deberia estar ordenado por inicio
	intersectionTree []Interval
	center uint64
}
func appendInterval(interval *Interval, t *Tree) *Tree {
	//el intervalo contiene mi centro
	if interval.isInternal(t.center) {
		t.intersectionTree = append(t.intersectionTree, *interval)
	} else if (interval.left < t.center) {
		//es menor va a la izq
		if t.leftTree == nil {
			//TODO crear el slice con el tamaño maximo
			t.leftTree = &Tree{nil, nil, nil, interval.left}
		}
		t.leftTree = appendInterval(interval, t.leftTree)
	} else {
		if t.rightTree == nil {
			//TODO crear el slice con el tamaño maximo
			t.rightTree = &Tree{nil, nil, nil, interval.right}
		}
		t.rightTree = appendInterval(interval, t.rightTree)
	}
	return t
}

func queryIntersection(target Interval, t *Tree)  {
	//fmt.Printf("comparation center t:%d target:%d target:%d\n", t.center,target.left, target.right)
	if target.isInternal(t.center) {
		for _,ci := range t.intersectionTree {
			//fmt.Printf("comparation ci:%d target:%d \n", ci.left,target.left)
			if (ci.isIntersec(target))	{
				fmt.Printf("intersection left:%d right:%d value:%d\n", ci.left,ci.right,ci.value)
			}
		}
	}

	if target.left <= t.center && t.leftTree != nil{
		queryIntersection(target, t.leftTree)
	}

	if target.right >= t.center && t.rightTree != nil{
		queryIntersection(target, t.rightTree)
	}
}