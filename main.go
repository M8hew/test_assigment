package main

import (
	"fmt"
	"os"

	"test_assigment/internal/solver"
)

func main() {
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	result, err := solver.Solve(n, os.Stdin)
	if err != nil {
		fmt.Println("Error solving task:", err)
		os.Exit(1)
	}

	if result {
		fmt.Println("yes")
		return
	}
	fmt.Println("no")
}
