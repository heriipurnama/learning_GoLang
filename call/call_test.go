package main

import "testing"

func TestDiscountApplied(t *testing.T) {
	call := NewDiscountCalculator(100, 20)
	amount := call.Calculate(150)

	if amount != 10 {
		t.Fail()
	}
}

func main() {

}
