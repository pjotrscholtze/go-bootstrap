package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestCol(t *testing.T) {
	// Col(contents []htmlwrapper.Elm, size BsColSize, formGroup bool)
	html, err := Col([]htmlwrapper.Elm{}, BsColSizeAuto, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"col\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestColFormGroup(t *testing.T) {
	// Col(contents []htmlwrapper.Elm, size BsColSize, formGroup bool)
	html, err := Col([]htmlwrapper.Elm{}, BsColSizeAuto, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"col form-group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestColFormGroupColSize(t *testing.T) {
	// Col(contents []htmlwrapper.Elm, size BsColSize, formGroup bool)
	html, err := Col([]htmlwrapper.Elm{}, BsColSize10, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"colcol-10 form-group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestColFormGroupContent(t *testing.T) {
	// Col(contents []htmlwrapper.Elm, size BsColSize, formGroup bool)
	html, err := Col([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, BsColSize10, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"colcol-10 form-group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestFormRow(t *testing.T) {
	// FormRow(contents []htmlwrapper.Elm, formGroup bool)
	html, err := FormRow([]htmlwrapper.Elm{}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"form-row\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormRowFormGroup(t *testing.T) {
	// FormRow(contents []htmlwrapper.Elm, formGroup bool)
	html, err := FormRow([]htmlwrapper.Elm{}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"form-row form-group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormRowFormGroupContent(t *testing.T) {
	// FormRow(contents []htmlwrapper.Elm, formGroup bool)
	html, err := FormRow([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"form-row form-group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
