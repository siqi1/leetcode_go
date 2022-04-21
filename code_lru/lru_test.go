package code_lru

import (
	"github.com/siqi1/common/test"
	"testing"
)

func TestLru(t *testing.T) {
	co1 := Constructor(2)
	testCase1 := test.NewTest(t, co1.Get)
	co1.Put(1, 1)
	co1.Put(2, 2)
	testCase1.SetAndRun(1, 1)
	co1.Put(3, 3)
	testCase1.SetAndRun(2, -1)
	co1.Put(4, 4)
	testCase1.SetAndRun(1, -1)
	testCase1.SetAndRun(3, 3)
	testCase1.SetAndRun(4, 4)
}
