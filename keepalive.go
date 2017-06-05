package gofair

import (
	"encoding/json"
	"time"
)


type keepAliveResult struct {
	SessionToken	string `json:"sessionToken"`
	Token		string `json:"token"`
	Status		string `json:"status"`
	Error		string `json:"error"`
}


func (c *Client) KeepAlive() (keepAliveResult, error) {
	// make request
	resp, err := logoutRequest(c, identity_url, "keepAlive")
	if err != nil {
		return *new(keepAliveResult), err
	}

	var result keepAliveResult

	// parse json
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	c.session.SessionToken = result.Token
	c.session.LoginTime = time.Now().UTC()
	return result, nil
}
