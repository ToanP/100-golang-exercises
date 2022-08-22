package main

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included).  The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 001")
	res := Ex001(2000, 3200)
	fmt.Println("--Slice Declaration--")
	fmt.Println(res)
	fmt.Println("--Slice Declaration & Initialize--")
	res1 := Ex001A(2000, 3200)
	fmt.Println(res1)
	fmt.Println("--Slice created from new function as Pointer--")
	res2 := Ex001New(2000, 3200)
	fmt.Println(res2)
	fmt.Println("--Slice created with make function--")
	res3 := Ex001New(2000, 3200)
	fmt.Println(res3)
}

// Declare an empty Slide
func Ex001(low, high int) string {
	var numbers []string
	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	return strings.Join(numbers, ",")
}

// Declare and Initial empty Slide
func Ex001A(low, high int) (result string) {
	numbers := []string{}
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	result = strings.Join(numbers, ",")
	return

}

// Using Make function to create Slide with (high-low) members
func Ex001Make(low, high int) (result string) {
	numbers := make([]string, high-low)
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	result = strings.Join(numbers, ",")
	return

}

// Using new function to create Slide
func Ex001New(low, high int) (result string) {
	var numbers = new([]string)
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			*numbers = append(*numbers, strconv.Itoa(i))
		}
	}
	result = strings.Join(*numbers, ",")
	return

}
