package bootstrap

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func ModalHeader(headerText string, closeBtn bool) htmlwrapper.Elm {
	headerContent := []htmlwrapper.Elm{
		&htmlwrapper.HTMLElm{
			Tag: "h5",
			Attrs: map[string]string{
				"class": "modal-title",
			},
			Contents: []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: headerText}},
		}}
	if closeBtn {
		headerContent = append(headerContent,
			&htmlwrapper.HTMLElm{
				Tag: "button",
				Attrs: map[string]string{
					"type": "button", "class": "close", "data-dismiss": "modal", "aria-label": "Close",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "span",
						Attrs: map[string]string{
							"aria-hidden": "true",
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.TextElm{
								Content: "&times;",
							},
						},
					},
				},
			})
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "modal-header",
		},
		Contents: headerContent,
	}
}

func ModalBody(content []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "modal-body",
		},
		Contents: content,
	}
}
func ModalFooter(content []htmlwrapper.Elm, closeBtn *string) htmlwrapper.Elm {
	footerContent := content
	if closeBtn != nil {
		footerContent = append(content, &htmlwrapper.HTMLElm{
			Tag: "button",
			Attrs: map[string]string{
				"type": "button", "class": "btn btn-secondary", "data-dismiss": "modal",
			},
			Contents: []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: *closeBtn}},
		})
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "modal-footer",
		},
		Contents: footerContent,
	}
}

func Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm) htmlwrapper.Elm {
	addClass := string(modelSize)
	if modalDialogCentered {
		addClass += " modal-dialog-centered"
	}
	if fade {
		addClass += " fade"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "modal", "tabindex": "-1", "role": "dialog",
		},
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"class": "modal-dialog" + addClass, "role": "document",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "div",
						Attrs: map[string]string{
							"class": "modal-content",
						},
						Contents: modalContent,
					},
				},
			},
		},
	}
}

func ModelWithButton(modelSize ModelSize, fade bool, id string, modalDialogCentered bool, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, modalContent []htmlwrapper.Elm) htmlwrapper.Elm {
	btn := Button(btnType, btnKind, btnText, btnSize, btnState, nil).(*htmlwrapper.HTMLElm)
	btn.Attrs["data-toggle"] = "modal"
	btn.Attrs["data-target"] = "#" + id
	model := Modal(modelSize, fade, modalDialogCentered, modalContent).(*htmlwrapper.HTMLElm)
	model.Attrs["id"] = id
	model.Attrs["aria-hidden"] = "true"
	return &htmlwrapper.MultiElm{Contents: []htmlwrapper.Elm{model, btn}}
}
