package creekmail

import "regexp"

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
