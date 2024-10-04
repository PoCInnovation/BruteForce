package wordlistgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"bruteforce/src/models"
)

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

const WORDLIST_SYS_PROMPT string = "You are a wordlist generator knowledgeable about technology that generates a list and only a list of keywords and no - before the words, the words must follow the camel case and have no trailling spaces"

func CallopenAI(SysPrompt string, UserPrompt string) (error, string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenAI API key not found in environment")
	}

	url := "https://api.openai.com/v1/chat/completions"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "system", "content": SysPrompt},
			{"role": "user", "content": UserPrompt},
		},
		"max_tokens": 200,
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err, ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err, ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err, ""
	}

	var openAIResponse OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err, ""
	}

	if len(openAIResponse.Choices) > 0 {
		return nil, openAIResponse.Choices[0].Message.Content
	}
	return fmt.Errorf("No choices found in the response."), ""
}

func BuildPrompt(PromptInfo *models.PromptBuilder) (syspromt string, userpromt string) {
	SysPrompt := WORDLIST_SYS_PROMPT
	UserPrompt := fmt.Sprintf("Generate a wordlist of %d terms", PromptInfo.WordlistLen)

	if len(PromptInfo.SiteKeywords) > 0 {
		UserPrompt += " related to theses tech stack including"
		for _, techkeyword := range PromptInfo.TechKeywords {
			UserPrompt += techkeyword
		}
	}
	if len(PromptInfo.SiteKeywords) > 0 {
		UserPrompt += " and also terms related to theses topics"
		for _, sitekeyword := range PromptInfo.SiteKeywords {
			UserPrompt += sitekeyword
		}
	}
	return SysPrompt, UserPrompt
}

const BRUTEFORCE_DIR string = "BruteforceWordlists"

func generateNewFile() (*os.File, string) {
	re := regexp.MustCompile("[0-9]+")

	dirsinfo, err := os.ReadDir(BRUTEFORCE_DIR)
	if err != nil {
		err := os.Mkdir(BRUTEFORCE_DIR, os.ModePerm)
		if err != nil {
			panic(err)
		}
		dirsinfo, _ = os.ReadDir(BRUTEFORCE_DIR)
	}
	maxNumber := 0
	if len(dirsinfo) > 0 {
		for _, file := range dirsinfo {
			numstr := re.FindAllString(file.Name(), -1)
			if len(numstr) >= 1 {
				num, _ := strconv.Atoi(numstr[0])
				if num > maxNumber {
					maxNumber = num
				}
			}
		}
	}
	maxNumber += 1
	newfile := fmt.Sprintf("%s/gptwordlist%d.txt",
		BRUTEFORCE_DIR, maxNumber)
	f, err := os.OpenFile(newfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return f, newfile
}

func OverwriteWordlist(params *models.ForcingParams) {
	if params.BoolFlags.Generate == false {
		return
	}

	if params.BoolFlags.Verbose {
		log.Print("Generating wordlist ...")
	}
	f, newfile := generateNewFile()

	params.Wordlist = newfile
	userprompt, sysprompt := BuildPrompt(&params.PromptInfo)

	if params.BoolFlags.Verbose {
		log.Print("Contacting OpenAI API ...")
	}
	err, wordlist := CallopenAI(sysprompt, userprompt)

	f.WriteString(wordlist)
	f.Close()
	if params.BoolFlags.Verbose {
		log.Print("Clearing trailling spaces ...")
	}
	cmd := exec.Command("/bin/sh", "-c", "sed -i 's/[ \\t]*$//' "+newfile)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
	if params.BoolFlags.Verbose {
		log.Print("Done generating wordlist ...")
	}
}
