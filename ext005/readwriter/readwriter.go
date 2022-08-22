package readwriter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckModule() {
	fmt.Println("Exercise 005")
}

func ReadString() string {
	fmt.Printf("Please input: ")
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	return str
}

func PrintString(str string) {
	fmt.Println(strings.ToUpper(str))
}
