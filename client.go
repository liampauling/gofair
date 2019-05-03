package gofair

import (
	"crypto/tls"
	"time"
	"strings"
)

// betfair api endpoints
const (
	login_url       = "https://identitysso-api.betfair.com/api/"
	identity_url    = "https://identitysso.betfair.com/api/"
	api_betting_url = "https://api.betfair.com/exchange/betting/rest/v1.0/"
	api_account_url = "https://api.betfair.com/exchange/account/rest/v1.0/"
	navigation_url  = "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json"
)

// holds login data
type Config struct {
	Username string
	Password string
	AppKey   string
	CertFile string
	KeyFile  string
	Locale   string
}

// holds session data
type session struct {
	SessionToken string
	LoginTime    time.Time
}

// main client object
type Client struct {
	config       *Config
	session      *session
	certificates *tls.Certificate
	Betting      *Betting
	Account      *Account
	Streaming    *Streaming
	Historical   *Historical
}

// betting object
type Betting struct {
	Client *Client
}

// account object
type Account struct {
	Client *Client
}

// streaming object
type Streaming struct {
	Client *Client
}

// historical object
type Historical struct {
	Client *Client
}

// creates new client
func NewClient(config *Config) (*Client, error) {
	c := new(Client)

	// create session
	c.session = new(session)
	var cert tls.Certificate
	var err error
	// create certificates
	// ----- is obviously not a path, therefore load direct from the variables
	if strings.HasPrefix(config.CertFile, "------") {
		cert, err = tls.X509KeyPair([]byte(config.CertFile), []byte(config.KeyFile))
		if err != nil {
			return nil, err
		}
	} else {
		cert, err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		if err != nil {
			return nil, err
		}
	}
	c.certificates = &cert

	// set config
	c.config = config

	// create betting
	c.Betting = new(Betting)
	c.Betting.Client = c

	// create account
	c.Account = new(Account)
	c.Account.Client = c

	// create streaming
	c.Streaming = new(Streaming)
	c.Streaming.Client = c

	// create historical
	c.Historical = new(Historical)
	c.Historical.Client = c

	return c, nil
}

func (c *Client) SessionExpired() bool {
	// returns True if client not logged in or expired
	// betfair requires keep alive every 4hrs (20mins ITA)
	if c.session.SessionToken == "" {
		return true
	}
	duration := time.Since(c.session.LoginTime)
	if duration.Minutes() > 200 {
		return true
	} else {
		return false
	}
}
