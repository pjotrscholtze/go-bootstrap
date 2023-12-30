package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func Col(contents []htmlwrapper.Elm, size BsColSize, formGroup bool) htmlwrapper.Elm {
	colSizeStr := string(size)
	if size != BsColSizeAuto {
		size = " " + size
	}
	formGroupStr := ""
	if formGroup {
		formGroupStr = " form-group"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "col" + colSizeStr + formGroupStr,
		},
		Contents: contents,
	}
}

func FormRow(contents []htmlwrapper.Elm, formGroup bool) htmlwrapper.Elm {
	fg := ""
	if formGroup {
		fg = " form-group"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "form-row" + fg,
		},
		Contents: contents,
	}
}
