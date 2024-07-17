package query

import (
	"bruteforce/src/matching"
	"bruteforce/src/utils"
	"time"
)

type ForceData struct {
	Worker   int
	WordList string
	Url      string
}

func executeQueryFromFile(data ForceData, currentPath chan string) {
	for taskData := range currentPath {
		QueryExecute(data, taskData, "POST")
	}
}

func MainRequest(data ForceData, criteria matcher.MatchCriteria) {
	channel := make(chan string)
	wordArray := utils.GetFileContent(data.WordList)

	for i := 0; i < data.Worker; i++ {
		go executeQueryFromFile(data, channel)
	}
	for i := 0; i < len(wordArray); i++ {
		channel <- wordArray[i]
	}
	time.Sleep(1 * time.Second)
	close(channel)
}
