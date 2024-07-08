package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i > 1; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
