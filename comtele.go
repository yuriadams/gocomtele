// Package gocomtele is a library for interacting with http://sms.comtele.com.br/ API.
package comtele

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const baseURL string = "https://sms.comtele.com.br/api/"

// Comtele stores basic information important for connecting to the
// sms.comtele.com.br REST api such as AuthToken.
type Comtele struct {
	AuthToken string
}

// Create a new Comtele struct.
func NewComteleClient(authToken string) *Comtele {
	return &Comtele{authToken}
}

// SendTextMessage uses Comtele to send a text message.
func (comtele *Comtele) SendSMS(sender string, receivers string, content string) string {
	apiURL := baseURL + comtele.AuthToken + "/sendmessage?sender=" + sender + "&receivers=" + receivers + "&content=" + url.QueryEscape(content)
	form := url.Values{}

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(form.Encode())))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	messageResp := ""

	json.Unmarshal(responseBody, messageResp)

	return messageResp
}
