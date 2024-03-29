package memo_test

import (
	"testing"

	"gopl.io/ch9/memo4"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func start(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
