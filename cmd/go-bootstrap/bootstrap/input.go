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

func Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool) htmlwrapper.Elm {
	plaintextStr := ""
	if plaintext {
		plaintextStr = " form-control-plaintext"
	}
	out := &htmlwrapper.HTMLElm{
		Tag: "input",
		Attrs: map[string]string{
			"type": string(kind) + string(size) + plaintextStr,
			"id":   id,
			"name": name,
		},
	}
	if checked {
		out.Attrs["checked"] = "checked"
	}
	if readonly {
		out.Attrs["readonly"] = "readonly"
	}
	if placeholder != nil {
		out.Attrs["placeholder"] = *placeholder
	}
	if value != nil {
		out.Attrs["value"] = *value
	}
	return out
}
