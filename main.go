package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func listStories(storyType string) (storyIds []int32, err error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/%sstories.json", storyType)

	client := http.Client{Timeout: 500 * time.Millisecond}
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

func getStory(storyId int32) (story map[string]interface{}, err error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", storyId)
	client := http.Client{Timeout: 500 * time.Millisecond}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// FIXME: unmarshall into a struct?
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func printStory(story map[string]interface{}) {
	fmt.Printf("%s ", story["title"])
	fmt.Printf("[%.f]\n", story["score"])
	// FIXME: url can be nil (e.g., LaunchHN, etc.)
	fmt.Printf("%s\n", story["url"])
	fmt.Printf("https://news.ycombinator.com/item?id=%.f\n", story["id"])
}

func main() {
	maxResults := flag.Int("n", 5, "Chose randomly from this many top results")
	flag.Parse()

	storyIds, err := listStories("top")
	if err != nil {
		panic(err)
	}

	selectionSize := min(*maxResults, len(storyIds))
	storyId := storyIds[rand.Intn(selectionSize)]
	story, err := getStory(storyId)
	if err != nil {
		panic(err)
	}

	printStory(story)
}
