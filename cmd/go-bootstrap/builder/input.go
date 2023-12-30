package builder

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type InputBuilder interface {
	AsHTML() (string, error)
	AsElm() htmlwrapper.Elm
	SetKind(kind bootstrap.BsInputType) InputBuilder
	SetId(id string) InputBuilder
	SetName(name string) InputBuilder
	SetPlaceholder(placeholder *string) InputBuilder
	SetValue(value *string) InputBuilder
	SetSize(size bootstrap.BsInputSize) InputBuilder
	SetReadonly(readonly bool) InputBuilder
	SetPlaintext(plaintext bool) InputBuilder
	SetChecked(checked bool) InputBuilder
	SetRequired(required bool) InputBuilder
	SetLabel(label *string) InputBuilder
}

type inputBuilder struct {
	kind        bootstrap.BsInputType
	id          string
	name        string
	placeholder *string
	value       *string
	size        bootstrap.BsInputSize
	readonly    bool
	plaintext   bool
	checked     bool
	required    bool
	label       *string
}

func (ib *inputBuilder) AsElm() htmlwrapper.Elm {
	elm := bootstrap.Input(ib.kind, ib.id, ib.name, ib.placeholder, ib.value,
		ib.size, ib.readonly, ib.plaintext, ib.checked).(*htmlwrapper.HTMLElm)
	if ib.required {
		elm.Attrs["required"] = "required"
	}
	content := []htmlwrapper.Elm{}
	if ib.label != nil {
		content = append(content, &htmlwrapper.HTMLElm{
			Tag: "label",
			Attrs: map[string]string{
				"for":   ib.id,
				"class": "form-label",
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.TextElm{
					Content: *ib.label,
				},
			},
		})
	}
	content = append(content, elm)
	return (&htmlwrapper.MultiElm{
		Contents: content,
	})
}
func (ib *inputBuilder) AsHTML() (string, error) {
	return ib.AsElm().AsHTML()
}

func (ib *inputBuilder) SetKind(kind bootstrap.BsInputType) InputBuilder {
	ib.kind = kind
	return ib
}
func (ib *inputBuilder) SetId(id string) InputBuilder {
	ib.id = id
	return ib
}
func (ib *inputBuilder) SetName(name string) InputBuilder {
	ib.name = name
	return ib
}
func (ib *inputBuilder) SetPlaceholder(placeholder *string) InputBuilder {
	ib.placeholder = placeholder
	return ib
}
func (ib *inputBuilder) SetValue(value *string) InputBuilder {
	ib.value = value
	return ib
}
func (ib *inputBuilder) SetSize(size bootstrap.BsInputSize) InputBuilder {
	ib.size = size
	return ib
}
func (ib *inputBuilder) SetReadonly(readonly bool) InputBuilder {
	ib.readonly = readonly
	return ib
}
func (ib *inputBuilder) SetPlaintext(plaintext bool) InputBuilder {
	ib.plaintext = plaintext
	return ib
}
func (ib *inputBuilder) SetChecked(checked bool) InputBuilder {
	ib.checked = checked
	return ib
}
func (ib *inputBuilder) SetRequired(required bool) InputBuilder {
	ib.required = required
	return ib
}
func (ib *inputBuilder) SetLabel(label *string) InputBuilder {
	ib.label = label
	return ib
}

func NewInput() InputBuilder {
	return &inputBuilder{
		kind:      bootstrap.BsInputTypeText,
		id:        "to-fill-in",
		name:      "to-fill-in",
		readonly:  false,
		plaintext: false,
		checked:   false,
		size:      bootstrap.BsInputSizeNormal,
		required:  false,
	}
}
