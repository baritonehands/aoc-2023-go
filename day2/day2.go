package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"maps"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseRounds(rounds []string) []map[string]int {
	ret := make([]map[string]int, len(rounds))
	for i, round := range rounds {
		ret[i] = make(map[string]int)
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			parts := strings.Split(cube, " ")
			ret[i][parts[1]], _ = strconv.Atoi(parts[0])
		}
	}
	return ret
}

func part1(limit map[string]int) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		id, _ := strconv.Atoi(parts[0][5:])
		roundsStr := strings.Split(parts[1], "; ")

		valid := true
		for _, round := range parseRounds(roundsStr) {
			for cube, cnt := range round {
				if cnt > limit[cube] {
					valid = false
				}
			}

		}
		if valid {
			sum += id
		}
	}
	return sum
}

func part2() int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		//id, _ := strconv.Atoi(parts[0][5:])
		roundsStr := strings.Split(parts[1], "; ")
		mins := map[string]int{}

		for _, round := range parseRounds(roundsStr) {
			for cube, cnt := range round {
				if cnt > mins[cube] {
					mins[cube] = cnt
				}
			}
		}
		fmt.Printf("%v\n", mins)
		power := it.Fold(maps.Values(mins), func(p int, cnt int) int {
			return p * cnt
		}, 1)
		println(power)
		sum += power
	}
	return sum
}

func main() {
	println(part1(map[string]int{"red": 12, "green": 13, "blue": 14}))
	println(part2())
}
