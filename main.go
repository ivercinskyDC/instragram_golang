package main

import (
	"net/http"

	"golang.org/x/oauth2"
)

func main() {
	igConf = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		RedirectURL: RedirectURI,
		Scopes:      []string{"public_content", "comments"},
	}

	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/profile", profilePage)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
