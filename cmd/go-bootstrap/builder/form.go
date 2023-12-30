package builder

import (
	"reflect"
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/typetohtml"
)

type HTTPMethod string

const (
	HTTPMethodGet  HTTPMethod = "get"
	HTTPMethodPost            = "post"
)

type formBuilder[T interface{}] struct {
	ref            T
	method         HTTPMethod
	action         string
	typeMapperConv typetohtml.TypeMapperConv
}
type FormBuilder interface {
	ActionMethod(method HTTPMethod, action string) FormBuilder
	GetTypeMapperConv() typetohtml.TypeMapperConv

	RegisterMappingMulti(mapping map[string]func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm)
	SetDefaultMappings() FormBuilder
	AsElm() htmlwrapper.Elm
}

func (fb *formBuilder[T]) RegisterMappingMulti(mapping map[string]func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm) {
	for k, v := range mapping {
		fb.GetTypeMapperConv().RegisterMapping(k, v)
	}
}

func (fb *formBuilder[T]) GetTypeMapperConv() typetohtml.TypeMapperConv {
	return fb.typeMapperConv
}

func (fb *formBuilder[T]) ActionMethod(method HTTPMethod, action string) FormBuilder {
	fb.method = method
	fb.action = action
	return fb
}
func (fb *formBuilder[T]) getFormContent() ([]htmlwrapper.Elm, error) {
	t := reflect.TypeOf(fb.ref)
	rv := reflect.Indirect(reflect.ValueOf(&fb.ref))
	formContent := []htmlwrapper.Elm{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if fb.typeMapperConv.HasMapping(f.Type.String()) {
			elm, err := fb.typeMapperConv.MapToHTML(f, rv.FieldByName(f.Name), fb.ref)
			if err != nil {
				return nil, err
			}
			formContent = append(formContent, elm)
		}
	}
	return formContent, nil
}

func (fb *formBuilder[T]) AsElm() htmlwrapper.Elm {
	content, _ := fb.getFormContent()
	return &htmlwrapper.HTMLElm{
		Tag: "form",
		Attrs: map[string]string{
			"class":  "container",
			"action": fb.action,
			"method": string(fb.method),
		},
		Contents: content, //[]htmlwrapper.Elm{},
	}
}
func (fb *formBuilder[T]) SetDefaultMappings() FormBuilder {
	prependStr, appendStr := "", ""
	fb.RegisterMappingMulti(map[string]func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm{
		"string": func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
			valueStr, nameId := value.String(), ""
			if len(prependStr) > 0 {
				nameId = prependStr + "-"
			}
			nameId += structField.Name
			if len(appendStr) > 0 {
				nameId += "-" + appendStr
			}

			return NewInput().
				SetName(nameId).
				SetId(nameId).
				SetLabel(&structField.Name).
				SetValue(&valueStr).
				AsElm()
		},
		"int64": func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
			valueStr := strconv.FormatInt(value.Int(), 10)
			nameId := ""
			if len(prependStr) > 0 {
				nameId = prependStr + "-"
			}
			nameId += structField.Name
			if len(appendStr) > 0 {
				nameId += "-" + appendStr
			}
			return bootstrap.InputGroup([]htmlwrapper.Elm{
				bootstrap.InputGroupText([]htmlwrapper.Elm{
					&htmlwrapper.TextElm{Content: structField.Name},
				}, nameId+"-label"),
				bootstrap.Input(bootstrap.BsInputTypeNumber, nameId, nameId, nil, &valueStr, bootstrap.BsInputSizeNormal, false, false, false),
			})
		},
	})
	return fb
}

func NewFormBuilder[T interface{}](ref T) FormBuilder {
	return &formBuilder[T]{
		ref:            ref,
		typeMapperConv: typetohtml.NewTypeMapperConv(),
	}
}
