package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// исходная разница весов Васи и Маши
func getDiff(weights []int) int {
	difference := 0
	for i := range weights {
		if i%2 == 0 {
			difference += weights[i]
		} else {
			difference -= weights[i]
		}
	}
	return difference
}

// вычисление максимальной разницы соседних весов Васи и Маши
func getMaxDiff(weights []int) int {
	minVasya := weights[0]
	maxMasha := weights[1]

	for i, v := range weights {
		if i%2 == 0 {
			if minVasya > v {
				minVasya = v
			}
		}
		if i%2 == 1 {
			if maxMasha < v {
				maxMasha = v
			}
		}
	}
	return (maxMasha - minVasya) * 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	nData := scanner.Text()
	n, _ := strconv.Atoi(strings.TrimSpace(nData))

	scanner.Scan()
	a := strings.Fields(scanner.Text())
	weights := make([]int, n)
	for i := 0; i < n; i++ {
		weights[i], _ = strconv.Atoi(a[i])
	}
	difference := getDiff(weights)
	maxDiff := getMaxDiff(weights)
	if difference < maxDiff {
		fmt.Println(difference + maxDiff)
	} else {
		fmt.Println(difference)
	}

}
