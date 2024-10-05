package scarppingTech

import (
	"bruteforce/src/models"
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func analyzeScripts(params *models.ForcingParams, scripts []string) {
	TechKeywords, _ := readLines("keywordDetection/Techkeywords.txt")

	for _, script := range scripts {
		for _, keyword := range TechKeywords {
			if strings.Contains(script, keyword) {
				params.PromptInfo.TechKeywords = append(params.PromptInfo.TechKeywords, keyword)
			}
		}
	}
}

func ScrapScripts(params *models.ForcingParams, url string) {
	c := colly.NewCollector()

	var scriptURLs []string

	c.OnHTML("script", func(e *colly.HTMLElement) {
		scriptSrc := e.Attr("src")
		if scriptSrc != "" {
			scriptURLs = append(scriptURLs, scriptSrc)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		analyzeScripts(params, scriptURLs)
	})

	err := c.Visit(url)
	if err != nil {
		log.Println("Error visiting the site:", err)
	}
}
