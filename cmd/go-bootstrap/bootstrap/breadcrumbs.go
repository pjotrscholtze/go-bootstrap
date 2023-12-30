package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

type BreadcrumbEntry struct {
	Href string
	Text string
}

func Breadcrumb(items []BreadcrumbEntry) htmlwrapper.Elm {
	contents := make([]htmlwrapper.Elm, len(items))
	for i, item := range items {
		var linkItem htmlwrapper.Elm
		linkItem = &htmlwrapper.HTMLElm{
			Tag: "a",
			Attrs: map[string]string{
				"href": item.Href,
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.TextElm{
					Content: item.Text,
				},
			},
		}
		activeStr := ""
		if len(items)-1 == i {
			activeStr = " active"
			linkItem = &htmlwrapper.TextElm{Content: item.Text}
		}
		contents[i] = &htmlwrapper.HTMLElm{
			Tag: "li",
			Attrs: map[string]string{
				"class":        "breadcrumb-item" + activeStr,
				"aria-current": "page",
			},
			Contents: []htmlwrapper.Elm{
				linkItem,
			},
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag: "nav",
		Attrs: map[string]string{
			"aria-label": "breadcrumb",
		},
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "ol",
				Attrs: map[string]string{
					"class": "breadcrumb",
				},
				Contents: contents,
			},
		},
	}
}
