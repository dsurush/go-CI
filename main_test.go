package main

import "testing"

func Test_fourthProblem(t *testing.T) {
	sales := []int{5000, 1500, 500, 5000}
	result := fourthProblem(sales)
	want := 5000
	got := result
	if want != result{
	t.Error("nps with args", sales, "want:", want, "got", got)
	}
}