package main

type Offer struct {
	Azienda     string `json:"azienda"`
	Email       string `json:"email"`
	Ruolo       string `json:"ruolo"`
	Descrizione string `json:"descrizione"`
	Competenze  string `json:"competenze"`
	Benefits    string `json:"benefits"`
	FullTime    bool   `json:"fulltime"`
	PartTime    bool   `json:"parttime"`
}
