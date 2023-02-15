package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("test01")
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
