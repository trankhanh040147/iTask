package googleprovider

import (
	"paradise-booking/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig(cfg *config.Config) *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		RedirectURL:  cfg.Google.RedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return config
}
