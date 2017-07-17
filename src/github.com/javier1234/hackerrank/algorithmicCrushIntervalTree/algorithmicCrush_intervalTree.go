package main

import (
	"os"
	"log"
	"fmt"
	"sort"
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

func main_read_from_file() (max uint64) {
	var N,M uint64
	fmt.Fscanf(stdin, "%v %v\n",&N,&M)
	var l,r,v uint64
	var tree = &Tree{center:N/2}
	var intervals = make([]Interval, M)
	for i:=uint64(0);i<M;i++ {
		fmt.Fscanf(stdin, "%v %v %v\n",&l,&r,&v)
		var interval = &Interval{l, r, v}
		intervals[i] = *interval
		tree = appendInterval(interval, tree)
	}
	//tree.printCompleteTree()


	for _,interval := range intervals {
		maxLeftAux, maxRightAux := tree.queryMaxByValue(interval)
		//fmt.Printf("max for interval %d/%d is %d\n", interval.left, interval.right, maxAux)
		if (maxLeftAux > max) {
			max = maxLeftAux
		}
		if (maxRightAux > max) {
			max = maxRightAux
		}
	}
	fmt.Printf("%v \n", max)
	return max
}


func main() {
	var N,M uint64
	fmt.Scanf("%v %v\n",&N,&M)
	var l,r,v uint64
	var tree = &Tree{center: N/2}
	var intervals = make([]Interval, M)
	for i:=uint64(0);i<M;i++ {
		fmt.Scanf("%v %v %v\n",&l,&r,&v)
		var interval = &Interval{l, r, v}
		intervals[i] = *interval
		tree = appendInterval(interval, tree)
	}
	var max uint64
	for _,interval := range intervals {
		maxLeftAux, maxRightAux := tree.queryMaxByValue(interval)
		if (maxLeftAux > max) {
			max = maxLeftAux
		}
		if (maxRightAux > max) {
			max = maxRightAux
		}
	}
	fmt.Printf("%v\n", max)
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
	intersectionLeftSortedTree []Interval
	center uint64
}

type IntervalsLeftSorter struct {
	intervals []Interval
	by func(i1, i2 *Interval) bool
}
// Len is part of sort.Interface.
func (s *IntervalsLeftSorter) Len() int {
	return len(s.intervals)
}

// Swap is part of sort.Interface.
func (s *IntervalsLeftSorter) Swap(i, j int) {
	s.intervals[i], s.intervals[j] = s.intervals[j], s.intervals[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *IntervalsLeftSorter) Less(i, j int) bool {
	return s.by(&s.intervals[i], &s.intervals[j])
}
func byleft (i1, i2 *Interval) bool {
	return i1.left < i2.left
}


func appendInterval(interval *Interval, t *Tree) *Tree {
	//el intervalo contiene mi centro
	if interval.isInternal(t.center) {
		t.intersectionLeftSortedTree = append(t.intersectionLeftSortedTree, *interval)
		sort.Sort(&IntervalsLeftSorter{t.intersectionLeftSortedTree, byleft})
	} else if (interval.left <= t.center) {
		//es menor va a la izq
		if t.leftTree == nil {
			t.leftTree = &Tree{center: interval.left}
		}
		t.leftTree = appendInterval(interval, t.leftTree)
	} else {
		if t.rightTree == nil {
			t.rightTree = &Tree{center: interval.right}
		}
		t.rightTree = appendInterval(interval, t.rightTree)
	}
	return t
}

func (t *Tree) queryIntersection(target Interval)  {
	fmt.Printf("comparation center t:%d target:%d target:%d\n", t.center,target.left, target.right)
	if target.isInternal(t.center) {
		for _,ci := range t.intersectionLeftSortedTree {
			fmt.Printf("comparation ci:%d target:%d \n", ci.left,target.left)
			if (ci.isIntersec(target))	{
				fmt.Printf("intersection left:%d right:%d value:%d\n", ci.left,ci.right,ci.value)
			} else if (target.right > ci.left) {
				fmt.Printf("salgo por la lista ordenada:t.right=%d ci.lefr=%d", target.right, ci.left)
				break;
			}
		}
	}

	if target.left < t.center && t.leftTree != nil{
		t.leftTree.queryIntersection(target)
	}

	if target.right > t.center && t.rightTree != nil{
		t.rightTree.queryIntersection(target)
	}
}


func (t *Tree) queryMaxByValue(target Interval) (maxLeft, maxRight uint64)  {
	//fmt.Printf("comparation center t:%d target:%d target:%d\n", t.center,target.left, target.right)
	for _,ci := range t.intersectionLeftSortedTree {
		//fmt.Printf("comparation ci:%d target:%d \n", ci.left,target.left)
		if (ci.isIntersec(target))	{
			//fmt.Printf("intersection left:%d right:%d value:%d\n", ci.left,ci.right,ci.value)
			if (ci.isInternal(target.left)) {
				//fmt.Printf("left %d+%d\n", maxLeft, target.value)
				maxLeft += ci.value
			}
			if (ci.isInternal(target.right)) {
				//fmt.Printf("right %d+%d\n", maxRight, target.value)
				maxRight += ci.value
			}
		} else if (target.right < ci.left) {
			//fmt.Printf("salgo por la lista ordenada:t.right=%d ci.lefr=%d", target.right, ci.left)
			break;
		}
	}

	var maxLeftAux,maxRightAux uint64
	if target.left <= t.center && t.leftTree != nil{
		//fmt.Printf("izquierda \n")
		maxLeftAux,maxRightAux = t.leftTree.queryMaxByValue(target)
		maxLeft += maxLeftAux
		maxRight += maxRightAux
	}

	if target.right >= t.center && t.rightTree != nil{
		//fmt.Printf("dere \n")
		maxLeftAux,maxRightAux = t.rightTree.queryMaxByValue(target)
		maxLeft += maxLeftAux
		maxRight += maxRightAux
	}
	return maxLeft,maxRight
}

func (t *Tree) printCompleteTree()  {
	fmt.Printf("Center:%d \n", t.center)
	for _,ci := range t.intersectionLeftSortedTree {
		fmt.Printf("    %d-%d-%d\n", ci.left,ci.right,ci.value)
	}

	if t.leftTree != nil{
		fmt.Println("inz")
		t.leftTree.printCompleteTree()
	}

	if t.rightTree != nil{
		fmt.Println("der")
		t.rightTree.printCompleteTree()
	}
}