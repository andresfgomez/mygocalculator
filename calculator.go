package mygocalculator

import (
	"fmt"
)

type Calculator struct {
}

func (oper *Calculator) Add(a int, b int) (c int) {
	return a + b
}

func Start() {
	fmt.Println("Starting calculator")
}
