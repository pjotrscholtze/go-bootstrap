package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize) htmlwrapper.Elm {
	baseClass := "btn-group"
	if vertical {
		baseClass = "btn-group-vertical"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class":      baseClass + " " + string(size),
			"role":       "group",
			"aria-label": ariaLabel,
		},
		Contents: content,
	}
}

func ButtonToolbar(content []htmlwrapper.Elm, ariaLabel string) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class":      "btn-toolbar",
			"role":       "toolbar",
			"aria-label": ariaLabel,
		},
		Contents: content,
	}
}
