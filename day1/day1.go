package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

var numWords = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
	"six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func findFirstNumber(line string) int {
	for i := 0; i < len(line); i++ {
		s := line[i:]
		if strings.IndexAny(s[:1], "0123456789") != -1 {
			if len(s[:1]) != 1 {
				panic("Must compare 1!")
			}
			return int(s[0] - '0')
		}

		for word, n := range numWords {
			if strings.Index(s, word) == 0 {
				return n
			}
		}
	}
	panic("Shouldn't happen")
}

func findLastNumber(line string) int {
	for i := 0; i < len(line); i++ {
		s := line[:len(line)-i]
		if strings.LastIndexAny(s[len(s)-1:], "0123456789") != -1 {
			if len(s[len(s)-1:]) != 1 {
				panic("Must compare 1!")
			}
			return int(s[len(s)-1] - '0')
		}

		sWord := s
		if len(sWord) > 5 {
			sWord = sWord[len(s)-5:]
		}
		for word, n := range numWords {
			idx := strings.LastIndex(sWord, word)
			if idx != -1 && idx == len(sWord)-len(word) {
				return n
			}
		}
	}
	panic("Shouldn't happen")
}

func part1() []int {
	lines := strings.Split(input, "\n")
	ret := make([]int, len(lines))
	for i, line := range lines {
		xIdx := strings.IndexAny(line, "0123456789")
		yIdx := strings.LastIndexAny(line, "0123456789")
		x := int(line[xIdx] - '0')
		y := int(line[yIdx] - '0')
		ret[i] = x*10 + y
	}
	return ret
}

func part2() []int {
	lines := strings.Split(input, "\n")
	ret := make([]int, len(lines))
	for i, line := range lines {
		x := findFirstNumber(line)
		y := findLastNumber(line)
		ret[i] = x*10 + y
	}
	return ret
}

func main() {
	nums := part1()
	part1total := 0
	for _, num := range nums {
		part1total += num
	}
	println(part1total)

	nums = part2()
	part2Total := 0
	for _, num := range nums {
		part2Total += num
	}
	println(part2Total)
}
