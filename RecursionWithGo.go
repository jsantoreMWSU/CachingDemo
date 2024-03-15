package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lookupTable [256]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("calculate which fibonacci number:")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading - giving up ", err)
	}
	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("you didn't put in an int - giving up ", err)
	}
	fmt.Println("Fib of ", num, " is ", fib(num))
	//fmt.Println("Factorial of ", num, " is ", fact(num))
}

func fib(num int) int {
	if num <= 2 {
		return 1
	}
	if lookupTable[num] != 0 {
		return lookupTable[num]
	}
	result := fib(num-1) + fib(num-2)
	lookupTable[num] = result
	return result
}

func fact(num int) int64 {
	if num <= 1 {
		return 1
	}
	return int64(num) * fact(num-1)
}
