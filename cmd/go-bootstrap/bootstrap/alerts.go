package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

type AlertType string

const (
	AlertPrimary   AlertType = "alert-primary"
	AlertSecondary           = "alert-secondary"
	AlertSuccess             = "alert-success"
	AlertDanger              = "alert-danger"
	AlertWarning             = "alert-warning"
	AlertInfo                = "alert-info"
	AlertLight               = "alert-light"
	AlertDark                = "alert-dark"
)
const (
	AlertDismissible  string = "alert-dismissible"
	AlertLinkClass           = "alert-link"
	AlertHeadingClass        = "alert-heading"
)

func Alert(alertType AlertType, contents []htmlwrapper.Elm, dismissible bool) htmlwrapper.Elm {
	dismissibleStr := ""
	alertContents := contents
	if dismissible {
		dismissibleStr = " " + AlertDismissible
		alertContents = append(alertContents, &htmlwrapper.HTMLElm{
			Tag: "button",
			Attrs: map[string]string{
				"type":            "button",
				"class":           "btn-close",
				"data-bs-dismiss": "alert",
				"aria-label":      "Close",
			},
		})
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "alert " + string(alertType) + dismissibleStr,
			"role":  "alert",
		},
		Contents: alertContents,
	}
}
