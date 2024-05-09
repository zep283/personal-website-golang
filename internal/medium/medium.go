package medium

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	gofeed "github.com/mmcdole/gofeed"
	"github.com/zep283/personal-website-golang/internal/common"
)

type DummyStoryLinks struct {
	Links []DummyStory `json:"links"`
}

type DummyStory struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func ParseMediumRSSFeed() []common.Story {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://medium.com/feed/@zep283")
	var stories []common.Story
	for _, s := range feed.Items {
		story := common.Story{
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

func dummyLinks(stories []common.Story) []common.Story {
	jsonFile, err := os.Open("../web/assets/dummy.json")
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
		story := common.Story(s)
		stories = append(stories, story)
	}
	return stories
}
