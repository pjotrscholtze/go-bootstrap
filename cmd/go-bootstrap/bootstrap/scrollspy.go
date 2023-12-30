package bootstrap

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func Scrollspy(dataTarget, dataOffset string, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"data-spy": "scroll", "data-target": "#" + dataTarget, "data-offset": dataOffset,
		},
		Contents: contents,
	}
}
