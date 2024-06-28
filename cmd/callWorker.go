package main

import (
  "time"
)

type forceData struct {
	worker int
	wordList string
	url string
}

func executeQueryFromFile(currentPath chan string) {
  for taskData := range currentPath{
    queryExecute(taskData, "POST")
    }
}

func mainRequest(worker int) {
  channel := make(chan string)
  wordArray := GetFileContent("../wordList/rootList")

  for i := 0 ;i < worker; i++ {
		go executeQueryFromFile(channel)
	}
  for i := 0; i < len(wordArray); i++ {
    channel <- wordArray[i]
  }
  time.Sleep(1 * time.Second)
  close(channel)
}

func main () {
  mainRequest(3);
}