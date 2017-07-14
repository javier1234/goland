package main

import (
	"testing"
)


func Test_mainOld_set2(t *testing.T) {
	setStdinFromFile("test2_imput.txt")
	result := main_local()
	if (result != uint64(200)){
		t.Errorf("el resultado esperado era:%v pero dio:%v", 200, result)
	}
}

func Test_mainOld_set3(t *testing.T) {
	//2490686975
	setStdinFromFile("test3_imput.txt")
	result := main_local()
	if (result != uint64(2490686975)){
		t.Errorf("el resultado esperado era:%v pero dio:%v", 2490686975, result)
	}
}

