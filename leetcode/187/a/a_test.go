// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]`, 
			`"Sao Paulo"`, 
		},
		{
			`[["B","C"],["D","B"],["C","A"]]`, 
			`"A"`, 
		},
		{
			`[["A","Z"]]`, 
			`"Z"`, 
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, destCity, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-187/problems/destination-city/
