package smsru

type SmsResponse struct {
	Status     string                      `json:"status"`
	StatusCode int                         `json:"status_code"`
	StatusText string                      `json:"status_text"`
	Sms        map[string]SmsPhoneResponse `json:"sms"`
	Balance    float64                     `json:"balance"`
}

type SmsPhoneResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	SmsId      string `json:"sms_id,omitempty"`
	StatusText string `json:"status_text,omitempty"`
}

type CallResponse struct {
	Status     string  `json:"status"`
	StatusText string  `json:"status_text,omitempty"`
	Code       int     `json:"code"`
	CallId     string  `json:"call_id,omitempty"`
	Cost       float64 `json:"cost,omitempty"`
	Balance    float64 `json:"balance,omitempty"`
}
