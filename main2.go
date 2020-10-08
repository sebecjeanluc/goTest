package main

import "fmt"

func todo() {
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}

	arr2 = append(arr2, "is tak", "yo")
	fmt.Println(arr, arr2)
}

func main() {
	todo()
}
