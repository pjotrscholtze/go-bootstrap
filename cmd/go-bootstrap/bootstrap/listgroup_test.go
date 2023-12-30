package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

// type ListGroupItem struct {
// 	Content        htmlwrapper.Elm
// 	Active         bool
// 	Disabled       bool
// 	JustifyContent bool
// 	Href           *string
// 	Kind           *ListGroupItemKind
// }

func TestListgroup(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestListgroupFlush(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{}, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group list-group-flush\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestListgroupContent(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: ""},
		Active:         false,
		Disabled:       false,
		JustifyContent: false,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item\"></li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestListgroupContentText(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         false,
		Disabled:       false,
		JustifyContent: false,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item\">Text</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestListgroupContentTextActive(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         true,
		Disabled:       false,
		JustifyContent: false,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item active\">Text</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestListgroupContentTextDisabled(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         false,
		Disabled:       true,
		JustifyContent: false,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item disabled\">Text</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestListgroupContentTextDisabledActive(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         true,
		Disabled:       true,
		JustifyContent: false,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item active disabled\">Text</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestListgroupContentTextJustifyContent(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         false,
		Disabled:       false,
		JustifyContent: true,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item justify-content-between align-items-center\">Text</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestListgroupContentTextJustifyContentActive(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         true,
		Disabled:       false,
		JustifyContent: true,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item active justify-content-between align-items-center\">Text</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestListgroupContentMulti(t *testing.T) {
	// Listgroup(content []ListGroupItem, flush bool)
	html, err := Listgroup([]ListGroupItem{{
		Content:        &htmlwrapper.TextElm{Content: "Text"},
		Active:         true,
		Disabled:       false,
		JustifyContent: true,
	}, {
		Content:        &htmlwrapper.TextElm{Content: "ABCDEFGH"},
		Active:         false,
		Disabled:       true,
		JustifyContent: false,
	}}, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"list-group\"><li class=\"list-group-item active justify-content-between align-items-center\">Text</li><li class=\"list-group-item disabled\">ABCDEFGH</li></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
