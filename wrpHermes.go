package creekmail

import (
	"fmt"
	"log"

	"github.com/matcornic/hermes/v2"
)

// mailGenerator is a exposed hermes object for transparent interaction with
// body generator function(s)
var mailGenerator hermes.Hermes
var mailGeneratorWasUserDefined bool

// mainColor is the default color used for buttons
var mainColor string

// Setup returns a hermes object who is needed for mail generation
func Setup(coName, coWebSite, coLogo string,
	coFoundationYear uint, color string) {

	mainColor = validatedHexColor(color, errColor)

	copyright := fmt.Sprintf("Copyright © %d %s. All rights reserved.",
		coFoundationYear, coName)

	mailGenerator = hermes.Hermes{
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			Name:      coName,
			Link:      coWebSite,
			Logo:      coLogo,
			Copyright: copyright,
		},
	}

	mailGeneratorWasUserDefined = true
}

// EmailBody is a struct that wrap an email body in all its possible formats
type EmailBody struct {
	TEXT string
	HTML string
}

// GenEmailBody returns a EmailBody object with its fields filled
func GenEmailBody(targetName string, msg, outro []string,
	buttons ...hermes.Action) EmailBody {

	if !mailGeneratorWasUserDefined {
		log.Fatalln("'GenEmailBody' shouldn't be called without 'Setup' the mail generator")
		return EmailBody{}
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name:    targetName,
			Intros:  msg,
			Actions: buttons,
			Outros:  outro,
		},
	}

	text, err := mailGenerator.GeneratePlainText(email)
	if err != nil {
		log.Println("Err in PlainTEXT generation, blank is returned")
	}

	html, err := mailGenerator.GenerateHTML(email)
	if err != nil {
		log.Println("Err in HMTL generation, blank is returned")
	}

	return EmailBody{
		TEXT: text,
		HTML: html,
	}
}

// Button returns a hermes action that implies a call to action text
// with a button with customizable color
func Button(msg, text, link, color string) hermes.Action {
	return hermes.Action{
		Instructions: msg,
		Button: hermes.Button{
			Color: validatedHexColor(color, mainColor),
			Text:  text,
			Link:  link,
		},
	}
}
