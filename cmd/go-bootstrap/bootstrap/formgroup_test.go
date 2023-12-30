package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestFormGroup(t *testing.T) {
	// FormGroup(contents []htmlwrapper.Elm)
	html, err := FormGroup([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"form-group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormGroupContent(t *testing.T) {
	// FormGroup(contents []htmlwrapper.Elm)
	html, err := FormGroup([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"form-group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
