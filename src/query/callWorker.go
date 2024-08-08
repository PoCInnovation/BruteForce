package query

import (
	"bruteforce/src/matching"
	"bruteforce/src/utils"
	"sync"
)

type ForceData struct {
	Worker   int
	WordList string
	Url      string
}

func executeQueryFromFile(wg *sync.WaitGroup, data ForceData, currentPath chan string) {
	defer wg.Done()
	for taskData := range currentPath {
		QueryExecute(data, taskData, "POST")
	}
}

func MainRequest(data ForceData, criteria matcher.MatchCriteria) {
	wg := &sync.WaitGroup{}
    wg.Add(data.Worker)
	channel := make(chan string)
	wordArray := utils.GetFileContent(data.WordList)

	for i := 0; i < data.Worker; i++ {
		go executeQueryFromFile(wg, data, channel)
	}
	for i := 0; i < len(wordArray); i++ {
		channel <- wordArray[i]
	}
	wg.Wait()
	close(channel)
}
