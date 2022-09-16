package model

type Result struct {
	Code      int       `json:"code"`
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Clipboard Clipboard `json:"clipboard"`
}
