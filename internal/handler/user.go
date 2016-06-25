package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	a "elo-api/internal/config"

	"golang.org/x/oauth2"
)

// Login is a handler to create account or login via Oauth
func Login(w http.ResponseWriter, r *http.Request) {
	conf := &oauth2.Config{
		ClientID:     a.Config.ClientID,
		ClientSecret: a.Config.ClientSecret,
		Scopes:       a.Config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  a.Config.EndpointAuthURL,
			TokenURL: a.Config.EndpointTokenURL,
		},
	}

	var code string
	code = r.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// Get the email
	// Check if user exists, if not create entry in db
	// return email+token
	chocoCookie := &http.Cookie{Name: "Token", Value: tok.AccessToken, HttpOnly: false, Expires: tok.Expiry}
	http.SetCookie(w, chocoCookie)
	fmt.Fprintf(w, string(body))

}

// GetUser is a handler to create a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	//player := model.GetPlayerFromId(r.URL.Get("player_id"))
	//json.NewEncoder(w).Encode(player)
}

// GetTopPlayers is a handler to create a user
func GetTopPlayers(w http.ResponseWriter, r *http.Request) {

}
