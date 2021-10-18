// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][3]string{
		{
			`["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]`,
			`[[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10, 50]]`,
			`[null, true, true, true, false, false]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeClassWithExamples(t, Constructor, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-263/problems/simple-bank-system/
