package auth

type SendSmsRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
	// e164 This validates that a string value contains a valid E.164 Phone number https://en.wikipedia.org/wiki/E.164 (ex. +1123456789)
}

type SendSmsResponse struct {
	Message string `json:"message"`
	Session string `json:"session"`
	Code    string `json:"code"` // Для тестовых целей
}

type LoginRequest struct {
	Session string `json:"session" validate:"required"`
	Code    string `json:"code" validate:"required,number"`
}

type LoginResponse struct {
	Message string `json:"message" enum:"success,error" validate:"required"`
	Token   string `json:"token"`
}
