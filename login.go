package gofair

import (
	"strings"
	"encoding/json"
	"fmt"
)


type loginResult struct {
	LoginStatus	string	`json:"loginStatus"`
	SessionToken	string 	`json:"sessionToken"`
}


type Auth struct {
	config		*Config
	session		*session
}


func (a *Auth) Login() {
	// build body
	body := strings.NewReader("username=" + a.config.Username + "&password=" + a.config.Password)

	// make request
	resp, err := loginRequest(identity_url, "certlogin",  body)
	check(err)

	var result loginResult

	// parse json
	err = json.Unmarshal(resp, &result)
	check(err)

	a.session.LoginStatus = result.LoginStatus
	a.session.SessionToken = result.SessionToken
}


func (a Auth) PrintSession() {
	fmt.Println(a.session)
}
