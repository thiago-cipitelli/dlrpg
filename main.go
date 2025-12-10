package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type CommitsQuery struct {
	User struct {
		ContributionsCollection struct {
			CommitContributionsByRepository []struct {
				Contributions struct {
					TotalCount int
				}
				Repository struct {
					Name string
				}
			}
		} `graphql:"contributionsCollection(from: $from, to: $to)"`
	} `graphql:"user(login: $login)"`
}

func queryTotalCommitsToday(client *githubv4.Client, username string) (CommitsQuery, error) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now().In(loc)
	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	to := from.Add(24 * time.Hour)

	var query CommitsQuery

	variables := map[string]interface{}{
		"login": githubv4.String(username),
		"from":  githubv4.DateTime{Time: from},
		"to":    githubv4.DateTime{Time: to},
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return query, err
	}
	return query, nil
}

func getTotalCommitsToday(query CommitsQuery) int {
	totalCommits := 0
	for _, repo := range query.User.ContributionsCollection.CommitContributionsByRepository {
		totalCommits += repo.Contributions.TotalCount
	}
	return totalCommits
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	username := os.Getenv("GITHUB_USER")
	if githubToken == "" || username == "" {
		log.Fatal("Defina GITHUB_TOKEN e GITHUB_USER no .env")
	}

	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	query, err := queryTotalCommitsToday(client, username)
	if err != nil {
		log.Fatal(err)
	}

	totalCommits := getTotalCommitsToday(query)

	fmt.Println("Commits hoje:", totalCommits)
}
