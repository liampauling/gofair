package gofair

import (
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
	"strings"
)

func createUrl(endpoint string, method string)(string){
	return endpoint + method
}


func (b *Betting) Request(url string, params *Params, v interface{}) (error) {
	//params.Locale = b.Client.config.Locale

	bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	body := strings.NewReader(string(bytes))

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// set headers
	req.Header.Set("X-Application", b.Client.config.AppKey)
	req.Header.Set("X-Authentication", b.Client.session.SessionToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection","keep-alive")

	client := &http.Client {}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}
