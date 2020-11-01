package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// if else, for, switch case, break continue

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i")
	}

	arr := []string{"Tak", "Sebec", "Kawamura"}
	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Tak"
	mymap["age"] = 33

	for k, v := range mymap {
		fmt.Printf("Key: %s, Value: %v, ", k, v)
	}
}
