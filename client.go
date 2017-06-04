package gofair

import (
	"fmt"
	"net/http"
	"time"
	"crypto/tls"
	"net"
	"strings"
	"io/ioutil"
	"errors"
)


// betfair api endpoints
const (
	identity_url = "https://identitysso-api.betfair.com/api/"
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


// main client object
type Client struct {
	config		*Config
	httpClient	*http.Client
	session		*session
	loginTime	time.Time

	Auth		*Auth
}


func (a *Client) PrintSession() {
	fmt.Println(a.session)
	//fmt.Println(a.Auth.session)
}


// holds session data
type session struct {
	SessionToken string
	LoginStatus  string
}


// creates new client
func NewClient(config *Config)(*Client, error){
	c := new(Client)

	// create session
	c.session = new(session)

	// set config
	c.config = config

	// HTTP client
	cert, err := tls.LoadX509KeyPair(c.config.CertFile, c.config.KeyFile)
	if err != nil {
		return c, err
	}
	ssl := &tls.Config {
		Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	c.httpClient = &http.Client {
		Transport: &http.Transport {
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(time.Second*3))
			},
			TLSClientConfig: ssl,
		},
	}

	c.Auth = &Auth{config, c.session}

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


func loginRequest(endpoint string, method string, body *strings.Reader) ([]byte, error){
	// build url
	url := createUrl(endpoint, method)

	// load certs
	cert, err := tls.LoadX509KeyPair("/gocerts/client-2048.crt", "/gocerts/client-2048.key")
	if err != nil {
		fmt.Println("Error loading certificate. ", err)
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
