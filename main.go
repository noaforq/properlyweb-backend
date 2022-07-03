package main

import (
	"strconv"

	"properly.jp/properlyweb-backend/infrastructures"
)

func FizzBuzz(num int) string {
	if num < 1 || 100 < num {
		panic("argument exception")
	}
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%5 == 0:
		return "Buzz"
	case num%3 == 0:
		return "Fizz"
	}
	return strconv.Itoa(num)
}

func main() {
	infrastructures.Run()
}
