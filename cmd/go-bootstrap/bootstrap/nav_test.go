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
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  \"><a class=\"nav-link\" href=\"#\"></a></nav>" {
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
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  \"><a class=\"nav-link\" href=\"link\"></a></nav>" {
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
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  \"><a class=\"nav-link disabled\" href=\"link\"></a></nav>" {
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
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  \"><a class=\"nav-link active\" href=\"link\"></a></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNav(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  \"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUl(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav  \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVertical(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentLeft, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav   flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  justify-content-center\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav  justify-content-center\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav  justify-content-center flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlVerticalJustifyCenter(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, true, BsNavJustifyContentCenter, BsNavKindNormal, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav  justify-content-center flex-column\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-pills \"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentLeft, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav nav-pills \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentLeft, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-pills  flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-pills justify-content-center\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav nav-pills justify-content-center\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-pills justify-content-center flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlVerticalJustifyCenterPills(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, true, BsNavJustifyContentCenter, BsNavKindPills, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav nav-pills justify-content-center flex-column\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentLeft, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-tabs \"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentLeft, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav nav-tabs \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentLeft, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-tabs  flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, false, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-tabs justify-content-center\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, false, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav nav-tabs justify-content-center\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavVerticalJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(false, true, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<nav class=\"nav nav-tabs justify-content-center flex-column\"></nav>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavUlVerticalJustifyCenterTabs(t *testing.T) {
	// Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, popover *TooltipPopover)
	html, err := Nav(true, true, BsNavJustifyContentCenter, BsNavKindTabs, []*NavItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"nav nav-tabs justify-content-center flex-column\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLink(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(false, BsNavStateNormal, &htmlwrapper.TextElm{Content: ""}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"nav-link\" href=\"#\"></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdown(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateNormal, &htmlwrapper.TextElm{Content: ""}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link dropdown-toggle\" data-toggle=\"dropdown\" href=\"#\" role=\"button\"><div class=\"dropdown-menu\"></div></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdownContent(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateNormal, &htmlwrapper.TextElm{Content: "Text"}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link dropdown-toggle\" data-toggle=\"dropdown\" href=\"#\" role=\"button\"><div class=\"dropdown-menu\">Text</div></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkContent(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(false, BsNavStateNormal, &htmlwrapper.TextElm{Content: "Text"}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"nav-link\" href=\"#\">Text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdownContentActive(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateActive, &htmlwrapper.TextElm{Content: "Text"}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-toggle=\"dropdown\" href=\"#\" role=\"button\"><div class=\"dropdown-menu\">Text</div></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkDropdownActive(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(true, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-toggle=\"dropdown\" href=\"#\" role=\"button\"><div class=\"dropdown-menu\"></div></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestNavLinkActive(t *testing.T) {
	// NavLink(isDropdownToggle bool, navState *BsNavState, content htmlwrapper.Elm, href *string, popover *TooltipPopover)
	html, err := NavLink(false, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, nil).AsHTML()
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
	html, err := NavLink(false, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, &href).AsHTML()
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
	html, err := NavLink(true, BsNavStateActive, &htmlwrapper.TextElm{Content: ""}, &href).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a aria-expanded=\"false\" aria-haspopup=\"true\" class=\"nav-link active dropdown-toggle\" data-toggle=\"dropdown\" href=\"test\" role=\"button\"><div class=\"dropdown-menu\"></div></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}
