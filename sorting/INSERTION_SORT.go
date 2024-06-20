package main

import "fmt"

func main() {
	s := []int{12, 32, 53, 16, 3, 1, 720, 832, 64, 88}
	insertionSort(s)
	fmt.Println(s)
}

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
