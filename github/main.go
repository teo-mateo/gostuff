package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	github "github.com/teo-mateo/gostuff/github/impl"
)

func main(){

	ageMonths, _ := strconv.ParseInt(os.Args[1], 10, 8)
	result, err := github.SearchIssues(int(ageMonths), os.Args[2:])
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %10.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}