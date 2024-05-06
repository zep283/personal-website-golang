package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	gofeed "github.com/mmcdole/gofeed"
)

type DummyStoryLinks struct {
	Links []DummyStory `json:"links"`
}

type DummyStory struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type Story struct {
	Title string
	Link  string
}

func ParseMediumRSSFeed() []Story {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://medium.com/feed/@zep283")
	var stories []Story
	for _, s := range feed.Items {
		story := Story{
			Title: s.Title,
			Link:  s.Link,
		}
		stories = append(stories, story)
	}
	if len(stories) < 5 {
		stories = dummyLinks(stories)
	}
	return stories
}

func dummyLinks(stories []Story) []Story {
	jsonFile, err := os.Open("./web/assets/dummy.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var dummyStories DummyStoryLinks
	json.Unmarshal(byteValue, &dummyStories)
	for _, s := range dummyStories.Links {
		story := Story(s)
		stories = append(stories, story)
	}
	return stories
}
