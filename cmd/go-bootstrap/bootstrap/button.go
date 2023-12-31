package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover) htmlwrapper.Elm {
	attrs := map[string]string{
		"class": "btn " + string(btnType) + " " + string(size),
		"value": text,
		"type":  string(kind),
	}
	if state == BsBtnStateDisabled {
		attrs["disabled"] = "disabled"
		attrs["aria-disabled"] = "true"
	} else if state == BsBtnStateActive {
		attrs["class"] += " active"
	}
	if popover != nil {
		attrs["title"] = popover.Title
		if popover.Content != nil {
			attrs["data-content"] = *popover.Content
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:      "input",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{},
	}
}
func ButtonAnchor(href string, btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover) htmlwrapper.Elm {
	attrs := map[string]string{
		"class": "btn " + string(btnType) + " " + string(size),
		"href":  href,
	}
	if state == BsBtnStateDisabled {
		attrs["disabled"] = "disabled"
		attrs["aria-disabled"] = "true"
	} else if state == BsBtnStateActive {
		attrs["class"] += " active"
	}
	if popover != nil {
		attrs["title"] = popover.Title
		if popover.Content != nil {
			attrs["data-content"] = *popover.Content
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:   "a",
		Attrs: attrs,
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.TextElm{Content: text},
		},
	}
}

func ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize) htmlwrapper.Elm {
	inputType := "radio"
	if checkbox {
		inputType = "checkbox"
	}
	return &htmlwrapper.MultiElm{
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "input",
				Attrs: map[string]string{
					"type":         inputType,
					"class":        "btn-check",
					"id":           id,
					"autocomplete": "off",
				},
				Contents: []htmlwrapper.Elm{},
			},
			&htmlwrapper.HTMLElm{
				Tag: "label",
				Attrs: map[string]string{
					"class": "btn " + string(btnType) + " " + string(size),
					"for":   id,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.TextElm{
						Content: value,
					},
				},
			},
		},
	}
}
