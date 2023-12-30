package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestAlert(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertPrimary, []htmlwrapper.Elm{}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-primary\" role=\"alert\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertDismissible(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertPrimary, []htmlwrapper.Elm{}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-primary alert-dismissible\" role=\"alert\"><button aria-label=\"Close\" class=\"btn-close\" data-bs-dismiss=\"alert\" type=\"button\"></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertContent(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-primary\" role=\"alert\">text</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertDismissibleContent(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-primary alert-dismissible\" role=\"alert\">text<button aria-label=\"Close\" class=\"btn-close\" data-bs-dismiss=\"alert\" type=\"button\"></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertSecondary(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertSecondary, []htmlwrapper.Elm{}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-secondary\" role=\"alert\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertDismissibleSecondary(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertSecondary, []htmlwrapper.Elm{}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-secondary alert-dismissible\" role=\"alert\"><button aria-label=\"Close\" class=\"btn-close\" data-bs-dismiss=\"alert\" type=\"button\"></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertContentSecondary(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertSecondary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-secondary\" role=\"alert\">text</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestAlertDismissibleContentSecondary(t *testing.T) {
	// Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool)
	html, err := Alert(AlertSecondary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"alert alert-secondary alert-dismissible\" role=\"alert\">text<button aria-label=\"Close\" class=\"btn-close\" data-bs-dismiss=\"alert\" type=\"button\"></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}
