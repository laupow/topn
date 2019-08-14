package main

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestStandardCase(t *testing.T) {
	nValue := 6
	testInput := "68\n1000\n34\n1\n5\n8\n9\n10\n23\n2\n3\n49"
	arrayShouldBe := []int{1000, 68, 49, 34, 23, 10}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestNvalueLargerThanInput(t *testing.T) {
	nValue := 6
	testInput := "101\n5\n394"
	arrayShouldBe := []int{394, 101, 5}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestLongInputFromFile(t *testing.T) {
	nValue := 10
	arrayShouldBe := []int{10000, 9999, 9998, 9997, 9996, 9995, 9994, 9993, 9992, 9991}

	filename := "test10kNumbers.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outputArray := buildTopNArray(nValue, bufio.NewScanner(file))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestEmptyFile(t *testing.T) {
	nValue := 10
	arrayShouldBe := []int{}

	filename := "testEmpty.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outputArray := buildTopNArray(nValue, bufio.NewScanner(file))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestNoNValue(t *testing.T) {
	nValue := 0
	testInput := "68\n1000\n34\n1\n5\n8\n9\n10\n23\n2\n3\n49"
	arrayShouldBe := []int{}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestNegativeNValue(t *testing.T) {
	nValue := -5
	testInput := "68\n1000\n34\n1\n5\n8\n9\n10\n23\n2\n3\n49"
	arrayShouldBe := []int{}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestGarbageLine(t *testing.T) {
	nValue := 50
	testInput := "\n\n\n   \n #\none\n913sad\n\n5.3\n "
	arrayShouldBe := []int{}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestZeroAtEnd(t *testing.T) {
	nValue := 4
	testInput := "\n\n\n885\n913\n\n5\n0"
	arrayShouldBe := []int{913, 885, 5, 0}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestEmptyLine(t *testing.T) {
	nValue := 50
	testInput := "\n\n\n885\n913\n\n5"
	arrayShouldBe := []int{913, 885, 5}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestNegativeInputValue(t *testing.T) {
	nValue := 50
	testInput := "885\n-860\n380\n333\n913\n-91\n172\n-321\n-500\n"
	arrayShouldBe := []int{913, 885, 380, 333, 172, -91, -321, -500, -860}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func TestNegativeInputValueWithZero(t *testing.T) {
	nValue := 50
	testInput := "885\n-860\n380\n333\n0\n913\n-91\n172\n-321\n-500\n"
	arrayShouldBe := []int{913, 885, 380, 333, 172, 0, -91, -321, -500, -860}

	outputArray := buildTopNArray(nValue, getScanner(testInput))

	if !reflect.DeepEqual(arrayShouldBe, outputArray) {
		t.Errorf("Expected array of %v, but it was %v instead.", arrayShouldBe, outputArray)
	}
}

func Benchmark10NvalueSort(b *testing.B) {
	nValue := 10

	input := rangeInt(1, 1000000000000000, 10000)

	var array = make([]int, nValue)

	for i := 0; i < b.N; i++ {
		for _, j := range input {
			array = sortedInsert(array, j)
		}
	}
}

func Benchmark100NvalueSort(b *testing.B) {
	nValue := 100

	input := rangeInt(1, 1000000000000000, 10000)

	var array = make([]int, nValue)

	for i := 0; i < b.N; i++ {
		for _, j := range input {
			array = sortedInsert(array, j)
		}
	}
}

func Benchmark1000NvalueSort(b *testing.B) {
	nValue := 1000

	input := rangeInt(1, 1000000000000000, 10000)

	var array = make([]int, nValue)

	for i := 0; i < b.N; i++ {
		for _, j := range input {
			array = sortedInsert(array, j)
		}
	}
}

func Benchmark10000NvalueSort(b *testing.B) {
	nValue := 10000

	input := rangeInt(1, 1000000000000000, 10000)

	var array = make([]int, nValue)

	for i := 0; i < b.N; i++ {
		for _, j := range input {
			array = sortedInsert(array, j)
		}
	}
}

func Benchmark100000NvalueSort(b *testing.B) {
	nValue := 100000

	input := rangeInt(1, 1000000000000000, 1000000)

	var array = make([]int, nValue)

	for i := 0; i < b.N; i++ {
		for _, j := range input {
			array = sortedInsert(array, j)
		}
	}
}

func BenchmarkTwoMillionInputs(b *testing.B) {
	nValue := 500

	input := rangeInt(1, 1000000000000000, 2000000)

	var array = make([]int, nValue)

	for i := 0; i < b.N; i++ {
		for _, j := range input {
			array = sortedInsert(array, j)
		}
	}
}

func getScanner(testInput string) *bufio.Scanner {
	var r io.Reader
	r = strings.NewReader(testInput)
	scanner := bufio.NewScanner(r)
	return scanner
}

func rangeInt(min int, max int, n int) []int {
	arr := make([]int, n)
	var r int
	for r = 0; r <= n-1; r++ {
		arr[r] = rand.Intn(max) + min
	}
	return arr
}
