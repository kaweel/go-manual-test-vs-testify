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

	fmt.Printf("login '%v' using name '%v' ?? : %v\n", username, name, resp)

}
