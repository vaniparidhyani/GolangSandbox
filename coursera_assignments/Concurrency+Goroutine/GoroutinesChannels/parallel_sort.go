package main

import (
	"fmt"
	"sync"
)

func BubbleSort(sli []int, wg *sync.WaitGroup) {
	defer wg.Done()
	not_sorted := true
	for not_sorted {
		not_sorted = false
		for i := 0; i < len(sli) - 1; i++ {
			if sli[i] > sli[i + 1] {
				not_sorted = true
				sli[i], sli[i + 1] = sli[i + 1], sli[i]
			}
		}
	}
	fmt.Println("Sorted part: ", sli)
}

func Merge(sli []int, first, middle, last int) {
	merged_sli := make([]int, last - first)

	first_ind := first
	second_ind := middle
	merged_ind := 0
	for first_ind < middle && second_ind < last {
		if sli[first_ind] < sli[second_ind] {
			merged_sli[merged_ind] = sli[first_ind]
			first_ind++
		} else {
			merged_sli[merged_ind] = sli[second_ind]
			second_ind++
		}
		merged_ind++
	}
	if first_ind == middle {
		for second_ind < last {
			merged_sli[merged_ind] = sli[second_ind]
			second_ind++
			merged_ind++
		}
	}
	if second_ind == last {
		for first_ind < middle {
			merged_sli[merged_ind] = sli[first_ind]
			first_ind++
			merged_ind++
		}
	}

	for i := 0; i < len(merged_sli); i++ {
		sli[first + i] = merged_sli[i]
	}
}

func main() {
	var array_size int
	var wg sync.WaitGroup
	fmt.Println("Enter array size")
	fmt.Scan(&array_size)

	array := make([]int, array_size)

	fmt.Println("Enter array")
	for i := 0; i < array_size; i++ {
		fmt.Scan(&array[i])
	}

	partition_size := array_size / 4
	wg.Add(4)
	go BubbleSort(array[0:partition_size], &wg)
	go BubbleSort(array[partition_size:2 * partition_size], &wg)
	go BubbleSort(array[2 * partition_size:3 * partition_size], &wg)
	go BubbleSort(array[3 * partition_size:array_size], &wg)
	wg.Wait()
	Merge(array, 0, partition_size, 2 * partition_size)
	Merge(array, 2 * partition_size, 3 * partition_size, array_size)
	Merge(array, 0, 2 * partition_size, array_size)

	fmt.Println("Sorted array: ", array)
}