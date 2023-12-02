package main

import (
	"fmt"
	"github.com/grambbledook/adventofcode2023/day1"
	"github.com/grambbledook/adventofcode2023/day2"
)

func main() {
	fmt.Println("Day 1", "Task 1")
	fmt.Println("Result", day1.Compute(day1.Task1))

	fmt.Println("Day 1", "Task 2")
	fmt.Println("Result", day1.Compute(day1.Task2))

	fmt.Println("Day 2", "Task 1")
	fmt.Println("Result", day2.Compute(day2.Task1))

	fmt.Println("Day 2", "Task 2")
	fmt.Println("Result", day2.Compute(day2.Task2))
}
