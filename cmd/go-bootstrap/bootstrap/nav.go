package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

type NavItem struct {
	Href          *string
	Content       htmlwrapper.Elm
	DropDownItems *[]*NavItem
	NavState      BsNavState
}

func recursivelyMakeNavLink(ni *NavItem) htmlwrapper.Elm {
	if ni == nil {
		return &htmlwrapper.HTMLElm{
			Tag:   "div",
			Attrs: map[string]string{"class": "dropdown-divider"},
		}
	}
	content := make([]htmlwrapper.Elm, len(*ni.DropDownItems))
	if ni.DropDownItems != nil {
		for i, item := range *ni.DropDownItems {
			content[i] = recursivelyMakeNavLink(item)
		}
	}
	return NavLink(
		len(*ni.DropDownItems) > 0,
		ni.NavState,
		&htmlwrapper.MultiElm{Contents: content},
		ni.Href,
	)
}

func Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem) htmlwrapper.Elm {
	attrs := map[string]string{
		"class": "nav " + string(tabKind) + " " + string(justifyContent),
	}
	if vertical {
		attrs["class"] += " flex-column"
	}
	navContents := make([]htmlwrapper.Elm, len(content))
	tag := "nav"
	for i, item := range content {
		navContents[i] = recursivelyMakeNavLink(item)
	}
	if ulElement {
		tag = "ul"
		for i, item := range navContents {
			navContents[i] = &htmlwrapper.HTMLElm{
				Tag: "li",
				Attrs: map[string]string{
					"class": "nav-item",
				},
				Contents: []htmlwrapper.Elm{item},
			}
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:      tag,
		Attrs:    attrs,
		Contents: navContents,
	}
}

func NavLink(isDropdownToggle bool, navState BsNavState, content htmlwrapper.Elm, href *string) htmlwrapper.Elm {
	addClass := ""
	if navState != BsNavStateNormal {
		addClass += " " + string(navState)
	}
	hrefStr := "#"
	if href != nil {
		hrefStr = *href
	}
	attrs := map[string]string{
		"class": "nav-link" + addClass,
		"href":  hrefStr,
	}
	if isDropdownToggle {
		attrs["class"] += " dropdown-toggle"
		attrs["data-toggle"] = "dropdown"
		attrs["role"] = "button"
		attrs["aria-haspopup"] = "true"
		attrs["aria-expanded"] = "false"
		content = &htmlwrapper.HTMLElm{
			Tag: "div",
			Attrs: map[string]string{
				"class": "dropdown-menu",
			},
			Contents: []htmlwrapper.Elm{content},
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:      "a",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{content},
	}
}
