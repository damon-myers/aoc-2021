package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(infile string) []int {
	filehandle, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(filehandle)

	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text()) // assume we're always given valid numbers
		nums = append(nums, num)
	}

	return nums
}

func calcNumIncreases(depths []int) int {
	counter := 0

	previousDepth := depths[0]
	for i, depth := range depths {
		if i == 0 {
			continue
		}

		if depth > previousDepth {
			counter++
		}

		previousDepth = depth
	}

	return counter
}

func sumArray(array []int) int {
	result := 0
	for _, n := range array {
		result += n
	}

	return result
}

func calcNumIncreasesSliding(depths []int, windowSize int) int {
	counter := 0

	previousSum := sumArray(depths[0:windowSize])
	for i := 1; i <= len(depths)-windowSize; i++ {
		currentSum := sumArray(depths[i : i+windowSize])

		if currentSum > previousSum {
			counter++
		}

		previousSum = currentSum
	}

	return counter
}

func main() {
	infilePtr := flag.String("infile", "inputs/day1.txt", "path to the input file")
	flag.Parse()

	depths := readInput(*infilePtr)

	numIncreases := calcNumIncreases(depths)
	fmt.Println("Number of increases:", numIncreases)

	numIncreasesSliding := calcNumIncreasesSliding(depths, 3)
	fmt.Println("Number of sliding window increases:", numIncreasesSliding)
}
