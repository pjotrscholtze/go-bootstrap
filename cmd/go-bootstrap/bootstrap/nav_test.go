package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestNavContent(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{{
		Content:       &htmlwrapper.TextElm{Content: "test"},
		DropDownItems: &[]*NavItem{},
	}}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  \"><a class=\"nav-link\" href=\"#\">test</a></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestNavContentHref(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	href := "link"
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{{
		Content:       &htmlwrapper.TextElm{Content: "test"},
		DropDownItems: &[]*NavItem{},
		Href:          &href,
	}}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  \"><a class=\"nav-link\" href=\"link\">test</a></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestNavContentHrefDisabled(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	href := "link"
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{{
		Content:       &htmlwrapper.TextElm{Content: "test"},
		DropDownItems: &[]*NavItem{},
		Href:          &href,
		NavState:      BsNavStateDisabled,
	}}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  \"><a class=\"nav-link disabled\" href=\"link\">test</a></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestNavContentHrefActive(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	href := "link"
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{{
		Content:       &htmlwrapper.TextElm{Content: "test"},
		DropDownItems: &[]*NavItem{},
		Href:          &href,
		NavState:      BsNavStateActive,
	}}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  \"><a class=\"nav-link active\" href=\"link\">test</a></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestNavContentHrefActiveColor(t *testing.T) {
	c := BsColorPrimary
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	href := "link"
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{{
		Content: &htmlwrapper.TextElm{Content: "test"},
		DropDownItems: &[]*NavItem{
			{
				Href:          &href,
				Content:       htmlwrapper.Text("asd"),
				DropDownItems: &[]*NavItem{},
				NavState:      BsNavStateActive,
			},
		},
		Href:     &href,
		NavState: BsNavStateActive,
	}}, &c).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav   navbar-primary bg-primary\"><a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-bs-toggle=\"dropdown\" href=\"link\" role=\"button\">test</a><div class=\"dropdown-menu navbar-primary bg-primary\"><a class=\"nav-link active\" href=\"link\">asd</a></div></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNav(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  \"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUl(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav  \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVertical(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav   flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  justify-content-center\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav  justify-content-center\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav  justify-content-center flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlVerticalJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, true, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav  justify-content-center flex-column\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-pills \"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentLeft, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav nav-pills \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentLeft, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-pills  flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-pills justify-content-center\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav nav-pills justify-content-center\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-pills justify-content-center flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlVerticalJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, true, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav nav-pills justify-content-center flex-column\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-tabs \"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentLeft, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav nav-tabs \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentLeft, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-tabs  flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-tabs justify-content-center\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav nav-tabs justify-content-center\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"navbar-nav nav-tabs justify-content-center flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlVerticalJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, true, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"navbar-nav nav-tabs justify-content-center flex-column\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLink(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(false, BsNavStateNormal, &htmlwrapper.TextElm{Content: ""}, nil, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"nav-link\" href=\"#\"></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdown(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateNormal, &htmlwrapper.TextElm{Content: ""}, &htmlwrapper.TextElm{Content: ""}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link dropdown-toggle\" data-bs-toggle=\"dropdown\" href=\"#\" role=\"button\"></a><div class=\"dropdown-menu\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdownContent(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateNormal, &htmlwrapper.TextElm{Content: "Text"}, &htmlwrapper.TextElm{Content: "hoi"}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link dropdown-toggle\" data-bs-toggle=\"dropdown\" href=\"#\" role=\"button\">Text</a><div class=\"dropdown-menu\">hoi</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkContent(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(false, BsNavStateNormal, &htmlwrapper.TextElm{Content: "Text"}, nil, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"nav-link\" href=\"#\">Text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdownContentActive(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateActive, &htmlwrapper.TextElm{Content: "Text"}, &htmlwrapper.TextElm{Content: "hoi"}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-bs-toggle=\"dropdown\" href=\"#\" role=\"button\">Text</a><div class=\"dropdown-menu\">hoi</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdownActive(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, &htmlwrapper.TextElm{Content: "asdf"}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-bs-toggle=\"dropdown\" href=\"#\" role=\"button\"></a><div class=\"dropdown-menu\">asdf</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkActive(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(false, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, nil, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"nav-link active\" href=\"#\"></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkActiveHref(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	href := "test"
	html, err := NavLink(false, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, nil, &href, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"nav-link active\" href=\"test\"></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkActiveHrefDropdown(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	href := "test"
	html, err := NavLink(true, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, &htmlwrapper.TextElm{Content: "asdf"}, &href, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-bs-toggle=\"dropdown\" href=\"test\" role=\"button\"></a><div class=\"dropdown-menu\">asdf</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
