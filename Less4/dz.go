package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputNums := []int64{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			panic(err)
		}
		inputNums = append(inputNums, value)

	}

	inputNumsValues := make([]int, len(inputNums))
	for i, val := range inputNums {
		inputNumsValues[i] = int(val)
	}
	sort.Ints(inputNumsValues)
	fmt.Println(inputNumsValues)
}
