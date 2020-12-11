package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find(nums []int, check int, c chan int) {
	base := nums[check]
	for index, num := range nums {
		if index == check {
			continue
		}

		if num+base == 2020 {
			c <- (num * base)
		}
	}
}

func main() {
	numbers := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("could not convert number: %v", err)
			os.Exit(1)
		}

		numbers = append(numbers, num)
	}

	c := make(chan int)
	for index := range numbers {
		go find(numbers, index, c)
	}

	result := <-c
	fmt.Println(result)
}
