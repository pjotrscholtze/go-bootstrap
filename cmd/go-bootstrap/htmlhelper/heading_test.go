package htmlhelper

import "testing"

func TestHeading1(t *testing.T) {
	html, err := Heading(1, "text").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h1>text</h1>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h1>text</h1>", got "%q"`, html)
	}
}

func TestHeading1Content(t *testing.T) {
	html, err := Heading(1, "lorem ipsum").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h1>lorem ipsum</h1>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h1>lorem ipsum</h1>", got "%q"`, html)
	}
}

func TestHeading2(t *testing.T) {
	html, err := Heading(2, "text").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h2>text</h2>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h2>text</h2>", got "%q"`, html)
	}
}

func TestHeading2Content(t *testing.T) {
	html, err := Heading(2, "lorem ipsum").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h2>lorem ipsum</h2>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h2>lorem ipsum</h2>", got "%q"`, html)
	}
}

func TestHeading7(t *testing.T) {
	html, err := Heading(7, "text").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h7>text</h7>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h7>text</h7>", got "%q"`, html)
	}
}

func TestHeading7Content(t *testing.T) {
	html, err := Heading(7, "lorem ipsum").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h7>lorem ipsum</h7>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h7>lorem ipsum</h7>", got "%q"`, html)
	}
}

func TestHeading8(t *testing.T) {
	html, err := Heading(8, "text").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h1>text</h1>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h1>text</h1>", got "%q"`, html)
	}
}

func TestHeading8Content(t *testing.T) {
	html, err := Heading(8, "lorem ipsum").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h1>lorem ipsum</h1>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h1>lorem ipsum</h1>", got "%q"`, html)
	}
}

func TestHeading0(t *testing.T) {
	html, err := Heading(0, "text").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h1>text</h1>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h1>text</h1>", got "%q"`, html)
	}
}

func TestHeading0Content(t *testing.T) {
	html, err := Heading(0, "lorem ipsum").AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<h1>lorem ipsum</h1>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<h1>lorem ipsum</h1>", got "%q"`, html)
	}
}
