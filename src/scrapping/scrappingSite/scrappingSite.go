package scarppingSite

import (
	"bruteforce/src/models"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func loadStopWords(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stopWords := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			stopWords[word] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stopWords, nil
}

func cleanText(text string) string {
	reg := regexp.MustCompile("[^a-zA-Z ]+")
	cleaned := reg.ReplaceAllString(strings.ToLower(text), "")
	spaceReg := regexp.MustCompile("\\s+")
	return spaceReg.ReplaceAllString(cleaned, " ")
}

func countWords(text string, stopWords map[string]bool) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		if !stopWords[word] && len(word) > 2 {
			wordCount[word]++
		}
	}
	return wordCount
}

func ScrapSite(params *models.ForcingParams, url string) {
	stopWords, err := loadStopWords("keywordDetection/stopwords.txt")
	if err != nil {
		fmt.Println("Error loading stop words:", err)
		return
	}

	c := colly.NewCollector()

	var siteText string

	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		description := strings.TrimSpace(e.Attr("content"))
		if params.BoolFlags.Verbose {
			fmt.Println("Meta Description:", description)
		}
		siteText += description + " "
	})

	c.OnHTML("meta[name=keywords]", func(e *colly.HTMLElement) {
		keywords := strings.TrimSpace(e.Attr("content"))
		if params.BoolFlags.Verbose {
			fmt.Println("Meta Keywords:", keywords)
		}
		siteText += keywords + " "
	})

	c.OnHTML("h1, h2, h3", func(e *colly.HTMLElement) {
		heading := strings.TrimSpace(e.Text)
		if params.BoolFlags.Verbose {
			fmt.Println("Heading:", heading)
		}
		siteText += heading + " "
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		paragraph := strings.TrimSpace(e.Text)
		siteText += paragraph + " "
	})

	c.Visit(url)

	cleanedText := cleanText(siteText)

	wordCount := countWords(cleanedText, stopWords)

	for word := range wordCount {
		params.PromptInfo.SiteKeywords = append(params.PromptInfo.SiteKeywords, word)
	}
}
