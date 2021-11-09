package twilio

import (
	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	"github.com/twilio/twilio-go"
)

var client *twilio.RestClient

func GetTwilioClient() *twilio.RestClient {
	return client
}
func init() {
	params := twilio.RestClientParams{
		Username: config.Config["TWILIO_ACC_SID"],
		Password: config.Config["TWILIO_AUTH_TOKEN"],
	}
	client = twilio.NewRestClientWithParams(params)
}
