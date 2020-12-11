package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := []int{0}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		adapters = append(adapters, next)
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	oneJumps := 0
	threeJumps := 0
	for i, jolts := range adapters[1:] {
		prevJolts := adapters[i]
		if jolts-prevJolts == 1 {
			oneJumps++
		}
		if jolts-prevJolts == 3 {
			threeJumps++
		}
	}

	fmt.Println(oneJumps * threeJumps)
}
