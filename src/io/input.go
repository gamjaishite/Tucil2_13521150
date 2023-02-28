package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Requesting input from user through stdin and do validation
func GetInput(text string, minRange uint32, maxRange uint32) (uint32, string) {
	fmt.Print(text, ": ")

	reader := bufio.NewReader(os.Stdin)

	rawInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Sprintf("Failed to Receive %s. Please Try Again!\n", text)
	}
	trimInput := strings.TrimSpace(rawInput)

	N, err := strconv.ParseUint(trimInput, 10, 32)
	if err != nil || N < uint64(minRange) || N > uint64(maxRange) {
		return 0, fmt.Sprintf("%s Must Be a Number Between %d - %d. Please Try Again!\n", text, minRange, maxRange)
	}
	return uint32(N), ""
}

// Requesting input from user using InputPointsAmount function
func Input() (uint32, uint32) {
	for {
		N, errN := GetInput("Points Amount", 2, 100000)
		dimension, errD := GetInput("Dimension", 1, 10)
		if errN != "" || errD != "" {
			fmt.Println()
			fmt.Print(errN)
			fmt.Print(errD)
			fmt.Println()
		} else {
			return N, dimension
		}
	}
}
