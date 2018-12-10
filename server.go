package main

import (
	"fmt"
	"os"
	"strings"

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
	e.GET("/callback/viarezo", func(c echo.Context) error {
		oauth := c.QueryParam("state")
		body := req.Param{
			"grant_type":    "authorization_code",
			"code":          c.QueryParam("code"),
			"redirect_uri":  fmt.Sprintf("%s/%s", os.Getenv("OAUTH_CALLBACK_URL"), oauth),
			"client_id":     os.Getenv(fmt.Sprintf("%s_CLIENT_ID", strings.ToUpper(oauth))),
			"client_secret": os.Getenv(fmt.Sprintf("%s_CLIENT_SECRET", strings.ToUpper(oauth))),
		}

		res, err := req.Post(
			OAUTH_TOKEN_URL[OAUTH[oauth]],
			body,
		)

		if err != nil {
			return err
		}

		token := Token{} // empty token
		err = res.ToJSON(&token)
		if err != nil {
			return err
		}

		return c.JSONPretty(http.StatusOK, token, "    ")
	})
	e.GET("/callback/github", func(c echo.Context) error {
		oauth := c.QueryParam("state")
		body := req.Param{
			"code":          c.QueryParam("code"),
			"redirect_uri":  fmt.Sprintf("%s/%s", os.Getenv("OAUTH_CALLBACK_URL"), oauth),
			"client_id":     os.Getenv(fmt.Sprintf("%s_CLIENT_ID", strings.ToUpper(oauth))),
			"client_secret": os.Getenv(fmt.Sprintf("%s_CLIENT_SECRET", strings.ToUpper(oauth))),
			"state":         oauth,
		}

		res, err := req.Post(
			OAUTH_TOKEN_URL[OAUTH[oauth]],
			req.BodyJSON(&body),
		)

		if err != nil {
			return err
		}

		token, err := res.ToString()

		return c.String(http.StatusOK, token)
	})
	e.GET("/login/viarezo", func(c echo.Context) error {
		url, err := GetCode("viarezo")
		if err != nil {
		}
		fmt.Println(url)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})
	e.GET("/login/github", func(c echo.Context) error {
		url, err := GetCode("github")
		if err != nil {
		}
		fmt.Println(url)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})
	e.Logger.Fatal(e.Start(":80"))
}
