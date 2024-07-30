package query

import (
	"bruteforce/src/matching"
	"bruteforce/src/models"
	"bruteforce/src/utils"
	"time"
)

func executeQueryFromFile(params *models.Forcing_params, currentPath chan string) {
	for taskData := range currentPath {
		QueryExecute(params, taskData, "POST")
	}
}

func MainRequest(params *models.Forcing_params, criteria matcher.MatchCriteria) {
	channel := make(chan string)
	wordArray := utils.GetFileContent(params.Wordlist)

	for i := 0; i < params.Workers; i++ {
		go executeQueryFromFile(params, channel)
	}
	for i := 0; i < len(wordArray); i++ {
		channel <- wordArray[i]
	}
	time.Sleep(1 * time.Second)
	close(channel)
}
