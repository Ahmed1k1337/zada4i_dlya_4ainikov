package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func getInput() string { //считывает ввод пользователя и возвращает его в виде строки
	fmt.Println("введите число")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении ввода: %v", err)
	}
	return ""
}

func parseNumber(input string) (int64, error) { //преобразует строку в число и возвращает его
	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Printf("Ошибка преобразования строки в число: %v", err)
		return 0, fmt.Errorf("введено не число или число слишком велико/мало")
	}
	return number, nil
}

func main() {

	num := rand.Intn(100)

	fmt.Println("числа должны быть от 1 до 100")
	for {
		pnum, err := parseNumber(getInput())
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		if pnum > 100 || pnum < 1 {
			fmt.Println("я же просил от 1 до 100")
			break
		}
		if int64(num) < pnum {
			fmt.Print("ваше больше задуманного, ")
		} else if int64(num) > pnum {
			fmt.Print("ваше число меньше задуманного, ")
		} else {
			fmt.Println("угадали, вы че за тигр")
			break
		}
	}

}
