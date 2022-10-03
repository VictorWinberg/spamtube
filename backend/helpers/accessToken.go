package helpers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"spamtube/backend/domain"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

func GetAccessToken(c *cache.Cache) (*domain.AccessToken, error) {
	client := &http.Client{}
	URL := "https://www.reddit.com/api/v1/access_token"
	v := url.Values{
		"grant_type": {"client_credentials"},
		"username":   {os.Getenv("REDDIT_USERNAME")},
		"password":   {os.Getenv("REDDIT_PASSWORD")},
	}
	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Getenv("REDDIT_APP_USERNAME"), os.Getenv("REDDIT_APP_PRIVATE_KEY"))
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.Wrap(err, "Could not retrieve access token from reddit")
	}

	token := &domain.AccessToken{}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&token)

	c.Set("token", token.AccessToken, cache.DefaultExpiration)

	return token, nil
}

func HandleTokenLogic(c *cache.Cache, con *gin.Context) (string, error) {
	var token string
	if val, found := c.Get("token"); found {
		token = val.(string)
	} else {
		t, err := GetAccessToken(c)
		if err != nil {
			return "", err
		}
		token = t.AccessToken
	}
	return token, nil
}
