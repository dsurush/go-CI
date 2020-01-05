package main

import (
	"fmt"
)

func firstProblem(sales []int){
	sum := 0
	for _, value := range sales{
		sum = sum + value
	}
	fmt.Println("Sum of rows", sum)
}

func secondProblem(score []int){
	sum := 0
	for _, value := range score{
		sum = sum + value
	}
	total := 1.0 * sum / len(score)
	fmt.Println("middle mark = ", total)
}


func thirdProblem(score []int) int {
	max := 0
	//var index = 0
	for i:=0; i<len(score); i++{
		if max <= score[i]{
			max = score[i]
		//	index = i
		}
	}
	//fmt.Println("Max of mark = ", max, "Index of max marks = ", index)
	return max
}

func fourthProblem(score []int) int{
	min := 2_000_000_000
	//id := 0
	for _, value := range score{
		if min >= value {
			min = value
	//		id = index
		}
	}
	//fmt.Println("Min of mark = ", min, "Index of min marks = ", id)
	return min
}


func main(){
//	sales := []int{1000, 2000, 500, 5000}
//	firstProblem(sales)

//	score := []int{8, 7, 4, 2, 10}
//	secondProblem(score)

	//score = []int{8, 7, 4, 2, 10}
//	thirdProblem(score)

//	fourthProblem(score)


}
