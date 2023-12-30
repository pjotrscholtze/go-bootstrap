package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestBadge(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgPrimary, []htmlwrapper.Elm{}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-primary\"></span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeRoundedPill(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgPrimary, []htmlwrapper.Elm{}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-primary rounded-pill\"></span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeSecondary(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgSecondary, []htmlwrapper.Elm{}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-secondary\"></span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeRoundedPillSecondary(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgSecondary, []htmlwrapper.Elm{}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-secondary rounded-pill\"></span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeContent(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-primary\">text</span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeRoundedPillContent(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-primary rounded-pill\">text</span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeSecondaryContent(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgSecondary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-secondary\">text</span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBadgeRoundedPillSecondaryContent(t *testing.T) {
	// Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool)
	html, err := Badge(BsTextBgSecondary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"badge text-bg-secondary rounded-pill\">text</span>" {
		t.Fatalf("HTML is not as expected!")
	}
}
