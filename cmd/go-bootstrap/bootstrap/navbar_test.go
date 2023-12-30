package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestNavBarBrand(t *testing.T) {
	// NavBarBrand(content htmlwrapper.Elm, href *string)
	html, err := NavBarBrand(&htmlwrapper.TextElm{Content: ""}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"navbar-brand\"></span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavBarBrandContent(t *testing.T) {
	// NavBarBrand(content htmlwrapper.Elm, href *string)
	html, err := NavBarBrand(&htmlwrapper.TextElm{Content: "a"}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"navbar-brand\">a</span>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavBarBrandContentLink(t *testing.T) {
	// NavBarBrand(content htmlwrapper.Elm, href *string)
	href := "#"
	html, err := NavBarBrand(&htmlwrapper.TextElm{Content: "a"}, &href).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"navbar-brand\" href=\"#\">a</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavBar(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorPrimary, BsLocationNormal, nil, &htmlwrapper.TextElm{Content: ""}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-primary bg-primary\"><div class=\"container-fluid\"></div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestNavBarContent(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorPrimary, BsLocationNormal, nil, &htmlwrapper.TextElm{Content: "text"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-primary bg-primary\"><div class=\"container-fluid\">text</div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}
func TestNavBarContentDiffBackground(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorSecondary, BsLocationNormal, nil, &htmlwrapper.TextElm{Content: "text"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-secondary bg-secondary\"><div class=\"container-fluid\">text</div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestNavBarContentDiffBackgroundFixedBottom(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorSecondary, BsLocationFixedBottom, nil, &htmlwrapper.TextElm{Content: "text"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-secondary bg-secondary\"><div class=\"container-fluid\">text</div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestNavBarContentDiffBackgroundFixedTop(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorSecondary, BsLocationFixedTop, nil, &htmlwrapper.TextElm{Content: "text"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-secondary bg-secondary\"><div class=\"container-fluid\">text</div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestNavBarContentDiffBackgroundStickyTop(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorSecondary, BsLocationStickyTop, nil, &htmlwrapper.TextElm{Content: "text"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-secondary bg-secondary\"><div class=\"container-fluid\">text</div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestNavBarContentDiffBackgroundStickyTopCollapseButton(t *testing.T) {
	// NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm)
	html, err := NavBar("id", BsColorSecondary, BsLocationStickyTop, &CollapseButton{
		BtnType: BsBtnPrimary,
		Kind:    BsBtnNormal,
		Size:    btnGroupSm,
		State:   BsBtnStateDisabled,
		Popover: nil,
	}, &htmlwrapper.TextElm{Content: "text"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-expand-lg navbar-secondary bg-secondary\"><div class=\"container-fluid\"><input aria-controls=\"id\" aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-primary btn-group-sm\" data-target=\"#id\" data-toggle=\"collapse\" disabled=\"disabled\" value=\"\"><span class=\"navbar-toggler-icon\"></span></input><div class=\"collapse navbar-collapse\" id=\"id\">text</div></div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestNavBarText(t *testing.T) {
	// NavBarText(content htmlwrapper.Elm)
	html, err := NavBarText(&htmlwrapper.TextElm{Content: "test"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<span class=\"navbar-text\">test</span>" {
		t.Fatalf("HTML is not as expected!")
	}
}
