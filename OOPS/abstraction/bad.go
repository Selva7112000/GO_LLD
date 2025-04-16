package abstraction

import "fmt"

type SendEmail struct{}

func (s *SendEmail) Connect() {
	fmt.Println("connecting to gmail server")
}

func (s *SendEmail) Authenticate() {
	fmt.Println("Authenticating...")
}

func (s *SendEmail) SendGmail() {
	fmt.Println("Email has been sent")
}

func (s *SendEmail) Disconnect() {
	fmt.Println("disconnecting from the gmail server")
}

func BadSendGmail() {
	sendGmail := SendEmail{}
	sendGmail.Connect()
	sendGmail.Authenticate()
	sendGmail.SendGmail()
	sendGmail.Disconnect()
}
