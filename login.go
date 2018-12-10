package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	viarezo = iota
	github
)

var (
	OAUTH = map[string]int{
		"viarezo": viarezo,
		"github":  github,
	}

	OAUTH_URL = map[int]string{
		viarezo: "https://auth.viarezo.fr/",
		github:  "https://github.com/login/",
	}

	OAUTH_AUTHORIZE_URL = map[int]string{
		viarezo: "https://auth.viarezo.fr/oauth/authorize",
		github:  "https://github.com/login/oauth/authorize",
	}

	OAUTH_TOKEN_URL = map[int]string{
		viarezo: "https://auth.viarezo.fr/oauth/token",
		github:  "https://github.com/login/oauth/access_token",
	}
)

func getOauth() ([]string, map[int]url.Values) {
	queries := map[int]url.Values{
		viarezo: url.Values{
			"scope":         {"default"},
			"response_type": {"code"},
		},
		github: url.Values{},
	}
	return nil, queries
}

func GetCode(oauth string) (string, error) {
	_, queries := getOauth()
	oauthID := OAUTH[oauth]
	query, ok := queries[oauthID]
	fmt.Println(query, ok)
	if !ok {
		return "", errors.New("oauth doesn't exist")
	}

	query.Add("redirect_uri", fmt.Sprintf("%s/%s", os.Getenv("OAUTH_CALLBACK_URL"), oauth))
	query.Add("client_id", os.Getenv(fmt.Sprintf("%s_CLIENT_ID", strings.ToUpper(oauth))))
	query.Add("state", oauth)
	return fmt.Sprintf(
		"%s?%s",
		OAUTH_AUTHORIZE_URL[oauthID],
		query.Encode()), nil
}
