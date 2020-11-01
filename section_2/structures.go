package main

import (
	"fmt"
)

// structure encapsulation
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{"honda", 2, 1001}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}
