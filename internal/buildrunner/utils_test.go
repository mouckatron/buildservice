package buildrunner

import "testing"

func TestGetNameFromURL(t *testing.T) {
	URL := "https://github.com/mouckatron/buildservice.git"
	expected := "buildservice"

	subject := getNameFromURL(URL)

	if subject != expected {
		t.Errorf("Wanted %s, got %s", expected, subject)
	}
}
