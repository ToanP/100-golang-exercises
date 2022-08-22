package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var strInput string
	fmt.Println("Enter a string ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() { // wait for user's input from stdin ending up new line/ ENTER
		strInput = scanner.Text()
	}

	// fmt.Println(strInput)

	result := Ex004(strInput)
	fmt.Printf("Map is %T, %v\n", result, result)

}

// Ex004 takes a string of comma-seperated numbers and returns a slice of int
func Ex004(input string) []int {
	// create a map with the size of n
	numbers := strings.Split(input, ",")

	// indentify length of Slide
	length := len(numbers)
	var num = make([]int, length)

	for index, v := range numbers {
		s := strings.Trim(v, " ")
		num[index], _ = strconv.Atoi(s)
	}
	return num
}
