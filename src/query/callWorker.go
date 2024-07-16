package main

import (
  "time"
)

type forceData struct {
	worker int
	wordList string
	url string
}

func executeQueryFromFile(data forceData, currentPath chan string) {
  for taskData := range currentPath{
    queryExecute(data, taskData, "POST")
    }
}

func mainRequest(data forceData) {
  channel := make(chan string)
  wordArray := GetFileContent("../wordList/rootList")

  for i := 0 ;i < data.worker; i++ {
		go executeQueryFromFile(data, channel)
	}
  for i := 0; i < len(wordArray); i++ {
    channel <- wordArray[i]
  }
  time.Sleep(1 * time.Second)
  close(channel)
}

