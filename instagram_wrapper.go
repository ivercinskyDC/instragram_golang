package main

import "net/http"
import "io/ioutil"
import "fmt"
import "html/template"
import "golang.org/x/oauth2"
import "encoding/json"

//ClientID from Instagram
var ClientID = "a9e0f62b9ab844cda85194bfcaa649df"

//ClientSecret from Instragram
var ClientSecret = "2b4cd2664a2e405faec355a8e9f6a3fc"

//RedirectURI from Instragram
var RedirectURI = "http://localhost:8080/redirect"

var authURL = "https://api.instagram.com/oauth/authorize"

var tokenURL = "https://api.instagram.com/oauth/access_token"

var templ = template.Must(template.New("index2.html").ParseFiles("index2.html"))

var profile = template.Must(template.New("profile.html").ParseFiles("profile.html"))

var profileResponse ProfileResponse

var igConf *oauth2.Config

//Profile response form Instagram
type Profile struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_Picture"`
}

//ProfileResponse response from Instagram
type ProfileResponse struct {
	Profile Profile `json:"data"`
}

func redirect(res http.ResponseWriter, req *http.Request) {

	code := req.FormValue("code")

	if len(code) != 0 {
		tok, err := igConf.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Println(err)
			http.NotFound(res, req)
			return
		}

		if tok.Valid() {
			client := igConf.Client(oauth2.NoContext, tok)

			request, err := http.NewRequest("GET", "https://api.instagram.com/v1/users/self/?access_token="+tok.AccessToken, nil)
			if err != nil {
				fmt.Println(err)
				http.NotFound(res, req)
				return
			}

			resp, err := client.Do(request)
			if err != nil {
				fmt.Println(err)
				http.NotFound(res, req)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			err = json.Unmarshal(body, &profileResponse)
			if err != nil {
				fmt.Println(err)
				http.NotFound(res, req)
				return
			}
			//res.Write(body)
			http.Redirect(res, req, "/profile", 301)

		}

		http.NotFound(res, req)
	}

}

func homePage(res http.ResponseWriter, req *http.Request) {
	url := igConf.AuthCodeURL("", oauth2.AccessTypeOffline)
	fmt.Println(url)
	err := templ.Execute(res, url)
	if err != nil {
		fmt.Println(err)
	}
}

func profilePage(res http.ResponseWriter, req *http.Request) {
	err := profile.Execute(res, profileResponse)
	if err != nil {
		fmt.Println(err)
	}
}
