package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 007")

	var s1, s2 string

	fmt.Print("Please enter 2 numbers seperated by space-separated token: ")
	_, err := fmt.Scanln(&s1, &s2)

	if err != nil {
		fmt.Println("Error occurs:", err)
	}

	X, _ := strconv.Atoi(strings.Trim(s1, ","))
	Y, _ := strconv.Atoi(strings.Trim(s2, ","))

	array := Ext007(X, Y)

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Printf("%v\n", array)
}

func Ext007(R, C int) [][]int {
	var array = make([][]int, R)
	for index := range array {
		array[index] = make([]int, C)
	}

	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			array[row][col] = row * col
		}
	}
	return array
}
