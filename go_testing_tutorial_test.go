package main

import "testing"

func TestCalculate(t *testing.T){
	if Calculate(2) != 4{
		t.Error("2 + 2 = 4")
	}
}

