package gofair

import (
	"strings"
	"encoding/json"
	"time"
	"crypto/tls"
	"net/http"
	"errors"
	"io/ioutil"
)


type loginResult struct {
	LoginStatus	string	`json:"loginStatus"`
	SessionToken	string 	`json:"sessionToken"`
}


func (c *Client) Login() (loginResult, error) {
	// build body
	body := strings.NewReader("username=" + c.config.Username + "&password=" + c.config.Password)

	// make request
	resp, err := loginRequest(c, login_url, "certlogin",  body)
	if err != nil {
		return *new(loginResult), err
	}

	var result loginResult

	// parse json
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	c.session.SessionToken = result.SessionToken
	c.session.LoginTime = time.Now().UTC()
	return result, nil
}


func loginRequest(c *Client, endpoint string, method string, body *strings.Reader) ([]byte, error){
	// build url
	url := createUrl(endpoint, method)

	// HTTP client
	ssl := &tls.Config {
		Certificates: []tls.Certificate{*c.certificates},
		InsecureSkipVerify: true,
	}

	client := &http.Client {
		Transport: &http.Transport {
			TLSClientConfig: ssl,
		},
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// set headers
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
