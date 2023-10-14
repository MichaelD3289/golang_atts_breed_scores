package main


func main() {
	// scrapeAndSaveToFile()
	s := scoresFromFile("scores.txt")
	s.filterOutLowTests(250)
	s.sortByHighestPercent()
	
	WriteScoresToFile(s)
}



