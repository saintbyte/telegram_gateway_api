package telegram_gateway_api

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

type SendVerificationMessageRequest struct {
	PhoneNumber    string `json:"phone_number"`              // Yes	The phone number to which you want to send a verification message, in the E.164 format.
	RequestId      string `json:"request_id,omitempty"`      // Optional See SetAutoRequestId
	SenderUsername string `json:"sender_username,omitempty"` // Optional See SetSenderUsername
	Code           string `json:"code,omitempty"`            // Optional See SetCode
	CodeLength     int    `json:"code_length,omitempty"`     //	Optional See SetCodeLength
	CallbackUrl    string `json:"callback_url,omitempty"`    // Optional See SetCallbackUrl
	Payload        string `json:"payload,omitempty"`         // Optional See SetPayload
	TTL            int    `json:"ttl,omitempty"`             // Optional See SetTTL
}

func NewSendVerificationMessageRequest(phone string) *SendVerificationMessageRequest {
	// Создает SendVerificationMessageRequest
	return &SendVerificationMessageRequest{
		PhoneNumber: phone,
	}
}
func (sv *SendVerificationMessageRequest) SetAutoRequestId() string {
	// Автоматически задает Request id и возращает его.
	//
	// Из документации телеграмм: The unique identifier of a previous request from checkSendAbility.
	// If provided, this request will be free of charge.
	uuidString := uuid.NewString()
	sv.RequestId = uuidString
	return uuidString
}
func (sv *SendVerificationMessageRequest) SetTTL(ttl int) error {
	// высталяет ttl с валидацией.
	//
	// Из документации телеграмм: Time-to-live (in seconds)
	// before the message expires and is deleted.
	// The message will not be deleted if it has already been read.
	//If not specified, the message will not be deleted. Supported values are from 60 to 86400.
	if ttl < 60 {
		return errors.New("ttl out of range. so small")
	}
	if ttl > 86400 {
		return errors.New("ttl out of range. to big")
	}
	sv.TTL = ttl
	return nil
}
func (sv *SendVerificationMessageRequest) SetPayload(payload string) error {
	// Выставляет дополнительную нагрузку.
	//
	// Из докментации телеграмм: Custom payload, 0-128 bytes. This will not be displayed to the user,
	// use it for your internal processes.
	if payload == "" {
		return errors.New("payload is empty")
	}
	if len(payload) > 128 {
		return errors.New("payload is too long")
	}
	sv.Payload = payload
	return nil
}
func (sv *SendVerificationMessageRequest) SetCallbackUrl(url string) error {
	// Куда возращать вебхук.
	//
	// Из докментации телеграмм: CAn HTTPS URL where you want to receive delivery reports
	//related to the sent message, 0-256 bytes.
	if url == "" {
		return errors.New("url is empty")
	}
	if len(url) > 256 {
		return errors.New("url is too long")
	}
	if strings.ToLower(url[0:7]) != "https://" {
		return errors.New("url should be https://")
	}
	sv.CallbackUrl = url
	return nil
}
func (sv *SendVerificationMessageRequest) SetSenderUsername(username string) error {
	// Выставляем Имя пользоватя или канал от имени которого придет код.
	//
	//Из документации телеграмм: Username of the Telegram channel from which the code will be sent. The specified channel,
	//if any, must be verified and owned by the same account who owns the Gateway API token.
	if username[0:1] == "@" {
		username = username[0:]
	}
	sv.SenderUsername = username
	return nil
}

func (sv *SendVerificationMessageRequest) SetCode(code string) error {
	// Поставить свой код.
	//
	// Из документации телеграмм: The verification code. Use this parameter
	// if you want to set the verification code yourself.
	// Only fully numeric strings between 4 and 8 characters in length are supported.
	// If this parameter is set, code_length is ignored.
	if code == "" {
		return errors.New("code is empty")
	}
	if len(code) < 4 {
		return errors.New("code is too small")
	}
	if len(code) > 8 {
		return errors.New("code is too long")
	}
	sv.Code = code
	return nil
}

func (sv *SendVerificationMessageRequest) SetCodeLength(codeLength int) error {
	// Выставляет длину кода который надо сгенери
	//
	//  Из документации телеграмм: The length of the verification code if Telegram needs to generate it for you.
	// Supported values are from 4 to 8. This is only relevant if you are not using the code parameter
	// to set your own code. Use the checkVerificationStatus method with the code
	// parameter to verify the code entered by the user.
	if codeLength < 4 {
		return errors.New("code is too small")
	}
	if codeLength > 8 {
		return errors.New("code is too big")
	}
	sv.CodeLength = codeLength
	return nil
}
