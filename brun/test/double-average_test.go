package main

import (
	"fmt"
	"testing"

	"imooc.com/resk/infra/algo"
)

func TestDoubleAverage(t *testing.T) {
	count, amount := int64(10), int64(100)
	for i := int64(0); i < count; i++ {
		x := algo.DoubleAverage(count, amount*10)
		fmt.Print(x, ",")
	}

	fmt.Println()
}
