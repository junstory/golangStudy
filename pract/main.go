package main

import (
	"fmt"

	"github.com/junstory/golangStudy/ex4"
)

func main() {
	var num1 int
	var num2 int
	var op string
	fmt.Println("24 45 47 50 54 76 94")
	fmt.Println("Calculator")
	fmt.Println("첫 번째 숫자 입력: ")
	fmt.Scan(&num1)
	fmt.Println("두 번째 숫자 입력: ")
	fmt.Scan(&num2)
	fmt.Println("연산자 입력(+, -, *, /) : ")
	fmt.Scan(&op)

	ex4.Calc(num1, num2, op)
}
