package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func getOauth() ([2]string, map[string]url.Values) {
	viarezo := "viarezo"
	github := "github"

	oauths := [2]string{viarezo, github}

	queries := map[string]url.Values{
		viarezo: url.Values{
			"scope":         {"default"},
			"response_type": {"code"},
		},
		github: url.Values{},
	}
	return oauths, queries
}

func GetCode(oauth string) (string, error) {
	_, queries := getOauth()
	query, ok := queries[oauth]
	fmt.Println(query, ok)
	if !ok {
		return "", errors.New("oauth doesn't exist")
	}

	query.Add("redirect_uri", os.Getenv(fmt.Sprintf("%s_CALLBACK_URL", strings.ToUpper(oauth))))
	query.Add("client_id", os.Getenv(fmt.Sprintf("%s_CLIENT_ID", strings.ToUpper(oauth))))
	query.Add("state", oauth)
	return fmt.Sprintf(
		"%soauth/authorize/?%s",
		os.Getenv(fmt.Sprintf("%s_OAUTH_URL", strings.ToUpper(oauth))),
		query.Encode()), nil
}
