package creekmail

type ConnData struct {
	Mail   string
	AppPwd string
	Host   string
	Port   int
}

type CompanyData struct {
	Name  string
	Link  string
	Logo  string
	Year  int
	Color string
}

type EmailActionData struct {
	Text        string
	ButtonText  string
	ButtonLink  string
	ButtonColor string
}

type EmailData struct {
	Subject  string
	Target   string
	IntroMsg []string
	Actions  []emailActionData
	OutroMsg []string
}
