package main

import (
	"fmt"
	"os"
	
	"net/http"

	"github.com/imroc/req"
	"github.com/labstack/echo"
)


type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt int `json:"expires_at"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
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

		body := req.Param {
			"grant_type": "authorization_code",
			"code": c.QueryParam("code"),
			"redirect_uri": os.Getenv("CALLBACK_URL"),
			"client_id": os.Getenv("CLIENT_ID"),
			"client_secret": os.Getenv("CLIENT_SECRET"),
		}
		res, err := req.Post(
			fmt.Sprintf("%soauth/token", os.Getenv("OAUTH_URL")),
			body,
		)

		if err != nil {}

		fmt.Println(res)
		
		token := Token {} // empty token
		err = res.ToJSON(&token)
		if err != nil {
			panic(err)
		}

		return c.JSONPretty(http.StatusOK, token, "    ")
	})
	e.GET("/login", func(c echo.Context) error {
 		return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%soauth/authorize/?redirect_uri=%s&client_id=%s&response_type=code&state=viarezo&scope=default", os.Getenv("OAUTH_URL"), os.Getenv("CALLBACK_URL"), os.Getenv("CLIENT_ID")))
	})
	e.Logger.Fatal(e.Start(":80"))
}
