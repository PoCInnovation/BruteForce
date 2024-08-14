package query

import (
	"bruteforce/src/models"
	"bruteforce/src/utils"
	"sync"
)

func executeQueryFromFile(wg *sync.WaitGroup, params *models.Forcing_params, currentPath chan string) {
	defer wg.Done()
	for taskData := range currentPath {
		QueryExecute(params, taskData, "GET")
	}
}

func MainRequest(params *models.Forcing_params) {
	wg := &sync.WaitGroup{}
	wg.Add(params.Workers)
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
