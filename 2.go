package main

import (
	"fmt"
	"strconv"
)

func t(a int) {
	if a%2 == 0 {
		fmt.Println("четное")
	} else {
		fmt.Println("нечетное")
	}
}

func main() {
	var xstr string
	var xint int
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
	t(xint)
}
