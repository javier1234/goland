package main

import (
	"testing"
)

const path  = "/Users/jrobles/git/goland/src/github.com/javier1234/hackerrank/algorithmicCrush/"

func Test_mainOld_set2(t *testing.T) {
	setStdinFromFile(path + "test2_imput.txt")
	result := main_local()
	if (result != uint64(1500)){
		t.Errorf("el resultado esperado era:%v pero dio:%v", 1500, result)
	}
}

func Test_mainOld_set3(t *testing.T) {
	//2490686975
	setStdinFromFile(path + "test3_imput.txt")
	result := main_local()
	if (result != uint64(2490686975)){
		t.Errorf("el resultado esperado era:%v pero dio:%v", 2490686975, result)
	}
}

