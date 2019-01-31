package validateuser_test

import (
	"net/http"
	"net/http/httptest"
	"testbadry/validateuser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUserInfo_not_found(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"message": "Not Found","documentation_url": "https://developer.github.com/v3/users/#get-a-single-user"}`))
	}))
	defer server.Close()
	githubAPI := validateuser.NewGithubAPI(server.Client(), server.URL)
	body, err := githubAPI.GetUserInfo("handsome")
	if assert.Error(t, err) {
		assert.Nil(t, body)
		assert.Equal(t, "Not Found", err.Error())
	}
}

func Test_GetUserInfo_fail_with_binding_error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{'login':"kaweel','id':14370817'}`))
	}))
	defer server.Close()
	githubAPI := validateuser.NewGithubAPI(server.Client(), server.URL)
	body, err := githubAPI.GetUserInfo("kaweel")
	if assert.Error(t, err) {
		assert.Nil(t, body)
		assert.Equal(t, "invalid character '\\'' looking for beginning of object key string", err.Error())
	}
}

func Test_GetUserInfo_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"login":"kaweel","id":14370817,"node_id":"MDQ6VXNlcjE0MzcwODE3","avatar_url":"https://avatars1.githubusercontent.com/u/14370817?v=4","gravatar_id":"","url":"https://api.github.com/users/kaweel","html_url":"https://github.com/kaweel","followers_url":"https://api.github.com/users/kaweel/followers","following_url":"https://api.github.com/users/kaweel/following{/other_user}","gists_url":"https://api.github.com/users/kaweel/gists{/gist_id}","starred_url":"https://api.github.com/users/kaweel/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/kaweel/subscriptions","organizations_url":"https://api.github.com/users/kaweel/orgs","repos_url":"https://api.github.com/users/kaweel/repos","events_url":"https://api.github.com/users/kaweel/events{/privacy}","received_events_url":"https://api.github.com/users/kaweel/received_events","type":"User","site_admin":false,"name":"Kawee Lertrungmongkol","company":null,"blog":"https://medium.com/@Kaweel","location":null,"email":null,"hireable":null,"bio":"Minimal Dev","public_repos":18,"public_gists":13,"followers":1,"following":4,"created_at":"2015-09-20T13:16:20Z","updated_at":"2019-01-24T04:13:57Z"}`))
	}))
	defer server.Close()
	githubAPI := validateuser.NewGithubAPI(server.Client(), server.URL)
	body, err := githubAPI.GetUserInfo("kaweel")
	if assert.NoError(t, err) {
		assert.Equal(t, "Kawee Lertrungmongkol", body.Name)
		assert.Equal(t, "https://medium.com/@Kaweel", body.Blog)
		assert.Equal(t, "Minimal Dev", body.Bio)
	}
}
