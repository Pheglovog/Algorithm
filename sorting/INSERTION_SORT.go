package main

import "fmt"

func useInsertSort() {
	s := []int{12, 32, 53, 16, 3, 1, 720, 832, 64, 88}
	insertionSort(s)
	fmt.Println(s)
}

/**
* 插入排序： 左手没牌， 拿第一张在左手，之后依次拿出之后的牌，找合适位置放入
* 时间复杂度：O(n2)
* 空间复杂度：O(1)
 */
func insertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		key := s[i]
		j := i - 1
		for j >= 0 && s[j] > key {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = key
	}
}
