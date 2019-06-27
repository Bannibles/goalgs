package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const maxItems = 100000

func test_sorting() {
	// allocate array to sort
	a := make([]int, maxItems)
	// randomly initialize
	for i, max := 0, len(a); i < max; i++ {
		a[i] = rand.Intn(maxItems)
	}

	// print input to be sorted
	fmt.Println(a[:16], "...", a[maxItems-16:])

	sorted0, perf0 := StdGoSort(a)
	fmt.Printf("Go Sort of %d took %q:\n", len(a), perf0)
	fmt.Println(sorted0[:16], "...", sorted0[maxItems-16:])

	sorted1, perf1 := BubbleSort(a)
	fmt.Printf("Bubled sort of %d took %q:\n", len(a), perf1)
	fmt.Println(sorted1[:16], "...", sorted1[maxItems-16:])

	sorted2, perf2 := CocktailShaker(a)
	fmt.Printf("Cocktail Shaker of %d took %q:\n", len(a), perf2)
	fmt.Println(sorted2[:16], "...", sorted2[maxItems-16:])
}

func copyArray(in []int) []int {
	out := make([]int, len(in))
	for i, max := 0, len(in); i < max; i++ {
		out[i] = in[i]
	}
	return out
}

func BubbleSort(input []int) ([]int, time.Duration) {
	// make a copy of original array to sort
	input = copyArray(input)
	max := len(input)

	// start time for perf
	start := time.Now()

	// bubble sort
	for x := 0; x < max; x++ {
		for i := 0; i < max-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	}
	return input, time.Now().Sub(start)
}

func CocktailShaker(input []int) ([]int, time.Duration) {
	input = copyArray(input)
	max := len(input)
	start := time.Now()

	//cocktail shaker
	for x := 0; x < max/2; x++ {
		for i := 0; i < max-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
		for i := max - 1; i > 0; i-- {
			if input[i] < input[i-1] {
				input[i], input[i-1] = input[i-1], input[i]
			}
		}
	}
	return input, time.Now().Sub(start)
}

func StdGoSort(input []int) ([]int, time.Duration) {
	input = copyArray(input)
	start := time.Now()

	sort.Ints(input)

	return input, time.Now().Sub(start)
}

/*
func InsertionSort(input []int) ([]int, time.Duration) {
	input = copyArray(input)
	max := len(input)
	start := time.Now()

	//insertion sort
	for x:= 0; x< max; x++ {

	}
	return input, time.Now().Sub(start)
}
*/
