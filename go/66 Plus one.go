package main

import (
	"fmt"
	"math"
)

func plusOne(digits []int) []int {
	sum := 0
	digitsLen := len(digits) - 1
	for _, v := range digits {
		sum += v * int(math.Pow(10.0, float64(digitsLen)))
		digitsLen--
	}
	sum++
	sumStr := fmt.Sprintf("%d", sum)
	ret := make([]int, len(sumStr))
	for i, v := range fmt.Sprintf("%d", sum) {
		ret[i] = int(v) - '0'
	}
	return ret
}

func plusOneV2(digits []int) []int {
	fmt.Printf("%#v\n", digits)
	ret := make([]int, len(digits)+1)
	retLen, digitsLen := len(digits)+1, len(digits)
	for i := len(digits) - 1; i >= 0; i-- {
		cur := digits[i] + 1
		digitsLen--
		retLen--
		if cur < 10 {
			ret[retLen] = cur
			break
		}
		ret[retLen] = 0
		if i == 0 {
			ret[retLen-1] = 1
			retLen--
		}
	}
	for digitsLen > 0 {
		digitsLen--
		retLen--
		ret[retLen] = digits[digitsLen]
	}
	return ret[retLen:]
}

func main() {
	fmt.Printf("%#v\n", plusOneV2([]int{1, 9, 9, 9, 9, 9, 9}))
}
