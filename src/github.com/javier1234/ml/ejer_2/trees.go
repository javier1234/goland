package main

import (
	"strconv"
	"os"
	"fmt"
)

type tree struct {
	value int
	leftTree *tree
	rightTree *tree
}


func append(t *tree , value int ) *tree {
	if (t == nil) {
		return &tree{value, nil, nil}
	}
	//evaluamos a que rama va
	if (value > t.value) {
		t.rightTree = append(t.rightTree, value)
	} else {
		t.leftTree = append(t.leftTree, value)
	}
	return t
}

func printPreOrderTree(t *tree) {
	if (t == nil) {
		return
	}
	printPreOrderTree(t.leftTree)
	fmt.Printf("%v ", t.value)
	printPreOrderTree(t.rightTree)
}

func parseArgs(argsCommandLine []string) []int {
	argSlice := make([]int, len(argsCommandLine)-1)
	for i := 1; i < len(argsCommandLine); i++ {
		argSlice[i-1], _ = strconv.Atoi(argsCommandLine[i])
	}
	return argSlice
}


/**
Hacer un programa que reciba por parametro una lista de enteros, los almacene en un slice,
popule un arbol binario y despues imprima los valores en orden ascendente
 */
func main() {
	arg := parseArgs(os.Args)
	fmt.Println(arg)
	var myTree *tree
	for _, value := range arg {
		myTree = append(myTree, value);
	}
	println("el arbol ordenado:")
	printPreOrderTree(myTree)
}

