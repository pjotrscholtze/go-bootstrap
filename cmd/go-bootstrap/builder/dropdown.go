package builder

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type selectItem struct {
	value    string
	content  string
	disabled bool
}

type SelectBuilder interface {
	AsHTML() (string, error)
	AsElm() htmlwrapper.Elm
	AddEntry(value, content string, disabled bool) SelectBuilder
	SetLabel(label *string) SelectBuilder
	SetSelectedItem(selectedItem *string) SelectBuilder
}

type selectBuilder struct {
	items        []selectItem
	id           string
	required     bool
	label        *string
	selectedItem *string
}

func (sb *selectBuilder) SetSelectedItem(selectedItem *string) SelectBuilder {
	sb.selectedItem = selectedItem
	return sb
}

func (sb *selectBuilder) SetLabel(label *string) SelectBuilder {
	sb.label = label
	return sb
}
func (sb *selectBuilder) AddEntry(value, content string, disabled bool) SelectBuilder {
	sb.items = append(sb.items, selectItem{
		value:    value,
		content:  content,
		disabled: disabled,
	})
	return sb
}

func (sb *selectBuilder) AsElm() htmlwrapper.Elm {
	out := &htmlwrapper.HTMLElm{
		Tag: "select",
		Attrs: map[string]string{
			"class": "form-select",
			"id":    sb.id,
		},
		Contents: []htmlwrapper.Elm{},
	}
	if sb.required {
		out.Attrs["required"] = "required"
	}
	for _, si := range sb.items {
		entry := &htmlwrapper.HTMLElm{
			Tag: "option",
			Attrs: map[string]string{
				"value": si.value,
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.TextElm{Content: si.content},
			},
		}
		if sb.selectedItem != nil && si.value == *sb.selectedItem {
			entry.Attrs["selected"] = "selected"
		}
		if si.disabled {
			entry.Attrs["disabled"] = "disabled"
		}
		out.Contents = append(out.Contents, entry)
	}
	content := []htmlwrapper.Elm{}
	if sb.label != nil {
		content = append(content, &htmlwrapper.HTMLElm{
			Tag: "label",
			Attrs: map[string]string{
				"for":   sb.id,
				"class": "form-label",
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.TextElm{
					Content: *sb.label,
				},
			},
		})
	}
	content = append(content, out)
	return &htmlwrapper.MultiElm{Contents: content}
}
func (sb *selectBuilder) AsHTML() (string, error) {
	return sb.AsElm().AsHTML()
}

func NewSelect() SelectBuilder {
	return &selectBuilder{
		items: []selectItem{},
		id:    "",
	}
}
