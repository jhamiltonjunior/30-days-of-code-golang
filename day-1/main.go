package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var i2 uint64
	var d2 float64
	var s2 string
	// Read and save an integer, double, and String to your variables.
	if scanner.Scan() {
		texto, _ := strconv.Atoi(scanner.Text())
		i2 = uint64(texto)
		fmt.Println(i2 + i)
	}
	if scanner.Scan() {
		texto := scanner.Text()
		d2, _ = strconv.ParseFloat(texto, 64)
		fmt.Printf("%0.1f\n", d2+d)
	}

	if scanner.Scan() {
		s2 = scanner.Text()
		fmt.Println(s + s2)
	}

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}
