package main

import "fmt"

func mapp(a []int, f func(int) int) []int {
	for i := 0; i < len(a); i++ {
		a[i] = f(a[i])
	}
	return a
}

func main() {
	a := []int{1, 2, 3}
	b := mapp(a, func(i int) int { return i + 1 })
	fmt.Println(b)
}
