package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ScanNumber(prompt string) int {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(input))
	return num
}

func ScanFloat(prompt string) float64 {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	num, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	return num
}

func ScanString(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}