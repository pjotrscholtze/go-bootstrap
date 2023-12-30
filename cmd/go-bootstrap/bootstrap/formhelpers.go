package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func FormHelper(id string, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "small",
		Attrs: map[string]string{
			"class": "form-text text-muted",
			"id":    id,
		},
		Contents: contents,
	}
}

func FormLabel(forId string, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "small",
		Attrs: map[string]string{
			"class": "form-text text-muted",
			"for":   forId,
		},
		Contents: contents,
	}
}

func FormInvalidFeedback(contents []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "invalid-feedback",
		},
		Contents: contents,
	}
}

func FormTooltip(valid bool, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	classStr := "invalid-tooltip"
	if valid {
		classStr = "valid-tooltip"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": classStr,
		},
		Contents: contents,
	}
}
