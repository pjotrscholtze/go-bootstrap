package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

type ListGroupItem struct {
	Content        htmlwrapper.Elm
	Active         bool
	Disabled       bool
	JustifyContent bool
	Href           *string
	Kind           *ListGroupItemKind
}

func Listgroup(content []ListGroupItem, flush bool) htmlwrapper.Elm {
	listgroupContent := make([]htmlwrapper.Elm, len(content))
	for i, item := range content {
		itemClass := "list-group-item"
		tag := "li"
		if item.Href != nil {
			tag = "a"
			itemClass += " list-group-item-action"
		}
		if item.Active {
			itemClass += " active"
		}
		if item.Disabled {
			itemClass += " disabled"
		}
		if item.JustifyContent {
			itemClass += " justify-content-between align-items-center"
		}
		if item.Kind != nil {
			itemClass += " " + string(*item.Kind)
		}
		listgroupContent[i] = &htmlwrapper.HTMLElm{
			Tag: tag,
			Attrs: map[string]string{
				"class": itemClass,
			},
			Contents: []htmlwrapper.Elm{item.Content},
		}
	}
	classStr := "list-group"
	if flush {
		classStr += " list-group-flush"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "ul",
		Attrs: map[string]string{
			"class": classStr,
		},
		Contents: listgroupContent,
	}
}
