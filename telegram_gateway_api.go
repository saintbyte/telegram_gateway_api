package telegram_gateway_api

type TelegramGatewayAPI struct{}

func NewTelegramGatewayAPI() *TelegramGatewayAPI {
	return &TelegramGatewayAPI{}
}

func (t *TelegramGatewayAPI) getUrl() string {
	return ""
}
func (t *TelegramGatewayAPI) getApiToken() string {
	return ""
}
func (t *TelegramGatewayAPI) getRequest() {

}

func (t *TelegramGatewayAPI) SendVerificationMessage() {

}
func (t *TelegramGatewayAPI) CheckSendAbility() {

}
func (t *TelegramGatewayAPI) CheckVerificationStatus() {

}
func (t *TelegramGatewayAPI) RevokeVerificationMessage() {

}
func (t *TelegramGatewayAPI) CheckCallbackRequest() bool {
	return false
}
