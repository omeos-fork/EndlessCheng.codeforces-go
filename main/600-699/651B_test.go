// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/651/B
// https://codeforces.com/problemset/status/651/problem/B?friends=on
func Test_cf651B(t *testing.T) {
	testCases := [][2]string{
		{
			`5
20 30 10 50 40`,
			`4`,
		},
		{
			`4
200 100 100 200`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf651B)
}
