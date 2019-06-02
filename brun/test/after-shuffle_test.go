package main

import (
	"testing"
	"imooc.com/resk/infra/algo"
)

func TestAfterShuffle(t *testing.T) {
	t.Log(algo.AfterShuffle(int64(10), int64(100*100)))
}
