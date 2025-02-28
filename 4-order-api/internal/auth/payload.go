package auth

type SendSmsRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
	// e164 This validates that a string value contains a valid E.164 Phone number https://en.wikipedia.org/wiki/E.164 (ex. +1123456789)
}

type SendSmsResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type LoginRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
	Code  string `json:"code" validate:"required,min=6,max=6,number"`
}

type LoginResponse struct {
	Message string `json:"message" enum:"success,error" validate:"required"`
	Token   string `json:"token"`
}
