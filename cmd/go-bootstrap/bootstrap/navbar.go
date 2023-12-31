package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func NavBarBrand(content htmlwrapper.Elm, href *string) htmlwrapper.Elm {
	attrs := map[string]string{
		"class": "navbar-brand",
	}
	if href == nil {
		return &htmlwrapper.HTMLElm{
			Tag:      "span",
			Attrs:    attrs,
			Contents: []htmlwrapper.Elm{content},
		}
	}
	attrs["href"] = *href
	return &htmlwrapper.HTMLElm{
		Tag:      "a",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{content},
	}

}

// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
type CollapseButton struct {
	BtnType BsBtnStyle
	Kind    BsBtnKind
	Size    BsBtnSize
	State   BsBtnState
	Popover *TooltipPopover
}

func NavBar(id string, color BsColor, location BsLocation, collapsable *CollapseButton, content htmlwrapper.Elm) htmlwrapper.Elm {
	if collapsable != nil {
		btn := Button(collapsable.BtnType, collapsable.Kind, "", collapsable.Size, collapsable.State, collapsable.Popover).(*htmlwrapper.HTMLElm)
		btn.Contents = []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "span",
				Attrs: map[string]string{
					"class": "navbar-toggler-icon",
				}}}

		btn.Attrs["data-toggle"] = "collapse"
		btn.Attrs["data-target"] = "#" + id
		btn.Attrs["aria-controls"] = id
		btn.Attrs["aria-expanded"] = "false"

		content = &htmlwrapper.MultiElm{
			Contents: []htmlwrapper.Elm{
				btn,
				&htmlwrapper.HTMLElm{
					Tag: "div",
					Attrs: map[string]string{
						"class": "collapse navbar-collapse",
						"id":    id,
					},
					Contents: []htmlwrapper.Elm{content},
				},
			},
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag: "nav",
		Attrs: map[string]string{
			"class": "navbar navbar-expand-lg navbar-" + string(color) + " bg-" + string(color),
		},
		Contents: []htmlwrapper.Elm{&htmlwrapper.HTMLElm{
			Tag: "div",
			Attrs: map[string]string{
				"class": "container-fluid",
			},
			Contents: []htmlwrapper.Elm{content},
		}},
	}
}

func NavBarText(content htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "span",
		Attrs: map[string]string{
			"class": "navbar-text",
		},
		Contents: []htmlwrapper.Elm{content},
	}
}
