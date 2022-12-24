package simple

import (
	"fmt"
	"testing"
)

func TestTwoReturns(t *testing.T) {

	for i := range SumOfSum(10) {
		fmt.Println(i)
	}
}

// 0+1
// 1+2
// 3+3
// 6+4
// 10+5

func TestPassInFunc(t *testing.T) {

	for i := range DoPassInFunc(10, SumOfSum) {
		fmt.Println(i)
	}
}
