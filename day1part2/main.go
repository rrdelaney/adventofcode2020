package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find(nums []int, check1, check2 int, c chan int) {
	base1 := nums[check1]
	base2 := nums[check2]
	for index, num := range nums {
		if index == check1 || index == check2 {
			continue
		}

		if num+base1+base2 == 2020 {
			c <- (num * base1 * base2)
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
	for check1 := range numbers {
		for check2 := range numbers {
			go find(numbers, check1, check2, c)
		}
	}

	result := <-c
	fmt.Println(result)
}
