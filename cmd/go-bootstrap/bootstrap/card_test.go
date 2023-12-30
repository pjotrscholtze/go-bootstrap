package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestCardImg(t *testing.T) {
	// CardImg(src string, alt *string)
	html, err := CardImg("src", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<img class=\"card-img-top\" src=\"src\"></img>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardImgAlt(t *testing.T) {
	// CardImg(src string, alt *string)
	alt := "alt"
	html, err := CardImg("src", &alt).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<img alt=\"alt\" class=\"card-img-top\" src=\"src\"></img>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardText(t *testing.T) {
	// CardText(content []htmlwrapper.Elm)
	html, err := CardText([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<p class=\"card-text\"></p>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardTextContent(t *testing.T) {
	// CardText(content []htmlwrapper.Elm)
	html, err := CardText([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<p class=\"card-text\">content</p>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardFooter(t *testing.T) {
	// CardFooter(content []htmlwrapper.Elm)
	html, err := CardFooter([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-footer\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardFooterContent(t *testing.T) {
	// CardFooter(content []htmlwrapper.Elm)
	html, err := CardFooter([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "text"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-footer\">text</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardBody(t *testing.T) {
	// CardBody(title *string, content []htmlwrapper.Elm)
	html, err := CardBody(nil, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-body\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestCardBodyTitle(t *testing.T) {
	// CardBody(title *string, content []htmlwrapper.Elm)
	title := "title"
	html, err := CardBody(&title, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-body\"><h5 class=\"card-title\">title</h5></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestCardBodyContent(t *testing.T) {
	// CardBody(title *string, content []htmlwrapper.Elm)
	html, err := CardBody(nil, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-body\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestCardBodyTitleContent(t *testing.T) {
	// CardBody(title *string, content []htmlwrapper.Elm)
	title := "title"
	html, err := CardBody(&title, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-body\">content<h5 class=\"card-title\">title</h5></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCard(t *testing.T) {
	// Card(content []htmlwrapper.Elm)
	html, err := Card([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestCardContent(t *testing.T) {
	// Card(content []htmlwrapper.Elm)
	html, err := Card([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestCardImgOverlay(t *testing.T) {
	// CardImgOverlay(content []htmlwrapper.Elm)
	html, err := CardImgOverlay([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-img-overlay\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCardImgOverlayContent(t *testing.T) {
	// CardImgOverlay(content []htmlwrapper.Elm)
	html, err := CardImgOverlay([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"card-img-overlay\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
