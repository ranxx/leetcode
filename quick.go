package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(nums)
	quickSortV2(nums, 0, len(nums)-1)
	fmt.Println(nums)
	sort.Ints(nums)

}

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func quickSortV2(values []int, left, right int) {
	if left >= right {
		return
	}
	temp := values[left]
	// p := left
	i, j := left, right

	for i <= j {
		for i <= j && values[j] >= temp {
			j--
		}
		for i <= j && values[i] <= temp {
			i++
		}
		if i <= j {
			values[i], values[j] = values[j], values[i]
		}
	}
	values[j] = temp
	// if p-left > 1 {
	quickSort(values, left, j-1)
	// }
	// if right-p > 1 {
	quickSort(values, j+1, right)
	// }
}
