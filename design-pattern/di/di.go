package di

//depencies injection example
import "fmt"

type MessageService interface {
	SendEmail(message string, receiver string)
}

type Consumer interface {
	ProcessMessage(message string, receiver string)
}

type MessageServiceInjector interface {
	getConsumer() Consumer
}

//start service
type EmailService struct{}

func (e EmailService) SendEmail(message string, receiver string) {
	fmt.Printf("Sending message to email %v. message: %v. \n", receiver, message)
}

type SmsService struct{}

func (s SmsService) SendEmail(message string, receiver string) {
	fmt.Printf("Sending SMS to number %v. message: %v. \n", receiver, message)
}

type EmailServiceInjector struct{}

func (serviceInjector EmailServiceInjector) getConsumer() Consumer {
	return NewDIApp(EmailService{})
}

type SMSServiceInjector struct{}

func (serviceInjector SMSServiceInjector) getConsumer() Consumer {
	return NewDIApp(SmsService{})
}

//end service

//start main code
type MyDIApplication struct {
	service MessageService
}

func NewDIApp(service MessageService) MyDIApplication {
	return MyDIApplication{service}
}

func (a MyDIApplication) ProcessMessage(message string, receiver string) {
	a.service.SendEmail(message, receiver)
}

//end main code

//start client app
// func main() {
// 	message := "this is text message"
// 	email := "ruly.anggriawan@live.com"
// 	phoneNumber := "007"
// 	var injector MessageServiceInjector
// 	var app Consumer

// 	injector = EmailServiceInjector{}
// 	app = injector.getConsumer()
// 	app.ProcessMessage(message, email)

// 	injector = SMSServiceInjector{}
// 	app = injector.getConsumer()
// 	app.ProcessMessage(message, phoneNumber)
// }

//end client app
