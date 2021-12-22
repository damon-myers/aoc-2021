package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type selectedNumber struct {
	row int
	col int
}

type bingoBoard struct {
	numbers         [][]int
	selectedNumbers []selectedNumber
	hasWon          bool
}

type bingo struct {
	boards        []bingoBoard
	randomNumbers []int
}

func convertStringsToInts(strings []string) []int {
	numbers := make([]int, len(strings))

	for index, str := range strings {
		numInt, _ := strconv.Atoi(str)
		numbers[index] = numInt
	}

	return numbers
}

// scanner passed in should have the first line of the board in memory
func parseBoard(scanner *bufio.Scanner, boardSize int) bingoBoard {
	var bingoBoard bingoBoard
	bingoBoard.numbers = make([][]int, boardSize)

	for i := 0; i < boardSize; i++ {
		line := scanner.Text()
		bingoBoard.numbers[i] = convertStringsToInts(strings.Fields(line))

		if i != boardSize-1 {
			// don't advance past the last line of the board
			scanner.Scan()
		}
	}

	return bingoBoard
}

// Assumes we're always given valid input
func readInput(infile string) bingo {
	filehandle, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(filehandle)

	var bingo bingo

	// first line is the list of numbers that will be picked
	scanner.Scan()
	selectedNumbersStr := scanner.Text()

	bingo.randomNumbers = convertStringsToInts(strings.Split(selectedNumbersStr, ","))

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else {
			bingo.boards = append(bingo.boards, parseBoard(scanner, 5))
		}
	}

	return bingo
}

func markCurrentNumber(board *bingoBoard, currentNumber int) {
	for rowIdx, row := range board.numbers {
		for colIdx, number := range row {
			if number == currentNumber {
				board.selectedNumbers = append(board.selectedNumbers, selectedNumber{rowIdx, colIdx})
			}
		}
	}
}

func isWinByRow(boardSize int, selectedNumber selectedNumber, selectedNumbers []selectedNumber) bool {
	numInRow := 0

	for _, boardIndex := range selectedNumbers {
		if boardIndex.row == selectedNumber.row {
			numInRow++
		}

		if numInRow == boardSize {
			return true
		}
	}

	return false
}

func isWinByCol(boardSize int, selectedNumber selectedNumber, selectedNumbers []selectedNumber) bool {
	numInRow := 0

	for _, boardIndex := range selectedNumbers {
		if boardIndex.col == selectedNumber.col {
			numInRow++
		}

		if numInRow == boardSize {
			return true
		}
	}

	return false
}

func isWinningBoard(board *bingoBoard) bool {
	boardSize := len(board.numbers)
	if len(board.selectedNumbers) < boardSize {
		return false
	}

	for _, selectedNumber := range board.selectedNumbers {
		if isWinByRow(boardSize, selectedNumber, board.selectedNumbers) || isWinByCol(boardSize, selectedNumber, board.selectedNumbers) {
			return true
		}
	}

	return false
}

func playBingo(bingo bingo) (bingoBoard, int) {
	for _, currentNumber := range bingo.randomNumbers {
		for boardIdx := range bingo.boards {
			boardPtr := &bingo.boards[boardIdx]
			markCurrentNumber(boardPtr, currentNumber)

			if isWinningBoard(boardPtr) {
				return *boardPtr, currentNumber
			}
		}
	}

	return bingoBoard{}, -1
}

func isLastWinningBoard(boards []bingoBoard) bool {
	remainingBoardCount := len(boards)

	for _, board := range boards {
		if board.hasWon {
			remainingBoardCount--
		}
	}

	return remainingBoardCount == 1
}

func playBingoToLose(bingo bingo) (bingoBoard, int) {
	for _, currentNumber := range bingo.randomNumbers {
		for boardIdx, currentBoard := range bingo.boards {
			if currentBoard.hasWon {
				continue
			}

			boardPtr := &bingo.boards[boardIdx]
			markCurrentNumber(boardPtr, currentNumber)

			if isWinningBoard(boardPtr) {
				if isLastWinningBoard(bingo.boards) {
					return *boardPtr, currentNumber
				} else {
					boardPtr.hasWon = true
				}
			}
		}
	}

	return bingoBoard{}, -1
}

func isSelected(selectedNumbers []selectedNumber, rowIdx int, colIdx int) bool {
	for _, selectedNumber := range selectedNumbers {
		if selectedNumber.col == colIdx && selectedNumber.row == rowIdx {
			return true
		}
	}

	return false
}

func scoreBoard(board bingoBoard, finalNumber int) int {
	sumUnmarked := 0

	for rowIdx, row := range board.numbers {
		for colIdx, num := range row {
			if isSelected(board.selectedNumbers, rowIdx, colIdx) {
				continue
			}

			sumUnmarked += num
		}
	}

	return sumUnmarked * finalNumber
}

func deepCopyBingo(originalBingo *bingo) bingo {
	var bingoCopy bingo

	for _, board := range originalBingo.boards {
		bingoCopy.boards = append(bingoCopy.boards, board)
	}

	bingoCopy.randomNumbers = originalBingo.randomNumbers

	return bingoCopy
}

func printBoard(board bingoBoard) {
	colorRed := "\033[31m"
	colorWhite := "\033[37m"

	for rowIdx, row := range board.numbers {
		for colIdx, num := range row {
			if isSelected(board.selectedNumbers, rowIdx, colIdx) {
				fmt.Print(string(colorRed))
				fmt.Printf("%2d", num)
			} else {
				fmt.Print(string(colorWhite))
				fmt.Printf("%2d", num)
			}
			fmt.Print(" ")
		}

		fmt.Println()
	}
}

func Run(filename string) {
	bingo := readInput(filename)
	bingoCopy := deepCopyBingo(&bingo)

	winningBoard, finalNumber := playBingo(bingo)

	fmt.Println("Winning board")
	printBoard(winningBoard)
	fmt.Println("Final number selected", finalNumber)

	scorePart1 := scoreBoard(winningBoard, finalNumber)

	fmt.Println("Part 1 winning board score:", scorePart1)

	losingBoard, finalNumber := playBingoToLose(bingoCopy)

	fmt.Println("Losing board")
	printBoard(losingBoard)
	fmt.Println("Final number selected", finalNumber)

	scorePart2 := scoreBoard(losingBoard, finalNumber)

	fmt.Println("Part 2 losing board score:", scorePart2)
}
