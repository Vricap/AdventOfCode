package main

import (
	"bufio"
	"os"
	"strconv"
)

func Day1()  {
	slice := make([]int, 0)
	file ,_ := os.Open("day1_input.txt")
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		slice  = append(slice, num)
	}

	newSlice := make([]int, 0)
	sum := 0
	for _, val := range slice {
		if val == 0 {
			newSlice = append(newSlice, sum)
			sum = 0
			continue
		}
		sum += val
	}

	findMax(newSlice)
	topThree(newSlice)
}

func findMax(newSlice []int) int {
	max := newSlice[0]
	for _, val := range newSlice {
		if val > max {
			max = val
		}
	}
	return max
}

func topThree(newSlice []int) int {
	one := 0
	for i, val := range newSlice {
		if val > newSlice[one] {
			one = i
		}
	}

	two := 0
	for i, val := range newSlice {
		if i == one {
			continue
		}
		if val > newSlice[two] {
			two = i
		}
	}

	three := 0
	for i, val := range newSlice {
		if i == one || i == two {
			continue
		}

		if val > newSlice[three] {
			three = i
		}
	}

	return newSlice[one] + newSlice[two] + newSlice[three]
}