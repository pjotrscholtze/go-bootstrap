package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func CardImg(src string, alt *string) htmlwrapper.Elm {
	attrs := map[string]string{
		"class": "card-img-top",
		"src":   src,
	}
	if alt != nil {
		attrs["alt"] = *alt
	}
	return &htmlwrapper.HTMLElm{
		Tag:   "img",
		Attrs: attrs,
	}
}

func CardText(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "p",
		Attrs: map[string]string{
			"class": "card-text",
		},
		Contents: content,
	}
}

func CardFooter(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "card-footer",
		},
		Contents: content,
	}
}

func CardBody(title *string, content []htmlwrapper.Elm) htmlwrapper.Elm {
	bodyContent := content
	if title != nil {
		bodyContent = append(content, &htmlwrapper.HTMLElm{
			Tag: "h5",
			Attrs: map[string]string{
				"class": "card-title",
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.TextElm{
					Content: *title,
				},
			},
		})
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "card-body",
		},
		Contents: bodyContent,
	}
}
func Card(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "card",
		},
		Contents: content,
	}
}

func CardImgOverlay(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "card-img-overlay",
		},
		Contents: content,
	}
}
