package structdescription

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"reflect"
	"runtime"
	"strings"
)

func GetSourceFolderPath(p interface{}) string {
	t := reflect.TypeOf(p)
	_, filename, _, _ := runtime.Caller(0)
	kindPath := strings.Split(filename[len(build.Default.GOPATH):], "/")[1]
	pkg := t.PkgPath()
	out := build.Default.GOPATH + "/" + kindPath + "/" + pkg
	return out
}

type structFieldInfo struct {
	Name string
	Doc  string
	Tag  string
	Kind string
}
type StructFieldInfo interface {
	GetName() string
	GetDoc() string
	GetTag() string
	GetKind() string
	LoopkupTagKey(key string) (value string, ok bool)
}

func (sfi *structFieldInfo) GetName() string {
	return sfi.Name
}
func (sfi *structFieldInfo) GetDoc() string {
	return sfi.Doc
}
func (sfi *structFieldInfo) GetTag() string {
	return sfi.Tag
}
func (sfi *structFieldInfo) GetKind() string {
	return sfi.Kind
}
func (sfi *structFieldInfo) LoopkupTagKey(key string) (value string, ok bool) {
	return reflect.StructTag(strings.Trim(sfi.GetTag(), "`")).Lookup(key)
}

type structDescription struct {
	Name   string
	Fields []StructFieldInfo
	Doc    string
}
type StructDescription interface {
	GetName() string
	GetFields() []StructFieldInfo
	GetDoc() string
	GetFieldByName(name string) StructFieldInfo
}

func (sd *structDescription) GetName() string {
	return sd.Name
}
func (sd *structDescription) GetFields() []StructFieldInfo {
	return sd.Fields
}
func (sd *structDescription) GetFieldByName(name string) StructFieldInfo {
	var sfi StructFieldInfo
	for i := range sd.Fields {
		if sd.Fields[i].GetName() == name {
			sfi = sd.Fields[i]
			break
		}
	}
	return sfi
}
func (sd *structDescription) GetDoc() string {
	return sd.Doc
}

func castIfGenDecl(declRaw ast.Decl) *ast.GenDecl {
	switch declRaw.(type) {
	case *ast.GenDecl:
		return declRaw.(*ast.GenDecl)
	}
	return nil
}
func castIfTypeSpec(specRaw ast.Spec) *ast.TypeSpec {
	switch specRaw.(type) {
	case *ast.TypeSpec:
		return specRaw.(*ast.TypeSpec)
	}
	return nil
}
func GetKind(fieldType ast.Expr) (string, error) {
	kind := ""
	var typeInvestigate ast.Expr
	typeInvestigate = fieldType
	searchTypeRecursively := true
	for searchTypeRecursively {
		switch typeInvestigate.(type) {
		case *ast.Ident:
			kind += (typeInvestigate.(*ast.Ident)).Name
			searchTypeRecursively = false
			break
		case *ast.StarExpr:
			kind += "*"
			typeInvestigate = (typeInvestigate.(*ast.StarExpr)).X
		case *ast.SelectorExpr:
			// This one is incorrect, and deals with e.g. `strfmt.DateTime` fields
			xStr, err := GetKind((typeInvestigate.(*ast.SelectorExpr)).X)
			if err != nil {
				return "", err
			}
			selStr, err := GetKind((typeInvestigate.(*ast.SelectorExpr)).Sel)
			if err != nil {
				return "", err
			}
			return kind + xStr + "." + selStr, nil
		case *ast.ArrayType:
			kind += "[]"
			typeInvestigate = (typeInvestigate.(*ast.ArrayType)).Elt
		default:
			return "", fmt.Errorf("Unkown type given!")
		}
	}
	return kind, nil
}
func getFields(spec *ast.TypeSpec) []StructFieldInfo {
	fields := []StructFieldInfo{}
	structType := spec.Type.(*ast.StructType)
	for _, field := range structType.Fields.List {
		kind, err := GetKind(field.Type)
		if err != nil {
			continue
		}
		fields = append(fields, &structFieldInfo{
			Name: field.Names[0].Name,
			Tag:  field.Tag.Value,
			Kind: kind,
			Doc:  field.Doc.Text(),
		})
	}
	return fields
}
func getStructDescriptionWhenKind(ref interface{}, declRaw ast.Decl) StructDescription {
	decl := castIfGenDecl(declRaw)
	if decl == nil || decl.Tok.String() != "type" {
		return nil
	}
	fields := []StructFieldInfo{}
	var name *string
	found := false
	for _, specRaw := range decl.Specs {
		spec := castIfTypeSpec(specRaw)
		if spec == nil { //|| spec.Name.Name != reflect.TypeOf(ref).Name() {
			continue
		}
		name = &spec.Name.Name

		fields = getFields(spec)
		found = true
		break
	}
	if !found {
		return nil
	}
	return &structDescription{
		Name:   *name,
		Fields: fields,
		Doc:    decl.Doc.Text(),
	}
}

func GetStructDescription(ref interface{}) (StructDescription, error) {
	fset := token.NewFileSet() // positions are relative to fset
	// for ref != nil && reflect.ValueOf(ref).Kind() == reflect.Ptr {
	// 	ref = *(ref.(*interface{}))
	// }
	d, err := parser.ParseDir(fset, GetSourceFolderPath(ref), nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	refName := reflect.TypeOf(ref).Name()
	var sd StructDescription
	for _, f := range d {
		for _, f := range f.Files {
			for _, declRaw := range f.Decls {
				sdPntr := getStructDescriptionWhenKind(ref, declRaw)
				if sdPntr == nil || sdPntr.GetName() != refName {
					continue
				}
				sd = sdPntr
				break
			}
			if sd != nil {
				break
			}
		}
		if sd != nil {
			break
		}
	}
	return sd, nil
}
