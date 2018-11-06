package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := make([]string, 0, 3)
	var T int
	var line string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&line)
		_, err := strconv.Atoi(line)
		if err != nil {
			lines = append(lines, line)
		} else {
			continue
		}
	}

	for i := 0; len(lines)-1 >= i; i++ {
		var even []string
		var odd []string

		for ind, v := range lines[i] {
			if ind%2 == 0 {
				even = append(even, string(v))
			} else {
				odd = append(odd, string(v))
			}
		}

		fmt.Printf("%s %s\n", strings.Join(even, ""), strings.Join(odd, ""))
	}
}

// Split defines the delimeters we will split the user input with
func Split(r rune) bool {
	return r == ' ' || r == ','
}
