package main

import (
	"fmt"
	"log"
	"net/http"
	"testbadry/validateuser"
	"time"
)

func main() {

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	var username = "kaweel"
	var name = "Kawee Lertrungmongkol"

	githubAPI := validateuser.NewGithubAPI(httpClient, "https://api.github.com/users")
	service := validateuser.NewValidateUserService(githubAPI)

	resp, err := service.ValidateUser(username, name)
	if err != nil {
		log.Fatalf("error is %v", err)
	}
	fmt.Printf("Q: user '%v' still used name '%v' ??\n", username, name)
	fmt.Printf("A: anwser is %v\n", resp)
}
