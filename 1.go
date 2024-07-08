package main

import (
	"fmt"
	"strconv"
)

func calc(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	var x, y int
	var xstr, ystr string

	fmt.Println("введите число")
	_, err := fmt.Scanln(&xstr)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}
	x, err = strconv.Atoi(xstr)
	if err != nil {
		fmt.Println("Ошибка конвертации числа:", err)
		return
	}

	fmt.Println("введите число")
	_, err1 := fmt.Scanln(&ystr)
	if err1 != nil {
		fmt.Println("Ошибка ввода числа:", err1)
		return
	}
	y, err1 = strconv.Atoi(ystr)
	if err1 != nil {
		fmt.Println("Ошибка конвертации числа:", err1)
		return
	}

	calc(x, y)
}
