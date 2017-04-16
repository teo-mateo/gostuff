package impl

import "fmt"
import "time"
import "net/url"
import "strings"
import "net/http"
import "encoding/json"

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login	string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number 		int
	HTMLURL 	string 		`json:"html_url"`
	Title		string
	State		string
	User 		*User
	CreatedAt	time.Time	`json:"created_at"` 
	Body		string		// in Markdown format
}

type IssueSearchResult struct {
	TotalCount 	int `json:"total_count"`
	Items		[]*Issue
}

func SearchIssues(ageMonths int, terms []string) (*IssueSearchResult, error){
	
	t := time.Now().AddDate(0, -ageMonths, 0)
	aMonthAgo := "created:>=" + t.Format("2006-01-02")
	query := strings.Join(terms, " ")
	query = strings.Join([]string {query, aMonthAgo}, " ")
	fmt.Println(query)
	query = url.QueryEscape(query)
	query = IssuesURL + "?q=" + query
	fmt.Println(IssuesURL + "?q=" + query)
	

	resp, err := http.Get(query)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
