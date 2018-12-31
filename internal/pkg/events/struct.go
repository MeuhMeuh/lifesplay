package events

// Defines the credentials structure used for the Calendar OAuth2 authentication.
type credentials struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURI  string `mapstructure:"redirect_uri"`
	AuthURI      string `mapstructure:"auth_uri"`
	TokenURI     string `mapstructure:"token_uri"`
}

type token struct {
	AccessToken  string `mapstructure:"access_token"`
	TokenType    string `mapstructure:"token_type"`
	RefreshToken string `mapstructure:"refresh_token"`
	Expiry       string `mapstructure:"expiry"`
}
