package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func Badge(bgType BsTextBg, contents []htmlwrapper.Elm, roundedPill bool) htmlwrapper.Elm {
	roundedPillStr := ""
	if roundedPill {
		roundedPillStr = " " + BsRoundedPill
	}
	return &htmlwrapper.HTMLElm{
		Tag: "span",
		Attrs: map[string]string{
			"class": "badge " + string(bgType) + roundedPillStr,
		},
		Contents: contents,
	}
}
