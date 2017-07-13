package main

import "testing"

func Test_mainOld_set2(t *testing.T) {
	setStdinFromFile("test1_imput.txt")
	main_read_from_file()
	/*
	result := main_read_from_file()
	if (result != uint64(200)){
		t.Errorf("el resultado esperado era:%v pero dio:%v", 1500, result)
	}
	*/
}
