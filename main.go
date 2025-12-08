package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()
	private_key := os.Getenv("GITHUB_APP_PRIVATE_KEY")

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(private_key)})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	events, _, err := client.Activity.ListEventsPerformedByUser(ctx, "thiago-cipitelli", true, nil)
	if err != nil {
		fmt.Printf("deu erro")
		log.Fatal()
	}

	today := time.Now().UTC().Format("2006-01-02")

	for _, e := range events {
		if *e.Type == "PushEvent" {
			fmt.Println(e.CreatedAt.Format("2006-01-02"))
			if e.CreatedAt.Format("2006-01-02") == today {
				fmt.Println(string(e.GetRawPayload()))
				fmt.Printf("oi")
			}
		}
	}
}
