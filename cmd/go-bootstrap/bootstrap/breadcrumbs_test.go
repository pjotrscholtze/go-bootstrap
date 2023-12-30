package bootstrap

import "testing"

func TestBreadcrumb(t *testing.T) {
	// Breadcrumb(items []BreadcrumbEntry)
	html, err := Breadcrumb([]BreadcrumbEntry{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav aria-label=\"breadcrumb\"><ol class=\"breadcrumb\"></ol></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBreadcrumbSingleItem(t *testing.T) {
	// Breadcrumb(items []BreadcrumbEntry)
	html, err := Breadcrumb([]BreadcrumbEntry{{
		Href: "href",
		Text: "text",
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav aria-label=\"breadcrumb\"><ol class=\"breadcrumb\"><li aria-current=\"page\" class=\"breadcrumb-item active\">text</li></ol></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestBreadcrumbMultipleItem(t *testing.T) {
	// Breadcrumb(items []BreadcrumbEntry)
	html, err := Breadcrumb([]BreadcrumbEntry{
		{
			Href: "href",
			Text: "text",
		},
		{
			Href: "href2",
			Text: "text2",
		}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav aria-label=\"breadcrumb\"><ol class=\"breadcrumb\"><li aria-current=\"page\" class=\"breadcrumb-item\"><a href=\"href\">text</a></li><li aria-current=\"page\" class=\"breadcrumb-item active\">text2</li></ol></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}
