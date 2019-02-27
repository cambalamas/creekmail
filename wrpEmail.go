package creekmail

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

var err error
var conn *email.Pool

var coMail string

// Conn defines a pool of connections accessible from the entire package
// return if it was succesful or not
func Conn(from, pass, host string, port uint) bool {
	coMail = from

	address := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("", from, pass, host)

	conn, err = email.NewPool(address, 4, auth)
	return err == nil
}

// Send wrap the process needed to send a mail, keeping in mind that
// to or cc should be filled if you wanna fill bcc filed
func Send(subject, plain, html string, to, cc, bcc []string) error {

	to = validStrSlice(to)
	cc = validStrSlice(cc)
	bcc = validStrSlice(bcc)

	e := email.NewEmail()
	e.From = coMail
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.HTML = []byte(html)
	e.Text = []byte(plain)

	return conn.Send(e, 10*time.Second)
}
