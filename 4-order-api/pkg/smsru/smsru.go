package smsru

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Send(apiId, phone, message string) (string, error) {
	fmt.Println(phone, message)
	phoneNums := strings.Replace(phone, "+", "", -1)
	msgForApi := strings.Replace(message, " ", "+", -1)
	smsUrl := fmt.Sprintf("https://sms.ru/sms/send?api_id=%s&to=%s&msg=%s&json=1", apiId, phoneNums, msgForApi)
	resp, err := http.Get(smsUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении тела ответа: %v", err)
		return "", err
	}

	var smsResponse SmsResponse
	if err := json.Unmarshal(body, &smsResponse); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err)
		return "", err
	}
	if smsResponse.Status == "ERROR" {
		return "", errors.New(smsResponse.StatusText)
	}
	return smsResponse.Status, nil
}

func Call(apiId, phone string) (string, error) {
	phoneNums := strings.Replace(phone, "+", "", -1)
	callUrl := fmt.Sprintf("https://sms.ru/code/call?phone=%s&ip=33.22.11.55&api_id=%s", phoneNums, apiId)
	resp, err := http.Get(callUrl)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении тела ответа: %v", err)
		return "", err
	}
	var callResponse CallResponse
	if err := json.Unmarshal(body, &callResponse); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err)
		return "", err
	}
	if callResponse.Status == "ERROR" {
		return "", errors.New(callResponse.StatusText)
	}
	return string(rune(callResponse.Code)), nil
}
