package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func InputGroup(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "input-group",
		},
		Contents: content,
	}
}

func InputGroupPrepend(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "input-group-prepend",
		},
		Contents: content,
	}
}

func InputGroupAppend(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "input-group-append",
		},
		Contents: content,
	}
}

func InputGroupText(content []htmlwrapper.Elm, id string) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "input-group-text",
			"id":    id,
		},
		Contents: content,
	}
}

func Input(kind BsInputType, id, name string, descriptiveText, value *string, size BsInputSize, readonly, plaintext, checked bool) htmlwrapper.Elm {
	plaintextStr := ""
	if plaintext {
		plaintextStr = " form-control-plaintext"
	}
	formControlClass := "form-control"
	if kind == BsInputTypeCheckbox {
		formControlClass = "form-check-input"
	}
	out := &htmlwrapper.HTMLElm{
		Tag: "input",
		Attrs: map[string]string{
			"type":  string(kind),
			"class": formControlClass + " " + string(size) + plaintextStr,
			"id":    id,
			"name":  name,
		},
	}
	if checked {
		out.Attrs["checked"] = "checked"
	}
	if readonly {
		out.Attrs["readonly"] = "readonly"
	}
	if value != nil {
		out.Attrs["value"] = *value
	}
	if kind == BsInputTypeCheckbox {
		out = &htmlwrapper.HTMLElm{
			Tag: "div",
			Attrs: map[string]string{
				"class": "form-check",
			},
			Contents: []htmlwrapper.Elm{out},
		}
		if descriptiveText != nil {
			out.Contents = append(out.Contents, &htmlwrapper.HTMLElm{
				Tag: "label",
				Attrs: map[string]string{
					"class": "form-check-label",
					"for":   id,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.TextElm{Content: *descriptiveText},
				},
			})
		}
	} else {
		if descriptiveText != nil {
			out.Attrs["placeholder"] = *descriptiveText
		}
	}
	return out
}
