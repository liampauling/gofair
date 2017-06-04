package gofair

import (
	"net/http"
	"time"
	"crypto/tls"
	"net"
	"strings"
	"io/ioutil"
	"errors"
	"encoding/json"
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
}


type loginResult struct {
	LoginStatus	string	`json:"loginStatus"`
	SessionToken	string 	`json:"sessionToken"`
}


type logoutResult struct {
	Token	string `json:"token"`
	Product	string `json:"product"`
	Status	string `json:"status"`
	Error	string `json:"error"`
}


func (c *Client) Login() (loginResult) {
	// build body
	body := strings.NewReader("username=" + c.config.Username + "&password=" + c.config.Password)

	// make request
	resp, err := loginRequest(c, login_url, "certlogin",  body)
	check(err)

	var result loginResult

	// parse json
	err = json.Unmarshal(resp, &result)
	check(err)

	c.session.SessionToken = result.SessionToken
	c.session.LoginTime = time.Now().UTC()
	return result
}


func (c *Client) KeepAlive() (logoutResult) {
	// make request
	resp, err := logoutRequest(c, identity_url, "keepAlive")
	check(err)

	var result logoutResult

	// parse json
	err = json.Unmarshal(resp, &result)
	check(err)

	c.session.SessionToken = result.Token
	c.session.LoginTime = time.Now().UTC()
	return result
}


func (c *Client) Logout() (logoutResult) {
	// make request
	resp, err := logoutRequest(c, identity_url, "logout")
	check(err)

	var result logoutResult

	// parse json
	err = json.Unmarshal(resp, &result)
	check(err)

	c.session.SessionToken = ""
	c.session.LoginTime = time.Time{}
	return result
}


// creates new client
func NewClient(config *Config)(*Client, error){
	c := new(Client)

	// create session
	c.session = new(session)

	// set config
	c.config = config

	return c, nil
}


func createUrl(endpoint string, method string)(string){
	return endpoint + method
}


func check(e error){
	if e != nil {
		panic(e)
	}
}


func loginRequest(c *Client, endpoint string, method string, body *strings.Reader) ([]byte, error){
	// build url
	url := createUrl(endpoint, method)

	// HTTP client
	cert, err := tls.LoadX509KeyPair(c.config.CertFile, c.config.KeyFile)
	if err != nil {
		return nil, err
	}
	ssl := &tls.Config {
		Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	client := &http.Client {
		Transport: &http.Transport {
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(time.Second*3))
			},
			TLSClientConfig: ssl,
		},
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Application", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}


func logoutRequest(c *Client, endpoint string, method string) ([]byte, error){
	// build url
	url := createUrl(endpoint, method)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept","application/json")
	req.Header.Set("X-Application", c.config.AppKey)
	req.Header.Set("X-Authentication", c.session.SessionToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client {}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
