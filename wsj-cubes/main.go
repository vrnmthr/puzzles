package main

import (
	"math"
	"sort"
	"github.com/kavehmz/prime"
	"fmt"
)

const (
	MinAge = 2
	MaxAge = 19
	NumChildren = 3
	YearDiff = 2
)

var primes = prime.Primes(10000)

func cube(x int) int {
	return int(math.Pow(float64(x), 3))
}

func intMin(ints []int) int {
	min := ints[0]
	for _, n := range ints {
		if n < min {
			min = n
		}
	}
	return min
}

func intMax(ints []int) int {
	max := ints[0]
	for _, n := range ints {
		if n > max {
			max = n
		}
	}
	return max
}

func sum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func cubedSum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += cube(i)
	}
	return sum
}

// combine creates all unique combinations of len non-unique integers in (min, max)
func combine(min, max, len int) [][]int {
	results := make([][]int, 0)
	if len == 0 {
		return results
	}
	// aggregate map starts with single numbers
	agg := make(map[int][][]int)
	for i := min; i <= max; i++ {
		agg[i] = [][]int{[]int{i}}
	}
	for l := 0; l < len - 1; l ++ {
		agg = combineAggregate(min, max, agg)
	}

	// flatten this map into results
	for _, slices := range agg {
		results = append(results, slices...)
	}
	return results
}

// combineAggregate appends every number (min, max) to each slice that does not contain any lower numbers.
// The agg map is a map of int:slices that only have numbers >= int
func combineAggregate(min, max int, agg map[int][][]int) map[int][][]int {
	res := make(map[int][][]int)
	// loops over all numbers in range
	for i := min; i <= max; i ++ {
		// only append this number to slices that do not contain any lower numbers
		for j := i; j <= max; j++ {
			for _, cmb := range agg[j] {
				// min of the new slice is the min of the key and the number added
				min := intMin([]int{i,j})
				res[min] = append(res[min], append(cmb, i))
			}
		}
	}
	return res
}

// id returns a unique ID identifying any slice of integers
func id(ints []int) uint64 {
	if intMax(ints) > len(primes) {
		panic("int slice is out of range for ID")
	}
	var id uint64
	id = 1
	for i := range ints {
		id *= primes[ints[i]]
	}
	return id
}

// solve solves the puzzle and outputs a slice of all possible solutions
func solve() [][]int {

	// get all possible solutions
	guesses := combine(MinAge, MaxAge, NumChildren)
	for _, guess := range guesses {
		sort.Ints(guess)
	}

	// find all solutions
	solutions := make(map[uint64][]int)
	for _, firstGuess := range guesses {
		for _, secondGuess := range guesses {
			if id(firstGuess) != id(secondGuess) && sum(firstGuess) == sum(secondGuess) && cubedSum(firstGuess) == cubedSum(secondGuess){
				solutions[id(firstGuess)] = firstGuess
				solutions[id(secondGuess)] = secondGuess
			}
		}
	}

	// find a set of solutions that are YearDiff apart
	answers := make([][]int, 0)
	for _, solution := range solutions {
		translated := make([]int, len(solution))
		for i := range solution {
			translated[i] = solution[i] - YearDiff
		}
		if _, ok := solutions[id(translated)]; ok {
			answers = append(answers, solution)
		}
	}

	return answers
}

func main() {
	fmt.Println(solve())
}
