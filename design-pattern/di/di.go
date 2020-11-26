package main

//depencies injection example
import "fmt"

//start service
type EmailService struct{}

func (e EmailService) SendEmail(message string, receiver string) {
	fmt.Printf("Sending Email : %v. to %v. \n", message, receiver)
}

//end service

//start main code
type MyApplication struct {
	email EmailService
}

func (a MyApplication) ProcessMessage(message string, receiver string) {
	a.email.SendEmail(message, receiver)
}

//end main code

//start client app
func main() {
	app := MyApplication{}
	app.ProcessMessage("hi ruly", "ruly.anggriawan@live.com")
}

//end client app
