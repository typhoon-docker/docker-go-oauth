package main

import (
	"fmt"
	"os"
	
	"net/http"
	"net/url"
	"encoding/json"
	
	"github.com/labstack/echo"
)


type Token struct {
	access_token string
	expires_at int
	expires_in int
	refresh_token string
	scope string
}

func init() {
	LoadEnv()
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/callback", func(c echo.Context) error {
		// fmt.Println(c.QueryParam("state"))
		token := Token{} // empty token

		resp, err := http.PostForm(
			fmt.Sprintf("%soauth/token", os.Getenv("OAUTH_URL")),
			url.Values {
				"grant_type": {"authorization_code"},
				"code": {c.QueryParam("code")},
				"redirect_uri": {os.Getenv("CALLBACK_URL")},
				"client_id": {os.Getenv("CLIENT_ID")},
				"client_secret": {os.Getenv("CLIENT_SECRET")},
			})

		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		
		err = json.NewDecoder(resp.Body).Decode(&token)
		if err != nil {}

		fmt.Println(token)
		fmt.Println(token.access_token)
		fmt.Println(token.expires_at)
		fmt.Println(token.expires_in)
		fmt.Println(token.refresh_token)
		fmt.Println(token.scope)

		return c.String(http.StatusOK, "Callback")
	})
	e.GET("/login", func(c echo.Context) error {
 		return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%soauth/authorize/?redirect_uri=%s&client_id=%s&response_type=code&state=viarezo&scope=default", os.Getenv("OAUTH_URL"), os.Getenv("CALLBACK_URL"), os.Getenv("CLIENT_ID")))
	})
	e.Logger.Fatal(e.Start(":80"))
}
