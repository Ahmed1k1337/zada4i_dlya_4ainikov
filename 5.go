package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 0
	b := 1
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
	if xint == 1 {
		fmt.Println(a)
	} else if xint == 2 {
		fmt.Println(b)

	} else if xint < 1 {
		fmt.Println("введите число >=1")
	} else {
		fmt.Print(a)
		fmt.Print("  ", b)
		for i := 0; i < xint-2; i++ {
			res := a + b
			a = b
			b = res
			fmt.Print("  ", res)
		}
	}
}
