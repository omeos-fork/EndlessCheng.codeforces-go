package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1 2
2 2
1 3
3 -2
3 3`,
			`8`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
