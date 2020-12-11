package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

type Mem struct {
	mu    sync.Mutex
	cache map[string]int
}

func CanRemove(apts []int, pos int) bool {
	if len(apts) <= 3 {
		return false
	}

	diff := apts[pos+1] - apts[pos-1]
	return diff <= 3
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func Hashcode(s []int) string {
	return fmt.Sprintf("%v", s)
}

func CountVariations(apts []int, m *Mem) int {
	if len(apts) == 3 {
		return 1
	}

	hash := Hashcode(apts)
	m.mu.Lock()
	cached := m.cache[hash]
	m.mu.Unlock()
	if cached != 0 {
		return cached
	}

	c := make(chan int, 2)
	go func() {
		c <- CountVariations(apts[1:], m)
	}()

	if CanRemove(apts, 1) {
		go func() {
			withRemoved := RemoveIndex(apts, 1)
			c <- CountVariations(withRemoved, m)
		}()
	} else {
		c <- 0
	}

	vars := <-c + <-c

	m.mu.Lock()
	if m.cache[hash] == 0 {
		m.cache[hash] = vars
	}
	m.mu.Unlock()

	return vars
}

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

	mem := Mem{cache: make(map[string]int)}
	fmt.Println(CountVariations(adapters, &mem))
}
