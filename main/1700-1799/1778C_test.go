// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1778/C
// https://codeforces.com/problemset/status/1778/problem/C?friends=on
func Test_cf1778C(t *testing.T) {
	testCases := [][2]string{
		{
			`6
3 1
abc
abd
3 0
abc
abd
3 1
xbb
xcd
4 1
abcd
axcb
3 10
abc
abd
10 3
lkwhbahuqa
qoiujoncjb`,
			`6
3
6
6
6
11`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1778C)
}
