package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment.

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var err error
	var i2 uint64
	var d2 float64
	var s2 string

	// Read and save an integer, double, and String to your variables.
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		_, err = strconv.ParseUint(scanner.Text(), 10, 64)
		if err == nil {
			i2, _ = strconv.ParseUint(scanner.Text(), 10, 64)
			continue
		}

		_, err = strconv.ParseFloat(scanner.Text(), 64)
		if err == nil {
			d2, _ = strconv.ParseFloat(scanner.Text(), 64)
			continue
		}

		s2 += scanner.Text() + " "
		continue
	}

	// Print the sum of both integer variables on a new line.
	fmt.Printf("%d\n", i+i2)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+d2)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("%s%s\n", s, s2)
}
