
import (
  "fmt"
  "time"
  "strings"
)

type forceData struct {
	worker int
	wordList string
	url string
}

func executeQueryFromFile(filePathQuery []string, filePathRoot chan string, worker int) {
  for eachtask := range filePathRoot {
    for i := 0 ; i < worker ; i++ {
      fmt.Println(splitData)
    }
  }
}

func mainRequest(worker int) {
  channel := make(chan string)
  wordArray := GetFileContent("./wordList/rootList")

  for i := 0 ;i < len(data_tab); i++ {
		go executeQueryFromFile(wordArray, channel)
	}

  for i := 0; i < len(data_tab); i++ {
    channel <- data_tab[i]
  }
  time.Sleep(1 * time.Second)
  close(channel)
}
