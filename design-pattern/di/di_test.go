package di

import "testing"

func TestProcessMessage(t *testing.T) {

	message := "this is text message"
	email := "ruly.anggriawan@live.com"
	phoneNumber := "007"
	var injector MessageServiceInjector
	var app Consumer

	injector = EmailServiceInjector{}
	app = injector.getConsumer()
	app.ProcessMessage(message, email)

	injector = SMSServiceInjector{}
	app = injector.getConsumer()
	app.ProcessMessage(message, phoneNumber)

	// t.Fatalf("process message fail")
}
