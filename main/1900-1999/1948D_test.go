// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1948/D
// https://codeforces.com/problemset/status/1948/problem/D?friends=on
func Test_cf1948D(t *testing.T) {
	testCases := [][2]string{
		{
			`4
zaabaabz
?????
code?????s
codeforces`,
			`6
4
10
0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1948D)
}
