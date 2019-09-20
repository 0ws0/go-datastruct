package main

import (
	"fmt"
)

var sortArr []int = []int{1,10,22,23,8,9,5}

func main() {
	sort("merge")
}

func sort(sortStr string) {
	var result []int

	switch sortStr {
	case "bubble":
		result = bubbleSort(sortArr)
	case "choose":
		result = chooseSort(sortArr)
	case "insert":
		result = insertSort(sortArr)
	case "quick":
		result = quickSort(sortArr)
	case "shell":
		result = shellSort(sortArr)
	case "merge":
		result = mergeSort(sortArr)
	default:
	}

	fmt.Println(result)
}
