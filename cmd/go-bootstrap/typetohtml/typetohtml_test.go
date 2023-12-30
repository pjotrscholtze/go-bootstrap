package typetohtml

import (
	"reflect"
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type MockStruct struct {
	NumberField int
	Mock        *MockStruct
}

func TestHasMapping(t *testing.T) {
	tmc := NewTypeMapperConv()
	// HasMapping(kind string) bool
	if !tmc.HasMapping("int8") {
		t.Fatalf("Was unable to find mapping for int8")
	}
	if !tmc.HasMapping("float32") {
		t.Fatalf("Was unable to find mapping for float32")
	}
	if !tmc.HasMapping("uint8") {
		t.Fatalf("Was unable to find mapping for uint8")
	}
	if tmc.HasMapping("asdfnasufdinsadifbsda") {
		t.Fatalf("Was able to find mapping for non existing type")
	}
}

func TestMapToHTML(t *testing.T) {
	tmc := NewTypeMapperConv()
	ms := MockStruct{
		NumberField: 1,
		Mock:        nil,
	}
	reflectedType := reflect.TypeOf(ms)
	rv := reflect.Indirect(reflect.ValueOf(&ms))
	f := reflectedType.Field(0)
	test := rv.FieldByName(f.Name)
	elm, err := tmc.MapToHTML(f, test, &ms)
	if err != nil {
		t.Fatalf("An error occured while trying to get the mapping of a field")
	}
	if elm == nil {
		t.Fatalf("No element find for mapping!")
	}
}
func TestMapToHTMLMockNil(t *testing.T) {
	tmc := NewTypeMapperConv()
	ms := MockStruct{
		NumberField: 1,
		Mock:        nil,
	}
	reflectedType := reflect.TypeOf(ms)
	rv := reflect.Indirect(reflect.ValueOf(&ms))
	f := reflectedType.Field(1)
	elm, err := tmc.MapToHTML(f, rv.FieldByName(f.Name), &ms)
	if err == nil {
		t.Fatalf("An error occured while trying to get the mapping of a field")
	}
	if elm != nil {
		t.Fatalf("Mapping for non existing field type!")
	}
}

func TestKnownTypes(t *testing.T) {
	tmc := NewTypeMapperConv()
	types := tmc.KnownTypes() // []stringConv()
	if len(types) == 0 {
		t.Fatalf("No default types were added")
	}
	if !containsString(types, "int") {
		t.Fatalf("int type was not found!")
	}
	if !containsString(types, "uint") {
		t.Fatalf("uint type was not found!")
	}
	if !containsString(types, "int32") {
		t.Fatalf("int32 type was not found!")
	}
	if !containsString(types, "uint32") {
		t.Fatalf("uint32 type was not found!")
	}
	if !containsString(types, "string") {
		t.Fatalf("string type was not found!")
	}
	if !containsString(types, "float64") {
		t.Fatalf("float64 type was not found!")
	}
}

func TestGetMapping(t *testing.T) {
	tmc := NewTypeMapperConv()
	fn, err := tmc.GetMapping("int") // (func(fieldValue reflect.Value) htmlwrapper.Elm, error)
	if err != nil {
		t.Fatalf("An error occured while getting a native datatype")
	}
	if fn == nil {
		t.Fatalf("No mapping function was returned!")
	}
	ms := MockStruct{
		NumberField: 1,
		Mock:        nil,
	}
	rv := reflect.Indirect(reflect.ValueOf(&ms))
	elm := fn(reflect.TypeOf(ms).Field(0), rv.FieldByName("NumberField"), nil)
	if elm == nil {
		t.Fatalf("Mapping function did not produce an element!")
	}
	html, err := elm.AsHTML()
	if err != nil {
		t.Fatalf("Something went wrong in the mapping function")
	}
	if len(html) == 0 {
		t.Fatalf("No HTML was produced")
	}
}
func TestGetMappingNonexesiting(t *testing.T) {
	tmc := NewTypeMapperConv()
	fn, err := tmc.GetMapping("asdfsdafadsfint") // (func(fieldValue reflect.Value) htmlwrapper.Elm, error)
	if err == nil {
		t.Fatalf("No error occured while getting a non-existing datatype")
	}
	if fn != nil {
		t.Fatalf("A mapping function was returned for a non-existing datatype!")
	}
}

func TestGetMappingRegistered(t *testing.T) {
	tmc := NewTypeMapperConv()
	fn, err := tmc.GetMapping("MockStruct") // (func(fieldValue reflect.Value) htmlwrapper.Elm, error)
	if err == nil {
		t.Fatalf("No error occured while getting a non default datatype")
	}
	if fn != nil {
		t.Fatalf("Mapping function was returned for a non default type!")
	}

	tmc.RegisterMapping("MockStruct", func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: "content"}
	})

	fn, err = tmc.GetMapping("MockStruct") // (func(fieldValue reflect.Value) htmlwrapper.Elm, error)
	if err != nil {
		t.Fatalf("An error occured while getting a newly registered datatype")
	}
	if fn == nil {
		t.Fatalf("No mapping function was returned after registering!")
	}

	ms := MockStruct{
		NumberField: 1,
		Mock:        nil,
	}
	rv := reflect.Indirect(reflect.ValueOf(&ms))
	elm := fn(reflect.TypeOf(ms).Field(1), rv.FieldByName("Mock"), nil)
	if elm == nil {
		t.Fatalf("Mapping function did not produce an element!")
	}
	html, err := elm.AsHTML()
	if err != nil {
		t.Fatalf("Something went wrong in the mapping function")
	}
	if len(html) == 0 {
		t.Fatalf("No HTML was produced")
	}
}

// Test a mapping function with a newly addded function

func TestRegisterMapping(t *testing.T) {
	tmc := NewTypeMapperConv()
	knownTypesLenBefore := len(tmc.KnownTypes())
	tmc.RegisterMapping("forKind", func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: "Test"}
	})
	knownTypesLenAfter := len(tmc.KnownTypes())
	if knownTypesLenBefore+1 != knownTypesLenAfter {
		t.Fatalf("Regsitering a type didnt increase the known types")
	}
	tmc.RegisterMapping("forKind", func(structField reflect.StructField, value reflect.Value, refStruct interface{}) htmlwrapper.Elm {
		return &htmlwrapper.TextElm{Content: "Test2"}
	})
	knownTypesLenAfterSecond := len(tmc.KnownTypes())
	if knownTypesLenAfter != knownTypesLenAfterSecond {
		t.Fatalf("Reregistering the type mapper increases the amount of types")
	}
}

func TestNewTypeMapperConv(t *testing.T) {
	if NewTypeMapperConv() == nil {
		t.Fatalf("NewTypeMapperConv returned nil, this should never occur")
	}
}

func TestNewTypeMapperConvShouldHaveDefaultMapping(t *testing.T) {
	if len(NewTypeMapperConv().KnownTypes()) == 0 {
		t.Fatalf("TypeMapperConv from NewTypeMapperConv should always have known types")
	}
}
