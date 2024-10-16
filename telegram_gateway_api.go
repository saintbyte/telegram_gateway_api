package telegram_gateway_api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type TelegramGatewayAPI struct {
	ApiToken string
	Debug    bool
}

func NewTelegramGatewayAPI() *TelegramGatewayAPI {
	return &TelegramGatewayAPI{
		ApiToken: "",
		Debug:    true,
	}
}

func (t *TelegramGatewayAPI) getUrl(method string) string {
	return "https://" + TelegramGatewayApiHost + "/" + method
}
func (t *TelegramGatewayAPI) getApiToken() string {
	if t.ApiToken != "" {
		return t.ApiToken
	}
	value, ok := os.LookupEnv("TELEGRAM_GATEWAY_API_TOKEN")
	if ok {
		t.ApiToken = value
		return value
	}
	return ""
}

func (t *TelegramGatewayAPI) getRequest(url string, body io.Reader) (*http.Request, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, err := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+t.getApiToken())
	if err != nil {

		return nil, err
	}
	return request, nil
}
func (t *TelegramGatewayAPI) makeRequest(request *http.Request) (string, error) {
	client := &http.Client{}
	response, e := client.Do(request)
	if e != nil {
		return "", e
	}
	if response.StatusCode != http.StatusOK {
		return "", errors.New("HTTP code:" + string(response.StatusCode))
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	log.Println(string(body))
	return string(body), nil
}
func (t *TelegramGatewayAPI) SendVerificationMessage(messageRequest SendVerificationMessageRequest) {
	url := t.getUrl("sendVerificationMessage")
	json_body, err := json.Marshal(messageRequest)
	request, err := t.getRequest(url, json_body)
	if err != nil {

	}
	response, err2 := t.makeRequest(request)
}
func (t *TelegramGatewayAPI) CheckSendAbility() {
	url := t.getUrl("checkSendAbility")
	request, err := t.getRequest(url)
}
func (t *TelegramGatewayAPI) CheckVerificationStatus() {
	url := t.getUrl("checkVerificationStatus")
	request, err := t.getRequest(url)
}
func (t *TelegramGatewayAPI) RevokeVerificationMessage() {
	url := t.getUrl("revokeVerificationMessage")
	request, err := t.getRequest(url)
}
func (t *TelegramGatewayAPI) CheckCallbackRequest() bool {
	return false
}
