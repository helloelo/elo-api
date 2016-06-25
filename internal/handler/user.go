package handler

import (
	"fmt"
	"log"
	"net/http"

	a "elo-api/internal/config"

	"golang.org/x/oauth2"
)

func Login2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login success")
}

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
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext, tok)
	profile, err := client.Get("profile")
	if err != nil {
		log.Fatal(err)
	}

	// Get the email
	// Check if user exists, if not create entry in db
	// return email+token
	fmt.Println("DEBUG: ", profile.Body.Read)
	fmt.Fprintf(w, "Login success")

}

// GetUser is a handler to create a user
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// GetTopPlayers is a handler to create a user
func GetTopPlayers(w http.ResponseWriter, r *http.Request) {

}
