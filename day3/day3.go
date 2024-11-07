package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/itx"
	"github.com/baritonehands/aoc-2021-go/utils"
	"iter"
	"maps"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func neighbors(p Pair, width int, xMax int, yMax int) iter.Seq[Pair] {
	ret := it.Exhausted[Pair]()
	for ri := p.y - 1; ri <= p.y+1; ri++ {
		for ci := p.x - 1; ci <= p.x+width; ci++ {
			if (ri != p.y || ci < p.x || ci >= p.x+width) &&
				(ri >= 0 && ri < xMax) &&
				(ci >= 0 && ci < yMax) {
				ret = it.Chain(ret, it.Once(Pair{ci, ri}))
			}
		}
	}
	return ret
}

func isDigit(s string) bool {
	return len(s) > 0 && s[0] >= '0' && s[0] <= '9'
}

func isSymbol(s string) bool {
	return len(s) == 1 && !isDigit(s) && s[0] != '.'
}

type Part1 struct {
	xMax, yMax int
	schematic  map[Pair]string
	numbers    map[Pair]int
	symbols    map[Pair]rune
}

func part1() Part1 {
	ret := Part1{schematic: map[Pair]string{}, numbers: map[Pair]int{}, symbols: map[Pair]rune{}}

	lines := strings.Split(input, "\n")
	ret.yMax = len(lines)
	ret.xMax = len(lines[0])
	for ri, line := range lines {
		for p := range utils.PartitionFunc2([]byte(line), func(b byte) int {
			if b >= '0' && b <= '9' {
				return 0
			} else if b == '.' {
				return 1
			} else {
				return 2
			}
		}) {

			indexes, chars := itx.From2(p).Collect()
			value := string(chars)
			if isDigit(value) {
				ret.schematic[Pair{indexes[0], ri}] = value
				intValue, _ := strconv.Atoi(value)
				for _, ci := range indexes {
					ret.numbers[Pair{ci, ri}] = intValue
				}
			} else if isSymbol(value) {
				ret.schematic[Pair{indexes[0], ri}] = value
				ret.symbols[Pair{indexes[0], ri}] = rune(value[0])
			}
		}

	}

	sum := 0
	seen := map[string]Pair{}
	for pair, value := range ret.schematic {
		if !isSymbol(value) {
			for neighbor := range neighbors(pair, len(value), ret.xMax, ret.yMax) {
				//_, found := seen[value]
				if isSymbol(ret.schematic[neighbor]) /*&& !found*/ {
					n, _ := strconv.Atoi(value)
					sum += n
					seen[value] = pair
				}
			}
		}
	}
	println("part1", sum)
	return ret
}

func part2(result Part1) {
	sum := 0
	for pair, value := range result.schematic {
		if value[0] == '*' {
			ratio := map[int]bool{}
			for neighbor := range neighbors(pair, 1, result.xMax, result.yMax) {
				if n, numberFound := result.numbers[neighbor]; numberFound {
					ratio[n] = true
				}
			}
			if len(ratio) == 2 {
				sum += it.Fold(maps.Keys(ratio), func(agg int, i int) int {
					return agg * i
				}, 1)
			}
		}
	}
	println("part2", sum)
}

func main() {
	fmt.Printf("%v\n", slices.Collect(neighbors(Pair{5, 0}, 3, 10, 10)))
	result := part1()
	part2(result)
}
