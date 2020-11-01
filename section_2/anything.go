package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim.go")
	Anything(2.44)
	Anything("Angad")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "saiudygapd"
	mymap["age"] = 10
	fmt.Println(mymap)
}
