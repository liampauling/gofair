package gofair

import (
	"encoding/json"
	"time"
	"net/http"
	"errors"
	"io/ioutil"
)


type logoutResult struct {
	Token	string `json:"token"`
	Product	string `json:"product"`
	Status	string `json:"status"`
	Error	string `json:"error"`
}


func (c *Client) Logout() (logoutResult, error) {
	// build url
	url := createUrl(identity_url, "logout")

	// make request
	resp, err := logoutRequest(c, url)
	if err != nil {
		return *new(logoutResult), err
	}

	var result logoutResult

	// parse json
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	c.session.SessionToken = ""
	c.session.LoginTime = time.Time{}
	return result, nil
}


func logoutRequest(c *Client, url string) ([]byte, error){

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	// set headers
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
