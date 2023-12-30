package bootstrap

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type DrowndownItem struct {
	// Divider bool
	Kind DrowndownItemKind
	Href string
	Text string
}

func Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem) htmlwrapper.Elm {
	itemElems := make([]htmlwrapper.Elm, len(items))
	for i, item := range items {
		if item.Kind == BsDrowndownItemKindDivider {
			itemElems[i] = &htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"class": "dropdown-divider",
				},
			}
		} else if item.Kind == BsDrowndownItemKindItem {
			itemElems[i] = &htmlwrapper.HTMLElm{
				Tag: "a",
				Attrs: map[string]string{
					"class": "dropdown-item",
					"href":  item.Href,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.TextElm{
						Content: item.Text,
					},
				},
			}
		} else if item.Kind == BsDrowndownItemKindItemActive {
			itemElems[i] = &htmlwrapper.HTMLElm{
				Tag: "a",
				Attrs: map[string]string{
					"class": "dropdown-item active",
					"href":  item.Href,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.TextElm{
						Content: item.Text,
					},
				},
			}
		} else if item.Kind == BsDrowndownItemKindItemDisabled {
			itemElems[i] = &htmlwrapper.HTMLElm{
				Tag: "a",
				Attrs: map[string]string{
					"class": "dropdown-item disabled",
					"href":  item.Href,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.TextElm{
						Content: item.Text,
					},
				},
			}
		} else {
			itemElems[i] = &htmlwrapper.HTMLElm{
				Tag: "h6",
				Attrs: map[string]string{
					"class": "dropdown-header",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.TextElm{
						Content: item.Text,
					},
				},
			}
		}
	}
	btn := Button(btnType, btnKind, btnText, btnSize, btnState, nil).(*htmlwrapper.HTMLElm)
	btn.Attrs["class"] += " dropdown-toggle"
	btn.Attrs["id"] = "dropdownMenuButton"
	btn.Attrs["type"] = "button"
	btn.Attrs["data-bs-toggle"] = "dropdown"
	btn.Attrs["aria-expanded"] = "false"
	btn.Tag = "button"
	btn.Contents = []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: btnText}}
	out := &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "dropdown",
		},
		Contents: []htmlwrapper.Elm{
			btn,
			&htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"class":           "dropdown-menu",
					"aria-labelledby": "dropdownMenuButton",
				},
				Contents: itemElems,
			},
		},
	}
	if direction != BsDropdownDirectionDownDown {
		out = &htmlwrapper.HTMLElm{
			Tag: "div",
			Attrs: map[string]string{
				"class": "btn-group dropleft",
			},
			Contents: []htmlwrapper.Elm{
				out,
			},
		}
	}
	return out
}
