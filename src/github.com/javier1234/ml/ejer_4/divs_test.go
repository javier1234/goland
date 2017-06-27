package main

import (
	"testing"
	"fmt"
)


type testpair struct {
	value int
	sum int
}
var tests = []testpair{
	{ 5, 8 },
	{ 6, 14 },
}

func TestSum(t *testing.T) {
	for _, pair := range tests {
		v := Sum(pair.value)
		if v != pair.sum {
			t.Error(
				"For", pair.value,
				"expected", pair.sum,
				"got", v,
			)
		}
	}

}

func BenchmarkSum10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(10)
	}
}

func BenchmarkSum100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(100)
	}
}

//suma de numeros div por 5 y 3
func ExampleSum() {
	//Input: 5
	sum := Sum(5);
	fmt.Println(sum)
	//Output: 8
}
