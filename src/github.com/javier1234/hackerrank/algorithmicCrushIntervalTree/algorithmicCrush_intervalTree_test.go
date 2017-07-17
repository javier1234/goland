package main

import "testing"

func Test_mainOld_set2(t *testing.T) {
	resultExpected := uint64(400)
	setStdinFromFile("test1_imput.txt")
	result := main_read_from_file()
	if (result != resultExpected){
		t.Errorf("el resultado esperado era:%v pero dio:%v", resultExpected, result)
	}
}


func Test_mainOld_set9(t *testing.T) {
	setStdinFromFile("test9_imput.txt")
	result := main_read_from_file()
	if (result != uint64(2501448788)){
		t.Errorf("el resultado esperado era:%v pero dio:%v", 2501448788, result)
	}
}


func Test_mainOld_set3(t *testing.T) {
	resultExpected := uint64(2490686975)
	setStdinFromFile("test3_imput.txt")
	result := main_read_from_file()
	if (result != resultExpected){
		t.Errorf("el resultado esperado era:%v pero dio:%v", resultExpected, result)
	}
}
