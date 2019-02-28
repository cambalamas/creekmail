package creekmail

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseYAML(t *testing.T) {
	// Test conn yml
	{
		yml := "testyml/Conn.yml"
		data := ConnDataFromYAML(yml)
		if spew.Sdump(data) == spew.Sdump(ConnData{}) {
			t.Errorf("Fail loading %s, struct is empty after load.", yml)
		}
	}

	// Test brand yml
	{
		yml := "testyml/Brand.yml"
		data := BrandDataFromYAML(yml)
		if spew.Sdump(data) != spew.Sdump(BrandData{}) {
			t.Errorf("Fail loading %s, struct is empty after load.", yml)
		}
	}

	// Test email yml
	{
		yml := "testyml/Email.yml"
		data := EmailDataFromYAML(yml)
		if spew.Sdump(data) == spew.Sdump(EmailData{}) {
			t.Errorf("Fail loading %s, struct is empty after load.", yml)
		}
	}

}
