package bootstrap

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func Collapse(btn *htmlwrapper.HTMLElm, id, target string) htmlwrapper.Elm {
	btn.Attrs["data-toggle"] = "collapse"
	if btn.Tag == "button" {
		btn.Attrs["data-target"] = target
	} else {
		btn.Attrs["href"] = target
		btn.Attrs["role"] = "button"
	}
	btn.Attrs["aria-expanded"] = "false"
	btn.Attrs["aria-controls"] = id

	return &htmlwrapper.MultiElm{
		Contents: []htmlwrapper.Elm{
			btn,
			&htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"class": "collapse",
					"id":    id,
				},
			},
		},
	}
}
