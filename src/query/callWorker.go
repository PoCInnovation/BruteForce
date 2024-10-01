package query

import (
	"bruteforce/src/models"
	"bruteforce/src/utils"
	"sync"
)

func executeQueryFromFile(wg *sync.WaitGroup, params *models.ForcingParams, currentPath chan string) {
	defer wg.Done()

	if params.Url[len(params.Url)-1] != '/' {
		params.Url = params.Url + "/"
	}

	for taskData := range currentPath {
		QueryExecute(params, taskData, params.Method)
	}
}

func MainRequest(params *models.ForcingParams) {
	wg := &sync.WaitGroup{}
	wg.Add(params.Workers)
	channel := make(chan string)
	wordArray := utils.GetFileContent(params.Wordlist)

	for i := 0; i < params.Workers; i++ {
		go executeQueryFromFile(wg, params, channel)
	}

	for i := 0; i < len(wordArray); i++ {
		if len(wordArray[i]) > 0 && wordArray[i][0] != '#' {
			channel <- wordArray[i]
		}
	}
	close(channel)
	wg.Wait()
}
