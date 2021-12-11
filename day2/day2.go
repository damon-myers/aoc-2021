package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SubmarineCommand struct {
	direction string
	amount    int
}

type SubmarinePosition struct {
	horizontalPosition, depth int
}

func MaxInts(x, y int) int {
	if x >= y {
		return x
	}

	return y
}

// Assumes we're always given valid input
func readInput(infile string) []SubmarineCommand {
	filehandle, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(filehandle)

	var commands []SubmarineCommand
	for scanner.Scan() {
		line := scanner.Text()

		splitResult := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splitResult[1])

		commands = append(commands, SubmarineCommand{direction: splitResult[0], amount: amount})
	}

	return commands
}

func calculatePosition(commands []SubmarineCommand) SubmarinePosition {
	var subPos SubmarinePosition

	for _, command := range commands {
		switch command.direction {
		case "forward":
			subPos.horizontalPosition += command.amount
		case "up":
			subPos.depth -= command.amount

			subPos.depth = MaxInts(0, subPos.depth) // don't go above 0
		case "down":
			subPos.depth += command.amount
		}
	}

	return subPos
}

func calculatePositionAndAim(commands []SubmarineCommand) (SubmarinePosition, int) {
	var subPos SubmarinePosition
	var aim int

	for _, command := range commands {
		switch command.direction {
		case "forward":
			subPos.horizontalPosition += command.amount

			subPos.depth += aim * command.amount
		case "up":
			aim -= command.amount
		case "down":
			aim += command.amount
		}
	}

	return subPos, aim
}

func Run(filename string) {
	commands := readInput(filename)

	result1 := calculatePosition(commands)

	fmt.Println("Resulting location", result1.horizontalPosition, result1.depth)
	fmt.Println("Multiplied:", result1.horizontalPosition*result1.depth)

	result2, aim := calculatePositionAndAim(commands)

	fmt.Println("Resulting location (part 2)", result2.horizontalPosition, result2.depth)
	fmt.Println("Resulting aim (part 2)", aim)
	fmt.Println("Multiplied:", result2.horizontalPosition*result2.depth)
}
