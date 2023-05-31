package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Story struct {
	By    string
	Time  string
	Title string
	Score int
	Url   string
	Id    int
}

const DefaultTemplate = `{{.Title}} [{{.Score}}]
{{.Url}}
https://news.ycombinator.com/item?id={{.Id}}
`

func listStories(storyType string, timeout time.Duration) ([]int32, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/%sstories.json", storyType)

	client := http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []int32
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getStory(storyId int32, timeout time.Duration) (*Story, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", storyId)
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	story := &Story{}
	json.NewDecoder(resp.Body).Decode(story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func printStory(story *Story, tmpl *template.Template) error {
	err := tmpl.Execute(os.Stdout, story)
	return err
}

func loadTemplate(tmplText string) (*template.Template, error) {
	if len(tmplText) == 0 {
		tmplText = DefaultTemplate
	}
	tmpl, err := template.New("hnwelcome").Parse(tmplText)
	return tmpl, err
}

func main() {
	newest := flag.Bool("newest", false, "Show newest stories (deafult is to show current top stories)")
	maxResults := flag.Int("n", 5, "Chose randomly from this many top results")
	timeoutText := flag.String("timeout", "1s", "Timeout duration, as text. Valid time units are m, s, ms, etc.")
	tmplText := flag.String("template", "", "Output formatting template")
	flag.Parse()

	timeout, err := time.ParseDuration(*timeoutText)
	if err != nil {
		panic(err)
	}
	timeout = timeout / 2 // there is a total of 2 requests made hereafter

	tmpl, err := loadTemplate(*tmplText)
	if err != nil {
		panic(err)
	}

	storyType := "top"
	if *newest {
		storyType = "new"
	}
	storyIds, err := listStories(storyType, timeout)
	if err != nil {
		panic(err)
	}

	selectionSize := *maxResults
	if len(storyIds) < *maxResults {
		selectionSize = len(storyIds)
	}
	storyId := storyIds[rand.Intn(selectionSize)]
	story, err := getStory(storyId, timeout)
	if err != nil {
		panic(err)
	}

	err = printStory(story, tmpl)
	if err != nil {
		panic(err)
	}
}
