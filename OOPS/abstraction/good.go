package abstraction

import (
	"fmt"
)

type sendMail struct{}

func (s *sendMail) connect() {
	fmt.Println("connecting to the gmail server")
}

func (s *sendMail) authenticate() {
	fmt.Println("authenticating...")
}

func (s *sendMail) disconnect() {
	fmt.Println("disconnected from the gmail server")
}

func (s *sendMail) SendGmail() {
	s.connect()
	s.authenticate()
	fmt.Println("sending the gmail")
	s.disconnect()
}

func GoodSendGmail() {
	sendGmail := sendMail{}
	sendGmail.SendGmail()
}
