package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zep283/personal-website-golang/internal/common"
)

func GetRepos() []common.Repo {
	resp, err := http.Get("https://api.github.com/users/zep283/repos")

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Print(sb)

	var jsonMap []map[string]interface{}
	json.Unmarshal([]byte(sb), &jsonMap)

	repos := []common.Repo{}

	for _, repo := range jsonMap {
		repoName := repo["name"].(string)
		repoDesc := repo["description"].(string)
		repoUrl := repo["url"].(string)
		repoInfo := common.Repo{
			Name:        repoName,
			Description: repoDesc,
			Url:         repoUrl,
		}
		repos = append(repos, repoInfo)
	}
	return repos
}
