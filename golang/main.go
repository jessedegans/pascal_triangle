package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//measureTime()
	ask()
}
func measureTime() {
	now := time.Now()
	defer func() {
		pascalTriangleOptimized(5)
		fmt.Println(time.Since(now))
	}()
}
func ask() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Put in pascal triangle depth number")
	for scanner.Scan() {
		if number, err := strconv.Atoi(scanner.Text()); err == nil {
			fmt.Println(pascalTriangleOptimized(number))
		} else {
			fmt.Println("That's not a number")
		}

	}

}

func pascalTriangle(length int) []int {
	var res []int
	i := 0

	for i < length+1 {
		res = append(res, findVal(i, length))
		i++
	}

	return res
}

func pascalTriangleOptimized(length int) []int {
	var res []int
	i := 0

	for i < length/2+1 {
		res = append(res, findVal(i, length))
		i++
	}

	if length%2 == 0 {
		i = i - 1
	}

	for i > 0 {
		res = append(res, res[i-1])
		i--
	}

	return res
}

func findVal(x int, y int) int {
	if x == 0 {
		return 1
	}
	if y == 0 {
		return 0
	}
	return findVal(x, y-1) + findVal(x-1, y-1)
}

func elapsed() func([]int) {
	start := time.Now()
	return func(res []int) {
		fmt.Printf("took %v\n", time.Since(start))
		fmt.Println(res)
	}
}
