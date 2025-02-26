package verify

type SendEmailResponse struct {
	Hash string `json:"hash"`
}

type CheckHashResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type SendEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
