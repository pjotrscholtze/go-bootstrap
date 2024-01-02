package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestCollapse(t *testing.T) {
	// Collapse(btn *htmlwrapper.HTMLElm, id, target string)
	html, err := Collapse(Button(BsBtnPrimary, BsBtnStateNormal, "text", BsBtnLg, BsBtnStateDisabled, nil).(*htmlwrapper.HTMLElm), "id", "target").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input aria-controls=\"id\" aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-primary btn-lg\" data-toggle=\"collapse\" disabled=\"disabled\" href=\"target\" role=\"button\" value=\"text\"></input><div class=\"collapse\" id=\"id\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCollapseIdAndTarget(t *testing.T) {
	// Collapse(btn *htmlwrapper.HTMLElm, id, target string)
	html, err := Collapse(Button(BsBtnPrimary, BsBtnStateNormal, "text", BsBtnLg, BsBtnStateDisabled, nil).(*htmlwrapper.HTMLElm), "fdsa", "asdf").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input aria-controls=\"fdsa\" aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-primary btn-lg\" data-toggle=\"collapse\" disabled=\"disabled\" href=\"asdf\" role=\"button\" value=\"text\"></input><div class=\"collapse\" id=\"fdsa\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
