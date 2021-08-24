package main

import "time"

type Offer struct {
	Azienda        string `json:"azienda"`
	Email          string `json:"email"`
	Ruolo          string `json:"ruolo"`
	Salario        string `json:"salario"`
	Descrizione    string `json:"descrizione"`
	Competenze     string `json:"competenze"`
	Benefits       string `json:"benefits"`
	FullTime       bool   `json:"fulltime"`
	PartTime       bool   `json:"parttime"`
	ReCaptchaToken string `json:"recaptcha_token"`
}

type ReCaptchaConfig struct {
	SiteKey     string
	Secret      string
	ApiEndpoint string
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type SiteVerifyRequest struct {
	RecaptchaResponse string `json:"g-recaptcha-response"`
}
