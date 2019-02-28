package creekmail

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

const (
	errColor = "#FF00FF"
)

// validatedHexColor takes a color string and verify that match with
// hexadecimal color pattern using regexp. If the color match but doesn't
// have the hashtag (#) add it before the color string.
func validatedHexColor(color, onFail string) string {

	match, _ := regexp.MatchString("^#[0-9A-Fa-f]{3,6}$", color)
	if match {
		return color
	}

	matchNoHashtag, _ := regexp.MatchString("^[0-9A-Fa-f]{3,6}$", color)
	if matchNoHashtag {
		return "#" + color
	}

	return onFail
}

// validStrSlice returns a empty string slice if the arg slice len is 0 or
// all its entries are empty
func validStrSlice(ss []string) []string {
	empty := 0
	for _, s := range ss {
		if s == "" {
			empty++
		}
	}
	if empty > 0 {
		return []string{}
	}
	return ss
}

// FillFromYAML gived a yaml path, fills a gived object with its data
func FillFromYAML(yamlPath string, obj interface{}) {
	rawYAML, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Printf("rawYAML.Get err   #%v ", err)
	}
	fmt.Println(string(rawYAML))
	err = yaml.Unmarshal(rawYAML, &obj)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

// --- YML Helpers ---

func ymlLoad(ymlPath string) []byte {
	yml, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		log.Printf("yml.Get err   #%v ", err)
	}
	return yml
}

func ymlErr(err error, callback ...func()) {
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	if len(callback) > 0 {
		callback[0]()
	}
}
