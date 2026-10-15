package main

import "errors"

func Add(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
