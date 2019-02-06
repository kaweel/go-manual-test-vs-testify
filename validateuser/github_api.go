package validateuser

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GithubAPI interface {
	GetUserInfo(user string) (*UserInformation, error)
}

type api struct {
	httpClient *http.Client
	baseURL    string
}

func NewGithubAPI(httpClient *http.Client, baseURL string) GithubAPI {
	return &api{httpClient, baseURL}
}

func (api *api) GetUserInfo(user string) (*UserInformation, error) {
	req, err := http.NewRequest("GET", api.baseURL+"/"+user, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("Not Found")
	}

	record := &UserInformation{}
	if err := json.NewDecoder(res.Body).Decode(record); err != nil {
		return nil, err
	}
	return record, nil
}

type UserInformation struct {
	Login       string    `json:"login"`
	ID          int       `json:"id"`
	NodeID      string    `json:"node_id"`
	Type        string    `json:"type"`
	SiteAdmin   bool      `json:"site_admin"`
	Name        string    `json:"name"`
	Blog        string    `json:"blog"`
	Bio         string    `json:"bio"`
	PublicRepos int       `json:"public_repos"`
	PublicGists int       `json:"public_gists"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
