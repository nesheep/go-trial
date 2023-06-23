package main

import "fmt"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(min(1, 2))
}
