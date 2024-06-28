package main

import (
  "fmt"
  "time"
)

type forceData struct {
	worker int
	wordList string
	url string
}

func executeQueryFromFile(currentPath chan string) {
  for eachtask := range currentPath{
      fmt.Printf("eachtask %s \n", eachtask)
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