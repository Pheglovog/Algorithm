package main

import (
	"fmt"
	"math"
)

func useMergeSort() {
	s := []int{12, 32, 53, 16, 3, 1, 720, 832, 64, 88}
	mergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}

// MergeSort 是一种分治思想的排序，划分，解决，合并
// 时间复杂度 ： O(nlgn)
// 空间复杂度 ： O(n)
func mergeSort(s []int, stat, end int) {
	if stat < end {
		mid := (stat + end) / 2
		mergeSort(s, stat, mid)
		mergeSort(s, mid+1, end)
		merge(s, stat, mid, end)
	}
}

func merge(s []int, stat, mid, end int) {
	newLeft := make([]int, mid-stat+1)
	copy(newLeft, s[stat:mid+1])
	newLeft = append(newLeft, math.MaxInt)
	newRight := make([]int, end-mid)
	copy(newRight, s[mid+1:end+1])
	newRight = append(newRight, math.MaxInt)

	i, j, k := 0, 0, stat
	for k <= end {
		if newLeft[i] <= newRight[j] {
			s[k] = newLeft[i]
			i++
			k++
		} else {
			s[k] = newRight[j]
			j++
			k++
		}
	}
}
