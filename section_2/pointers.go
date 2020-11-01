package main

import "fmt"

// func main() {
// 	var m1 = 2
// 	var ptr = m1
// 	fmt.Println(ptr)
// }

// func main() {
// 	m1 := 2
// 	ptr := &m1
// &はreference?
// fmt.Println(ptr)

// *を入れると整数に戻る
// 	fmt.Println(*ptr)
// }

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func main() {
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
}
