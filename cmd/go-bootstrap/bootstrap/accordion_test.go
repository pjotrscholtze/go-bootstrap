package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestAccordionItem(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "id", "parentId", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"id\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeading(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "id", "parentId", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"id\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemId(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "idTest", "parentId", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"idTest\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingId(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "idTest", "parentId", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"idTest\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemParent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "id", "parentIdTest", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"id\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingParent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "id", "parentIdTest", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"id\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemIdParent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "idTest", "parentIdTest", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"idTest\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingIdParent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "idTest", "parentIdTest", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"idTest\"><div class=\"accordion-body\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "id", "parentId", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"id\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "id", "parentId", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"id\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemIdContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "idTest", "parentId", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"idTest\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingIdContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "idTest", "parentId", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentId\" id=\"idTest\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemParentContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "id", "parentIdTest", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"id\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingParentContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "id", "parentIdTest", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-id\"><button aria-controls=\"id\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#id\" data-bs-toggle=\"collapse\" id=\"id\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-id\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"id\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemIdParentContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("heading", "idTest", "parentIdTest", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">heading</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"idTest\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionItemHeadingIdParentContent(t *testing.T) {
	// AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm)
	html, err := AccordionItem("headingBlah", "idTest", "parentIdTest", []htmlwrapper.Elm{
		&htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion-item\"><h2 class=\"accordion-header\" id=\"heading-idTest\"><button aria-controls=\"idTest\" aria-expanded=\"true\" class=\"accordion-button\" data-bs-target=\"#idTest\" data-bs-toggle=\"collapse\" id=\"idTest\" type=\"button\">headingBlah</button></h2><div aria-labelledby=\"heading-idTest\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#parentIdTest\" id=\"idTest\"><div class=\"accordion-body\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordion(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "id", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"id\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionFlush(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "id", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"id\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionId(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "idTest", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"idTest\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionFlushId(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "idTest", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"idTest\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionContent(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "id", []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"id\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionFlushContent(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "id", []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"id\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionIdContent(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "idTest", []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"idTest\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestAccordionFlushIdContent(t *testing.T) {
	// Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm)
	html, err := Accordion(false, "idTest", []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"accordion\" id=\"idTest\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
