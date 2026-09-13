package main

import "fmt"

func main() {
	collection1 := []int{9, 7, 4, 1}
	collection2 := []int{0, 2, 5, 8}
	collection3 := []int{0, 3, 6, 10}

	merged := Merge(collection1, collection2, collection3)
	fmt.Println(merged)
}
