package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var ReCaptchaConf ReCaptchaConfig

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

func CheckRecaptcha(secret, response string) error {
	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	// Add necessary request parameters.
	q := req.URL.Query()
	q.Add("secret", secret)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response.
	var body SiteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	log.Println(body)
	// Check recaptcha verification success.
	if !body.Success {
		return errors.New("unsuccessful recaptcha verify request")
	}

	// Check response score.
	if body.Score < 0.5 {
		return errors.New("lower received score than expected")
	}

	// Check response action.
	if body.Action != "submit" {
		return errors.New("mismatched recaptcha action")
	}

	return nil
}
