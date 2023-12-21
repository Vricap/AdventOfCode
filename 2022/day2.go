package main

import (
	"bufio"
	"os"
	"strings"
)

func Day2()  {
	file, _ := os.Open("day2_input.txt")
	scanner := bufio.NewScanner(file)
	partOne(scanner)
	partTwo(scanner)
}

func partOne(scanner *bufio.Scanner) int {
	maps := map[string]int{"AX": 4, "AY": 8, "AZ": 3, "BX": 1, "BY": 5, "BZ": 9, "CX": 7, "CY": 2, "CZ": 6}	
	total := 0
	for scanner.Scan(){
		teks := strings.Replace(scanner.Text(), " ", "", -1)
		total += maps[teks]
	}
	return total
}

func partTwo(scanner *bufio.Scanner) int {
	maps := map[string]int{"AX": 3, "AY": 4, "AZ": 8, "BX": 1, "BY": 5, "BZ": 9, "CX": 2, "CY": 6, "CZ": 7}
	total := 0
	for scanner.Scan(){
		teks := strings.Replace(scanner.Text(), " ", "", -1)
		total += maps[teks]
	}
	return total
}