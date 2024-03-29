package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

func sum(vals ...int) int {
	total := 0
	for _, vals := range vals {
		total += vals
	}
	return total
}
