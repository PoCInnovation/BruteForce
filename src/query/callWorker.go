package query

import (
	"bruteforce/src/matching"
	"bruteforce/src/models"
	"bruteforce/src/utils"
	"sync"
)


func executeQueryFromFile(wg *sync.WaitGroup, params *models.Forcing_params, currentPath chan string) {
	defer wg.Done()
	for taskData := range currentPath {
		QueryExecute(params, taskData, "POST")
	}
}

func MainRequest(params *models.Forcing_params, criteria matcher.MatchCriteria) {
	wg := &sync.WaitGroup{}
    wg.Add(data.Worker)
	channel := make(chan string)
	wordArray := utils.GetFileContent(params.Wordlist)

	for i := 0; i < params.Workers; i++ {
		go executeQueryFromFile(wg, params, channel)
	}
	for i := 0; i < len(wordArray); i++ {
		channel <- wordArray[i]
	}
	wg.Wait()
	close(channel)
}
