package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func aMonthAgo(time time.Time) time.Time {
	return time.AddDate(0, -1, 0)
}

func anYearAgo(time time.Time) time.Time {
	return time.AddDate(-1, 0, 0)
}

func isLessThanAMonthOld(item *Issue) bool {
	return item.CreatedAt.After(aMonthAgo(time.Now()))
}

func isLessThanAYearOld(item *Issue) bool {
	return item.CreatedAt.After(anYearAgo(time.Now()))
}

func isMoreThanAYearOld(item *Issue) bool {
	return item.CreatedAt.Before(anYearAgo(time.Now()))
}

func printIssue(issue *Issue) {
	fmt.Printf("#%-5d %9.9s %55.55s %v\n", issue.Number, issue.User.Login, issue.Title, issue.CreatedAt.Format("02-Jan-2006"))
}

func printLessThanAMonthOldIssues(issues []*Issue) {
	for _, issue := range issues {
		if isLessThanAMonthOld(issue) {
			printIssue(issue)
		}
	}
}

func printLessThanAYearOldIssues(issues []*Issue) {
	for _, issue := range issues {
		if isLessThanAYearOld(issue) {
			printIssue(issue)
		}
	}
}

func printMoreThanAYearOldIssues(issues []*Issue) {
	for _, issue := range issues {
		if isMoreThanAYearOld(issue) {
			printIssue(issue)
		}
	}
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	println("\nLess than a month old")
	printLessThanAMonthOldIssues(result.Items)
	println("\nLess than a year old")
	printLessThanAYearOldIssues(result.Items)
	println("\nMore than a year old")
	printMoreThanAYearOldIssues(result.Items)
}
