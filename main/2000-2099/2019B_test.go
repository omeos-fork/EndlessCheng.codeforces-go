// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2019/B
// https://codeforces.com/problemset/status/2019/problem/B?friends=on
func Test_cf2019B(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2 2
101 200
2 1
6 15
1 2 3 5 6 7
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
5 8
254618033 265675151 461318786 557391198 848083778
6 9 15 10 6 9 4 4294967300`,
			`0 100 
0 0 0 0 2 0 0 0 3 0 2 0 0 0 0 
291716045 0 0 0 291716045 0 301749698 0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2019B)
}
