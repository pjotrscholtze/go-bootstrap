package htmlwrapper_test

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestTextElmAsHTMLEmptyContent(t *testing.T) {
	te := htmlwrapper.TextElm{
		Content: "",
	}
	if te.Content != "" {
		t.Fatalf(`content of TextElm changed`)
	}
	html, err := te.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on empty text element, should not crash`)
	}
	if html != "" {
		t.Fatalf(`Content of empty text element should be empty got "%q"`, html)
	}
}

func TestTextElmAsHTMLTextContent(t *testing.T) {
	te := htmlwrapper.TextElm{
		Content: "Hoi",
	}
	if te.Content != "Hoi" {
		t.Fatalf(`content of TextElm changed`)
	}
	html, err := te.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on text element, should not crash, %v`, err)
	}
	if html != "Hoi" {
		t.Fatalf(`Content of text element should be "Hoi" got "%q"`, html)
	}
}

func TestTextElmAsHTMLTextEscapeContent(t *testing.T) {
	te := htmlwrapper.TextElm{
		Content: "Ho\"i",
	}
	if te.Content != "Ho\"i" {
		t.Fatalf(`content of TextElm changed`)
	}
	html, err := te.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on text element, should not crash, %v`, err)
	}
	if html != "Ho&#34;i" {
		t.Fatalf(`Content of text element should be "Ho&#34;i" got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElement(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "span",
	}
	html, err := he.AsHTML()
	println(html)
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<span></span>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<span></span>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementEscaped(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "spa\"n",
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<spa&#34;n></spa&#34;n>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<spa&#34;n></spa&#34;n>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementAttr(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "a",
		Attrs: map[string]string{
			"href": "#",
		},
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<a href=\"#\"></a>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<a href=\"#\"></a>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementAttrEscaped1(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "a\"",
		Attrs: map[string]string{
			"hr\"ef": "#\"",
		},
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<a&#34; hr&#34;ef=\"#&#34;\"></a&#34;>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<a&#34; hr&#34;ef=\"#&#34;\"></a&#34;>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementAttrEscaped2(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "a\"",
		Attrs: map[string]string{
			"href": "#\"",
		},
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<a&#34; href=\"#&#34;\"></a&#34;>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<a&#34; href=\"#&#34;\"></a&#34;>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementAttrEscaped3(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "a\"",
		Attrs: map[string]string{
			"hr\"ef": "#",
		},
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<a&#34; hr&#34;ef=\"#\"></a&#34;>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<a&#34; hr&#34;ef=\"#\"></a&#34;>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementAttrEscaped4(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "a",
		Attrs: map[string]string{
			"hr\"ef": "#\"",
		},
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<a hr&#34;ef=\"#&#34;\"></a>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<a hr&#34;ef=\"#&#34;\"></a>", got "%q"`, html)
	}
}

func TestHTMLElmAsHTMLSimpleElementAttrEscaped5(t *testing.T) {
	he := htmlwrapper.HTMLElm{
		Tag: "a\"",
		Attrs: map[string]string{
			"href": "#\"",
		},
	}
	html, err := he.AsHTML()
	if err != nil {
		t.Fatalf(`AsHTML() on HTML element, should not crash, %v`, err)
	}
	if html != "<a&#34; href=\"#&#34;\"></a&#34;>" {
		t.Fatalf(`AsHTML() on HTML element, expected "<a&#34; href=\"#&#34;\"></a&#34;>", got "%q"`, html)
	}
}
