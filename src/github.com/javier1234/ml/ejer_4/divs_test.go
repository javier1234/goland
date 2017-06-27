package main

import "testing"


type testpair struct {
	value int
	sum int
}
var tests = []testpair{
	{ 5, 5 },
	{ 6, 9 },
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
