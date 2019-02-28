package creekmail

import yaml "gopkg.in/yaml.v2"

///////////////////////////////////////////////////////////////////////////////
// ConnData
///////////////////////////////////////////////////////////////////////////////

// ConnData wraps data needed for smtp requests
type ConnData struct {
	Mail   string `yaml:"mail"`
	AppPwd string `yaml:"appPwd"`
	Host   string `yaml:"host"`
	Port   uint   `yaml:"port"`
}

// ConnDataFromYAML fills a ConnData object from YAML file
func ConnDataFromYAML(ymlPath string) ConnData {
	var o ConnData
	ymlBytes := ymlLoad(ymlPath)
	ymlErr(yaml.Unmarshal(ymlBytes, &o))
	return o
}

///////////////////////////////////////////////////////////////////////////////
// BrandData
///////////////////////////////////////////////////////////////////////////////

// BrandData wraps data needed for email template to looks
// as your Brand is
type BrandData struct {
	Name  string `yaml:"name"`
	Link  string `yaml:"link"`
	Logo  string `yaml:"logo"`
	Year  uint   `yaml:"year"`
	Color string `yaml:"color"`
}

// BrandDataFromYAML fills a BrandData object from YAML file
func BrandDataFromYAML(ymlPath string) BrandData {
	var o BrandData
	ymlBytes := ymlLoad(ymlPath)
	ymlErr(yaml.Unmarshal(ymlBytes, &o))
	return o
}

///////////////////////////////////////////////////////////////////////////////
// EmailData
///////////////////////////////////////////////////////////////////////////////

// EmailData wraps data needed to send an email
type EmailData struct {
	To       []string `yaml:"to"`
	Cc       []string `yaml:"cc"`
	Bcc      []string `yaml:"bcc"`
	Subject  string   `yaml:"subject"`
	Target   string   `yaml:"target"`
	IntroMsg []string `yaml:"introMsg"`
	Actions  []struct {
		Msg   string `yaml:"Text"`
		Text  string `yaml:"ButtonText"`
		Link  string `yaml:"ButtonLink"`
		Color string `yaml:"ButtonColor"`
	} `yaml:"actions"`
	OutroMsg []string `yaml:"outroMsg"`
}

// EmailDataFromYAML fills a EmailData object from YAML file
func EmailDataFromYAML(ymlPath string) EmailData {
	var o EmailData
	ymlBytes := ymlLoad(ymlPath)
	ymlErr(yaml.Unmarshal(ymlBytes, &o))
	return o
}
