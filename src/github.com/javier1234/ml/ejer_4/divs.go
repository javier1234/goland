package main

import (
	"fmt"
)

func main(){
	var param int
	fmt.Println("ingresa un numero entero:")
	fmt.Scanln(&param)

	fmt.Println(Sum(param))
}

func Sum(param int) int {
	sum := 0
	for  i := 1;i <= param ;i++  {
		if (i % 3 == 0 || i % 5 == 0) {
			sum += i
		}
	}
	return sum
}