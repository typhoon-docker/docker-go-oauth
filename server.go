package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/imroc/req"
	"github.com/labstack/echo"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresAt    int    `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
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
		body := req.Param{
			"grant_type":    "authorization_code",
			"code":          c.QueryParam("code"),
			"redirect_uri":  os.Getenv("OAUTH_CALLBACK_URL"),
			"client_id":     os.Getenv("VIAREZO_CLIENT_ID"),
			"client_secret": os.Getenv("VIAREZO_CLIENT_SECRET"),
		}
		res, err := req.Post(
			fmt.Sprintf("%soauth/token", OAUTH_URL[c.QueryParam("state")]),
			body,
		)

		if err != nil {
		}

		fmt.Println(res)

		token := Token{} // empty token
		err = res.ToJSON(&token)
		if err != nil {
			panic(err)
		}

		return c.JSONPretty(http.StatusOK, token, "    ")
	})
	e.GET("/login_viarezo", func(c echo.Context) error {
		url, err := GetCode("viarezo")
		if err != nil {
		}
		fmt.Println(url)
		return c.Redirect(http.StatusMovedPermanently, url)
	})
	e.GET("/login_github", func(c echo.Context) error {
		url, err := GetCode("github")
		if err != nil {
		}
		fmt.Println(url)
		return c.Redirect(http.StatusMovedPermanently, url)
	})
	e.Logger.Fatal(e.Start(":80"))
}
