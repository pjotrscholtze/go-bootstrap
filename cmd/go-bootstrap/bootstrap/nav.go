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
		ni.Content,
		&htmlwrapper.MultiElm{Contents: content},
		ni.Href,
	)
}

func Nav(ulElement, vertical bool, justifyContent BsNavJustifyContent, tabKind BsNavKind, content []*NavItem, color *BsColor) htmlwrapper.Elm {
	attrs := map[string]string{
		"class": "navbar-nav " + string(tabKind) + " " + string(justifyContent),
	}
	if color != nil {
		attrs["class"] += " navbar-" + string(*color) + " bg-" + string(*color)
	}
	if vertical {
		attrs["class"] += " flex-column"
	}
	navContents := make([]htmlwrapper.Elm, len(content))
	out := &htmlwrapper.HTMLElm{
		Tag:      "nav",
		Attrs:    attrs,
		Contents: navContents,
	}
	if ulElement {
		out.Tag = "ul"
		for i, item := range content {
			class := "nav-item"
			if item.DropDownItems != nil {
				class += " dropdown"
			}
			navContents[i] = &htmlwrapper.HTMLElm{
				Tag: "li",
				Attrs: map[string]string{
					"class": class,
				},
				Contents: []htmlwrapper.Elm{recursivelyMakeNavLink(item)},
			}
		}
	} else {
		for i, item := range content {
			navContents[i] = recursivelyMakeNavLink(item)
		}
	}
	return out
}

func NavLink(isDropdownToggle bool, navState BsNavState, content htmlwrapper.Elm, dropdownContent htmlwrapper.Elm, href *string) htmlwrapper.Elm {
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
		attrs["data-bs-toggle"] = "dropdown"
		attrs["role"] = "button"
		attrs["aria-haspopup"] = "true"
		attrs["aria-expanded"] = "false"
		return &htmlwrapper.MultiElm{
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.HTMLElm{
					Tag:      "a",
					Attrs:    attrs,
					Contents: []htmlwrapper.Elm{content},
				},
				&htmlwrapper.HTMLElm{
					Tag: "div",
					Attrs: map[string]string{
						"class": "dropdown-menu",
					},
					Contents: []htmlwrapper.Elm{dropdownContent},
				},
			},
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:      "a",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{content},
	}
}
