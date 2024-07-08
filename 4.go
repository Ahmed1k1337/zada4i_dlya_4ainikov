package main

import (
	"fmt"
	"strconv"
)

func main() {
	var xint int
	var xstr string
	fmt.Println("введите число")
	_, err := fmt.Scanln(&xstr)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}
	xint, err = strconv.Atoi(xstr)
	if err != nil {
		fmt.Println("Ошибка конвертации числа:", err)
		return
	}
	answer := 1
	for i := 1; i <= xint; i++ {
		answer *= i
	}
	fmt.Println(answer)
}
