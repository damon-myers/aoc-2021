package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// We will count the number of 1s that appear in each column
// Together with the total number of lines, we can determine
// whether 1 or 0 was more common in that column
type bitCounter struct {
	oneCounter []int
	totalLines int
}

// Assumes we're always given valid input
func readInput(infile string) bitCounter {
	filehandle, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(filehandle)

	var counter bitCounter

	for scanner.Scan() {
		line := scanner.Text()

		counter.totalLines++

		// assume all rows in input have the same width
		if len(counter.oneCounter) == 0 {
			counter.oneCounter = make([]int, len(line))
		}

		for columnIdx, digitRune := range line {
			// assume it's a number that's either one or zero
			digit := int(digitRune - '0')

			if digit == 1 {
				counter.oneCounter[columnIdx]++
			}
		}
	}

	return counter
}

func calculatePowerConsumption(counter bitCounter) {
	mostCommonBits := make([]rune, len(counter.oneCounter))
	leastCommonBits := make([]rune, len(counter.oneCounter))

	for idx, oneCount := range counter.oneCounter {
		// in the event of an equal count, the "most common" value will be 0
		if oneCount > counter.totalLines/2 {
			mostCommonBits[idx] = '1'
			leastCommonBits[idx] = '0'
		} else {
			mostCommonBits[idx] = '0'
			leastCommonBits[idx] = '1'
		}
	}

	gammaRate, _ := strconv.ParseInt(string(mostCommonBits), 2, 0)
	epsilonRate, _ := strconv.ParseInt(string(leastCommonBits), 2, 0)

	fmt.Println("Gamma rate:", gammaRate)
	fmt.Println("Epsilon rate:", epsilonRate)
	fmt.Println("Power consumption:", gammaRate*epsilonRate)
}

func Run(filename string) {
	counter := readInput(filename)

	calculatePowerConsumption(counter)
}
