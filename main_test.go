package main

import "testing"

func Test_fourthProblem(t *testing.T) {
	testCases := []struct{
		name string
		score []int
		want int

	}{
		{ name: "All promoters", score:[]int{10, 10, 10}, want:10},
		{name: "All detractors", score:[]int{0, 80, 0}, want: 0},
		{name: "All neutrals", score:[]int{8, 58, 8}, want: 8},
		{name: "Mixed", score:[]int{10, 10, 8}, want: 8},
	}

	for _, testCases := range testCases{
		got := fourthProblem(testCases.score)
		if got != testCases.want{
			t.Error("fourthProblem test", testCases.name, "got", got, "want", testCases.want)
		}
	}
	/*	sales := []int{5000, 1500, 500, 5000}
	result := fourthProblem(sales)
	want := 5000
	got := result
	if want != result{
	t.Error("nps with args", sales, "want:", want, "got", got)
	}*/
}