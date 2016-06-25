package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is the global application config, set during init
var Config config

type config struct {
	// App Port
	Port int `config:"port"`

	// ClientID is the application's ID.
	ClientID string `config:"client_id"`

	// ClientSecret is the application's secret.
	ClientSecret string `config:"client_secret"`

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string `config:"redirect_url"`

	EndpointAuthURL  string `config:"endpoint_auth_url"`
	EndpointTokenURL string `config:"endpoint_token_url"`

	// Scope specifies optional requested permissions.
	Scopes []string `config:"scopes"`

	// Mysql is the dsn used
	Mysql string `config:"mysql_dsn"`
}

// InitConfig reads info from config file
func InitConfig() {
	var configfile = "/etc/app.ini"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	Config = config
}
