package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestFormHelper(t *testing.T) {
	// FormHelper(id string, contents []htmlwrapper.Elm)
	html, err := FormHelper("id", []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<small class=\"form-text text-muted\" id=\"id\"></small>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormHelperContent(t *testing.T) {
	// FormHelper(id string, contents []htmlwrapper.Elm)
	html, err := FormHelper("id", []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<small class=\"form-text text-muted\" id=\"id\">content</small>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormLabel(t *testing.T) {
	// FormLabel(forId string, contents []htmlwrapper.Elm)
	html, err := FormLabel("forId", []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<small class=\"form-text text-muted\" for=\"forId\"></small>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormLabelContent(t *testing.T) {
	// FormLabel(forId string, contents []htmlwrapper.Elm)
	html, err := FormLabel("forId", []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<small class=\"form-text text-muted\" for=\"forId\">text</small>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormInvalidFeedback(t *testing.T) {
	// FormInvalidFeedback(contents []htmlwrapper.Elm)
	html, err := FormInvalidFeedback([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"invalid-feedback\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormInvalidFeedbackContent(t *testing.T) {
	// FormInvalidFeedback(contents []htmlwrapper.Elm)
	html, err := FormInvalidFeedback([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"invalid-feedback\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormTooltip(t *testing.T) {
	// FormTooltip(valid bool, contents []htmlwrapper.Elm)
	html, err := FormTooltip(false, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"invalid-tooltip\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormTooltipValid(t *testing.T) {
	// FormTooltip(valid bool, contents []htmlwrapper.Elm)
	html, err := FormTooltip(true, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"valid-tooltip\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestFormTooltipValidContent(t *testing.T) {
	// FormTooltip(valid bool, contents []htmlwrapper.Elm)
	html, err := FormTooltip(true, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"valid-tooltip\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
