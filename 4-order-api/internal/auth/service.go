package auth

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/sessions"
	"go-ps-adv-homework/internal/smsru"
	"go-ps-adv-homework/internal/user"
	"go-ps-adv-homework/pkg/jwt"
	"net/http"
)

type AuthService struct {
	*configs.Config
	*user.UserRepository
	*sessions.SessionRepository
	SmsService *smsru.SmsruService
}

type AuthServiceDependencies struct {
	*configs.Config
	*user.UserRepository
	*sessions.SessionRepository
	SmsService *smsru.SmsruService
}

type SendCodeResponse struct {
	Session *sessions.Session `json:"session"`
}

func NewAuthService(dependencies AuthServiceDependencies) *AuthService {
	return &AuthService{
		Config:            dependencies.Config,
		UserRepository:    dependencies.UserRepository,
		SessionRepository: dependencies.SessionRepository,
		SmsService:        dependencies.SmsService,
	}
}

func (service *AuthService) SendCode(phone string, callToPhone bool) (*sessions.Session, error) {
	// Check user and create user
	foundedUser, err := service.UserRepository.GetOrCreateUser(phone)
	if err != nil {
		return nil, err
	}
	// Generate code and session
	code := ""
	// Call, no sms
	if callToPhone {
		// Get code from service
		code, err = service.SmsService.CallToPhone(phone)
		if err != nil {
			return nil, err
		}
	}
	newSession, err := sessions.NewSession(foundedUser.ID, foundedUser.Phone, code)
	if err != nil {
		return nil, err
	}
	session, err := service.SessionRepository.CreateOrUpdateSession(newSession)
	if err != nil {
		return nil, err
	}
	// Sms, no call
	if !callToPhone {
		// Send sms
		_, err = service.SmsService.SendSms(session.Phone, session.Code)
		if err != nil {
			return nil, err
		}
	}
	return session, nil
}

func (service *AuthService) CheckCode(sessionStr, code string) (string, error, int) {
	// Check session and code
	session, err := service.SessionRepository.GetSession(sessionStr)
	if err != nil {
		return "", err, http.StatusUnauthorized
	}
	if session.Code != code {
		return "", err, http.StatusUnauthorized
	}
	// JWT Token
	token, err := jwt.NewJWT(service.Config.Auth.Secret).SignToken(jwt.JWTData{
		Phone:   session.Phone,
		Session: session.Session,
	})
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	return token, nil, http.StatusCreated
}
