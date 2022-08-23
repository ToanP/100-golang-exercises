package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Excercise 006")
	fmt.Printf("Input list of numbers: ")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	numbers := Exe006(input)

	fmt.Printf("Output: %T, %v\n", numbers, strings.Join(numbers, ","))

}

func Exe006(input string) []string {
	const C, H = 50, 30
	numStrings := strings.Split(input, ",")

	length := len(numStrings)
	var numbers = make([]string, length)

	for index, v := range numStrings {
		D, _ := strconv.Atoi(strings.Trim(v, " "))
		calculatedVal := math.Sqrt(float64((2 * C * D) / H))
		numbers[index] = strconv.Itoa(int(math.Round(calculatedVal)))
	}

	return numbers
}
