package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestScrollspy(t *testing.T) {
	html, err := Scrollspy("target", "0", []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != `<div data-offset="0" data-spy="scroll" data-target="#target"></div>` {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestScrollspyTarget(t *testing.T) {
	html, err := Scrollspy("differentTarget", "0", []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != `<div data-offset="0" data-spy="scroll" data-target="#differentTarget"></div>` {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestScrollspyOffset(t *testing.T) {
	html, err := Scrollspy("target", "1234", []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != `<div data-offset="1234" data-spy="scroll" data-target="#target"></div>` {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestScrollspyContent(t *testing.T) {
	html, err := Scrollspy("target", "0", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "test"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != `<div data-offset="0" data-spy="scroll" data-target="#target">test</div>` {
		t.Fatalf("HTML is not as expected!")
	}
}
