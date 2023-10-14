package main

import (
	"os"

	"github.com/rodaine/table"
)

func WriteScoresToFile(s *BreedTestScoreSlice) {
	tbl := table.New("Breed", "Total", "Percent")
	os.WriteFile("output.txt", []byte{}, 0644)

	for _, b := range *s {
		tbl.AddRow(b.breedName, b.totalTested, b.percent)
	}
	tbl.WithWriter(&w{})
	tbl.Print()
}

type w struct {}

func (*w) Write(p []byte) (n int, err error) {
	f, e := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644) 
	
	if e != nil {
		return len(p), e
	}
	defer f.Close()

	n, er := f.Write(p)

	if er != nil {
		return len(p), e
	}


	return len(p), nil
}