package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	nValue, inputPath := getInputFlags()

	file, err := os.Open(*inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	array := buildTopNArray(*nValue, scanner)

	printTopN(array)
}

func getInputFlags() (*int, *string) {
	nValue := flag.Int("topn", 50, "the n-value in topn")
	inputPath := flag.String("filePath", "'input file'", "path to input file")
	flag.Parse()

	return nValue, inputPath
}

func buildTopNArray(nValue int, scanner *bufio.Scanner) []int {
	if nValue <= 0 {
		return []int{}
	}

	var array = make([]int, nValue)
	for i, _ := range array {
		array[i] = math.MinInt64
	}

	validRows := 0
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// not an int
			continue
		}
		array = sortedInsert(array, current)
		validRows++
	}

	// Sort the array one last time outside the loop
	// Catches cases when arraySize < nValue and is unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(array)))

	if validRows < len(array) {
		array = array[:validRows]
	}

	return array
}

func sortedInsert(array []int, current int) []int {
	len := len(array)

	i := sort.Search(len, func(i int) bool { return array[i] > current })
	if current >= array[0] {
		copy(array[:i-1], array[1:i])
		array[i-1] = current
	}
	return array
}

func printTopN(array []int) {
	for _, i := range array {
		fmt.Printf("%d\n", i)
	}
}
