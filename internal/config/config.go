package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is the global application config, set during init
var Config Conf

type Conf struct {
	// Port
	Port int    `toml:"port"`
	Host string `toml:"host"`

	// ClientID is the application's ID.
	ClientID string `toml:"client_id"`

	// ClientSecret is the application's secret.
	ClientSecret string `toml:"client_secret"`

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string `toml:"redirect_url"`

	EndpointAuthURL  string `toml:"endpoint_auth_url"`
	EndpointTokenURL string `toml:"endpoint_token_url"`

	// Scope specifies optional requested permissions.
	Scopes []string `toml:"scopes"`

	// Mysql is the dsn used
	Mysql string `toml:"mysql_dsn"`
}

// InitConfig reads info from config file
func InitConfig() {
	var configfile = "./etc/app.ini"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Conf
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	Config = config
}
