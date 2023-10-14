package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type BreedTestScoreSlice []BreedTestScore

func (btss *BreedTestScoreSlice) filterOutLowTests(num int) {
	var highTesters BreedTestScoreSlice

	for _, b := range *btss {
		if b.totalTested >= num {
			highTesters = append(highTesters, b)
		}
	}

	*btss = highTesters
}

func (btss *BreedTestScoreSlice) sortByHighestPercent() {
	sort.Slice(*btss, func(i, j int) bool {
		return (*btss)[i].percent > (*btss)[j].percent
	})
}

func (b BreedTestScoreSlice) createStringSlice() string {
	docString := ""
	for _, s := range b {
		f := fmt.Sprintf("%s::%v::%v::%v::%v>>>", s.breedName, s.totalTested, s.passed, s.failed, s.percent)
		docString += f
	}
	return docString
}

func (b *BreedTestScoreSlice) createByteSlice() []byte {
	s := b.createStringSlice()
	return []byte(s) 
}

func (b *BreedTestScoreSlice) saveToFile(fileName string) error {
	err := os.WriteFile(fileName, b.createByteSlice(), 0644)
	return err
}

func scoresFromFile(fileName string) *BreedTestScoreSlice {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error:", err)
	}
	str := string(f)
	sli := strings.Split(str, ">>>")
	btss := BreedTestScoreSlice{}
	for _, dogS := range sli {
		d := strings.Split(dogS, "::")
		if len(d) == 5 {
			btss = append(btss, *CreateScoreFromSlice(d))
		}
	}
	return &btss
}