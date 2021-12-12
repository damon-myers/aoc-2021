package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// We will count the number of 1s that appear in each column
// Together with the total number of lines, we can determine
// whether 1 or 0 was more common in that column
type bitCounter struct {
	oneCounter []int
	totalLines int
}

// Assumes we're always given valid input
func readInput(infile string) ([]string, bitCounter) {
	filehandle, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(filehandle)

	var counter bitCounter
	// the list of rows
	var diagReport []string

	for scanner.Scan() {
		line := scanner.Text()

		diagReport = append(diagReport, line)

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

	return diagReport, counter
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

func calculateDecimalValue(bitArray []int) int {
	strArray := make([]string, len(bitArray))

	for i, digit := range bitArray {
		strArray[i] = strconv.Itoa(digit)
	}

	decimalValue, _ := strconv.ParseInt(strings.Join(strArray, ""), 2, 0)

	return int(decimalValue)
}

func calculateOxygenRating(diagReport [][]int) int {
	// a "set" of numbers that are still being considered
	var numsConsidered = make(map[int][]int)

	// setup map
	for idx, row := range diagReport {
		numsConsidered[idx] = row
	}

	// filter results until one number is left (assume we'll always terminate)
	for bitPos := 0; bitPos < len(diagReport[0]); bitPos++ {
		remainingRows := getRows(numsConsidered)
		mostCommonNum, bothEqual := getMostCommonNumber(bitPos, remainingRows)

		for idx, row := range numsConsidered {
			if bothEqual && row[bitPos] == 1 {
				continue
			} else if row[bitPos] == mostCommonNum {
				continue
			}

			delete(numsConsidered, idx)
		}

		if len(numsConsidered) == 1 {
			break
		}
	}

	// should only be one value remaining
	for _, row := range numsConsidered {
		return calculateDecimalValue(row)
	}

	return -1
}

// First positional return - most common number
// Second position return - whether the frequency of numbers was equal
func getMostCommonNumber(col int, rows [][]int) (int, bool) {
	oneCount := 0

	for _, row := range rows {
		if row[col] == 1 {
			oneCount++
		}
	}

	if len(rows)%2 == 0 && oneCount == len(rows)/2 {
		return -1, true
	}

	if oneCount > len(rows)/2 {
		return 1, false
	}

	return 0, false
}

func getRows(numsConsidered map[int][]int) [][]int {
	var rows [][]int

	for _, row := range numsConsidered {
		rows = append(rows, row)
	}

	return rows
}

// could probably make these two functions share the map setup and iteration :shrug:
func calculateCO2Rating(diagReport [][]int) int {
	// a "set" of numbers that are still being considered
	var numsConsidered = make(map[int][]int)

	// setup map
	for idx, row := range diagReport {
		numsConsidered[idx] = row
	}

	// filter results until one number is left (assume we'll always terminate)
	for bitPos := 0; bitPos < len(diagReport[0]); bitPos++ {
		remainingRows := getRows(numsConsidered)
		mostCommonNum, bothEqual := getMostCommonNumber(bitPos, remainingRows)

		for idx, row := range numsConsidered {
			if row[bitPos] == mostCommonNum || (bothEqual && row[bitPos] == 1) {
				delete(numsConsidered, idx)
				continue
			}
		}

		if len(numsConsidered) == 1 {
			break
		}
	}

	// should only be one value remaining
	for _, row := range numsConsidered {
		return calculateDecimalValue(row)
	}

	return -1
}

func convertDiagReport(diagReport []string) [][]int {
	converted := make([][]int, len(diagReport))
	rowWidth := len(diagReport[0])
	for i := range converted {
		converted[i] = make([]int, rowWidth)
	}

	for i, row := range diagReport {
		for j, digitRune := range row {
			digit := int(digitRune - '0')

			converted[i][j] = digit
		}
	}

	return converted
}

func calculateLifeSupportRating(diagReport []string, counter bitCounter) {
	diagReportConverted := convertDiagReport(diagReport)
	oxygenRating := calculateOxygenRating(diagReportConverted)
	co2ScrubberRating := calculateCO2Rating(diagReportConverted)

	fmt.Println("Oxygen generator rating:", oxygenRating)
	fmt.Println("CO2 scrubber rating:", co2ScrubberRating)
	fmt.Println("Life support rating:", oxygenRating*co2ScrubberRating)
}

func Run(filename string) {
	diagReport, counter := readInput(filename)

	calculatePowerConsumption(counter)
	calculateLifeSupportRating(diagReport, counter)
}
