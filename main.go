package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments, please pass owner repo commitHash")
	}
	given := strings.Split(string(os.Args[1]), "/")

	owner := given[3]
	repo := given[4]
	commit := given[6]

	fmt.Println(fetchCommitDate(owner, repo, commit))
}

func fetchCommitDate(owner, repo, commitHash string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", owner, repo, commitHash)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var y commitResponse
	json.Unmarshal(data, &y)
	d, _ := time.Parse("2006-01-02T15:04:05Z", y.Committer.Date)

	return d.Format("2006010215064")
}

type commitResponse struct {
	Committer committer `json:"committer"`
}

type committer struct {
	Date string `json:"date"`
}
