package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(mealCost float64, tipPercent int32, taxPercent int32) {
	// Write your code here
	tip := (mealCost / 100) * float64(tipPercent)
	tax := (mealCost / 100) * float64(taxPercent)
	totalCost := mealCost + tax + tip
	fmt.Println(math.Round(totalCost))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	mealCost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
	checkError(err)

	tipPercenttemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tipPercent := int32(tipPercenttemp)

	taxPercenttemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	taxPercent := int32(taxPercenttemp)

	solve(mealCost, tipPercent, taxPercent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
