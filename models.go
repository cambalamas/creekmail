package creekmail

type ConnData struct {
	Mail   string
	AppPwd string
	Host   string
	Port   uint
}

type CompanyData struct {
	Name  string
	Link  string
	Logo  string
	Year  uint
	Color string
}

type EmailActionData struct {
	Text        string
	ButtonText  string
	ButtonLink  string
	ButtonColor string
}

type EmailData struct {
	To       []string
	Cc       []string
	Bcc      []string
	Subject  string
	Target   string
	IntroMsg []string
	Actions  []EmailActionData
	OutroMsg []string
}
