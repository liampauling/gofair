package gofair

import (
	"time"
	"crypto/tls"
)


// betfair api endpoints
const (
	login_url = "https://identitysso-api.betfair.com/api/"
	identity_url = "https://identitysso.betfair.com/api/"
	api_url = "https://api.betfair.com/exchange/"
	navigation_url = "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json"
)


// holds login data
type Config struct {
	Username 	string
	Password 	string
	AppKey		string
	CertFile 	string
	KeyFile 	string
	Locale		string
}


// holds session data
type session struct {
	SessionToken 	string
	LoginTime	time.Time
}


// main client object
type Client struct {
	config		*Config
	session		*session
	certificates	*tls.Certificate
}


// creates new client
func NewClient(config *Config)(*Client, error){
	c := new(Client)

	// create session
	c.session = new(session)

	// create certificates
	cert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		return nil, err
	}
	c.certificates = &cert

	// set config
	c.config = config

	return c, nil
}
