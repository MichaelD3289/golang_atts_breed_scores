package main

import (
	"strconv"
	"strings"
)

type BreedTestScore struct {
	breedName string
	totalTested int
	passed int
	failed int
	percent float64
}

func CreateScoreFromSlice(s []string) *BreedTestScore {
	total, _ := strconv.Atoi(s[1])
	passed, _ := strconv.Atoi(s[2])
	failed, _ := strconv.Atoi(s[3])

	p := strings.Replace(s[4], "%", "", 1)
	percent, _ := strconv.ParseFloat(p, 64)
	return &BreedTestScore{
		breedName: s[0],
		totalTested: total,
		passed: passed,
		failed: failed,
		percent: percent,
	}
}