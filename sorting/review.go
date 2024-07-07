package main

import (
	"fmt"
	"math"
)

func review() {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	reviewSort(s, 0, len(s)-1)
	fmt.Println(s)
}

func reviewSort(s []int, stat, end int) {
	if stat < end {
		mid := (stat + end) / 2
		reviewSort(s, stat, mid)
		reviewSort(s, mid+1, end)
		rmerge(s, stat, mid, end)
	}
}

func rmerge(s []int, stat, mid, end int) {
	left := make([]int, mid-stat+1)
	copy(left, s[stat:mid+1])
	left = append(left, math.MaxInt)
	right := make([]int, end-mid)
	copy(right, s[mid+1:end+1])
	right = append(right, math.MaxInt)

	for i, j, k := 0, 0, stat; k <= end; k++ {
		if left[i] <= right[j] {
			s[k] = left[i]
			i++
		} else {
			s[k] = right[j]
			j++
		}
	}
}
