package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func AccordionItem(heading, id, parentId string, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "accordion-item",
		},
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "h2",
				Attrs: map[string]string{
					"class": "accordion-header",
					"id":    "heading-" + id,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "button",
						Attrs: map[string]string{
							"class":          "accordion-button",
							"id":             id,
							"type":           "button",
							"data-bs-toggle": "collapse",
							"data-bs-target": "#" + id,
							"aria-expanded":  "true",
							"aria-controls":  id,
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.TextElm{Content: heading},
						},
					},
				},
			},
			&htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"id":              id,
					"class":           "accordion-collapse collapse show",
					"aria-labelledby": "heading-" + id,
					"data-bs-parent":  "#" + parentId,
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "div",
						Attrs: map[string]string{
							"class": "accordion-body",
						},
						Contents: contents,
					},
				},
			},
		},
	}
}
func Accordion(accodionFlush bool, id string, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	addClass := ""
	if accodionFlush {
		addClass = " accordion-flush"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "accordion" + addClass,
			"id":    id,
		},
		Contents: contents,
	}
}
