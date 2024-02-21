// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/466/C
// https://codeforces.com/problemset/status/466/problem/C
func Test_cf466C(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1 2 3 0 3`,
			`2`,
		},
		{
			`4
0 1 -1 0`,
			`1`,
		},
		{
			`2
4 1`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf466C)
}
