package typetohtml

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type typeMapperConv struct {
	mappings map[string]func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm
}

func (tmc *typeMapperConv) HasMapping(kind string) bool {
	_, ok := tmc.mappings[kind]
	return ok
}

// (kind string, fieldValue reflect.Value) (htmlwrapper.Elm, error)
func (tmc *typeMapperConv) MapToHTML(structField reflect.StructField, value reflect.Value, refStruct interface{}) (htmlwrapper.Elm, error) {
	if !tmc.HasMapping(structField.Type.String()) {
		return nil, fmt.Errorf("Unkown field value given! No mapping avaiable!")
	}

	html := tmc.mappings[structField.Type.String()](structField, value, refStruct)
	return html, nil
}

type TypeMapperConv interface {
	KnownTypes() []string
	HasMapping(kind string) bool
	// MapToHTML(kind string, fieldValue reflect.Value) (htmlwrapper.Elm, error)
	MapToHTML(structField reflect.StructField, value reflect.Value, refStruct interface{}) (htmlwrapper.Elm, error)
	RegisterMapping(forKind string, fn func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm)
	GetMapping(kind string) (func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm, error)
}

func (tmc *typeMapperConv) KnownTypes() []string {
	out := []string{}
	for key := range tmc.mappings {
		out = append(out, key)
	}
	return out
}

func (tmc *typeMapperConv) GetMapping(kind string) (func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm, error) {
	val, ok := tmc.mappings[kind]
	if !ok {
		return nil, fmt.Errorf("Failed to find type mapping!")
	}
	return val, nil
}
func (tmc *typeMapperConv) RegisterMapping(forKind string, fn func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm) {
	tmc.mappings[forKind] = fn
}

func NewTypeMapperConv() TypeMapperConv {

	// html := tmc.mappings[structField.Type.String()](structField reflect.StructField, value reflect.Value)
	// html := tmc.mappings[structField.Type.String()](value.FieldByName(structField.Name))

	out := &typeMapperConv{mappings: make(map[string]func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm)}
	out.mappings["int64"] = func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: strconv.FormatInt(value.Int(), 10)}
	}
	out.mappings["int32"] = out.mappings["int64"]
	out.mappings["int16"] = out.mappings["int64"]
	out.mappings["int8"] = out.mappings["int64"]
	out.mappings["int"] = out.mappings["int64"]

	out.mappings["uint64"] = func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: strconv.FormatUint(value.Uint(), 10)}
	}
	out.mappings["uint32"] = out.mappings["uint64"]
	out.mappings["uint16"] = out.mappings["uint64"]
	out.mappings["uint8"] = out.mappings["uint64"]
	out.mappings["uint"] = out.mappings["uint64"]

	out.mappings["bool"] = func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: strconv.FormatBool(value.Bool())}
	}

	out.mappings["float32"] = func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: strconv.FormatFloat(value.Float(), 'f', 2, 32)}
	}
	out.mappings["float64"] = out.mappings["float32"]

	out.mappings["string"] = func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: value.String()}
	}

	return out
}
