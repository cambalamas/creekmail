package creekmail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var err error
var conn *email.Pool

type CreekMail {
	//TODO: Define how to send mail...
	email.Email
}

// Conn defines a pool of connections accessible from the entire package
// return if it was succesful or not
func Conn(from, pass, host string, port uint) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("", from, pass, host)

	conn, err = email.NewPool(address, 4, auth)
	return err == nil
}

func Send()