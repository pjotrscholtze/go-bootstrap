package structdescription

import (
	"go/ast"
	"go/token"
	"strings"
	"testing"
)

type Test struct {
	Field string
}

func TestGetSourceFolderPath(t *testing.T) {
	// GetSourceFolderPath(p interface{}) string
	path := GetSourceFolderPath("")
	if len(path) == 0 {
		t.Fatalf("No path was returned!")
	}
}

func TestGetSourceFolderPathStruct(t *testing.T) {
	// GetSourceFolderPath(p interface{}) string
	path := GetSourceFolderPath(Test{})
	if len(path) == 0 {
		t.Fatalf("No path was returned!")
	}
	if !strings.Contains(path, "structdescription") {
		t.Fatalf("The package was not found in the path name!")
	}
}

// func TestStructFieldInfoGetName(t *testing.T) {
// 	// StructFieldInfoGetName() string
// 	// structFieldInfo{}
// }
// func TestStructFieldInfoGetDoc(t *testing.T) {
// 	// StructFieldInfoGetDoc() string
// 	// structFieldInfo{}
// }
// func TestStructFieldInfoGetTag(t *testing.T) {
// 	// StructFieldInfoGetTag() string
// 	// structFieldInfo{}
// }
// func TestStructFieldInfoGetKind(t *testing.T) {
// 	// StructFieldInfoGetKind() string
// 	// structFieldInfo{}
// }

// func TestStructDescriptionGetName(t *testing.T) {
// 	// StructDescriptionGetName() string
// 	// structDescription{}
// }
// func TestStructDescriptionGetFields(t *testing.T) {
// 	// StructDescriptionGetFields() []StructFieldInfo
// 	// structDescription{}
// }
// func TestStructDescriptionGetFieldByName(t *testing.T) {
// 	// StructDescriptionGetFieldByName(name string) *StructFieldInfo
// 	// structDescription{}
// }
// func TestStructDescriptionGetDoc(t *testing.T) {
// 	// StructDescriptionGetDoc() string
// 	// structDescription{}
// }

func TestCastIfGenDecl(t *testing.T) {
	// castIfGenDecl(declRaw ast.Decl) *ast.GenDecl
	if castIfGenDecl(&ast.GenDecl{}) == nil {
		t.Fatalf("Expected GenDecl")
	}
}
func TestCastIfGenDeclNil(t *testing.T) {
	// castIfGenDecl(declRaw ast.Decl) *ast.GenDecl
	if castIfGenDecl(nil) != nil {
		t.Fatalf("Expected nil")
	}
}
func TestCastIfTypeSpec(t *testing.T) {
	// castIfTypeSpec(specRaw ast.Spec) *ast.TypeSpec
	if castIfTypeSpec(&ast.TypeSpec{}) == nil {
		t.Fatalf("Expected TypeSpec")
	}
}
func TestCastIfTypeSpecNil(t *testing.T) {
	// castIfTypeSpec(specRaw ast.Spec) *ast.TypeSpec
	if castIfTypeSpec(nil) != nil {
		t.Fatalf("Expected nil")
	}
}
func TestGetKind(t *testing.T) {
	// GetKind(fieldType ast.Expr) (string, error)
	kind, err := GetKind(&ast.Ident{Name: "test"})
	if err != nil {
		t.Fatalf("No error expected")
	}
	if kind != "test" {
		t.Fatalf("Expected `test`")
	}
}
func TestGetKindNil(t *testing.T) {
	// GetKind(fieldType ast.Expr) (string, error)
	kind, err := GetKind(nil)
	if err == nil {
		t.Fatalf("Error expected")
	}
	if kind != "" {
		t.Fatalf("Expected no kind to be given")
	}
}
func TestGetKindArray(t *testing.T) {
	// GetKind(fieldType ast.Expr) (string, error)
	kind, err := GetKind(&ast.ArrayType{Elt: &ast.Ident{Name: "test"}})
	if err != nil {
		t.Fatalf("No error expected")
	}
	if kind != "[]test" {
		t.Fatalf("Expected `test`")
	}
}
func TestGetKindStarExpr(t *testing.T) {
	// GetKind(fieldType ast.Expr) (string, error)
	kind, err := GetKind(&ast.StarExpr{X: &ast.Ident{Name: "test"}})
	if err != nil {
		t.Fatalf("No error expected")
	}
	if kind != "*test" {
		t.Fatalf("Expected `test`")
	}
}
func TestGetKindSelector(t *testing.T) {
	// GetKind(fieldType ast.Expr) (string, error)
	kind, err := GetKind(&ast.SelectorExpr{
		X:   &ast.Ident{Name: "x"},
		Sel: &ast.Ident{Name: "sel"},
	})
	if err != nil {
		t.Fatalf("No error expected")
	}
	if kind != "x.sel" {
		t.Fatalf("Expected `x.sel`")
	}
}
func TestGetFields(t *testing.T) {
	// getFields(spec *ast.TypeSpec) []StructFieldInfo
	res := getFields(&ast.TypeSpec{
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{{
					Type: &ast.Ident{
						Name: "test",
					},
					Names: []*ast.Ident{
						{Name: "fieldname"},
					},
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "comment"},
						}},
					Tag: &ast.BasicLit{
						Value: "tag",
					},
				}}}}})
	if len(res) != 1 {
		t.Fatalf("Expected result")
	}
	if res[0].GetDoc() != "comment\n" {
		t.Fatalf("Doc string should be `comment`")
	}
	if res[0].GetKind() != "test" {
		t.Fatalf("Kind should be `test`")
	}
	if res[0].GetName() != "fieldname" {
		t.Fatalf("Name should be `fieldname`")
	}
	if res[0].GetTag() != "tag" {
		t.Fatalf("Tag should be `tag`")
	}
}
func TestGetFields2(t *testing.T) {
	// getFields(spec *ast.TypeSpec) []StructFieldInfo
	res := getFields(&ast.TypeSpec{
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{{
					Type: &ast.Ident{
						Name: "test2",
					},
					Names: []*ast.Ident{
						{Name: "fieldname2"},
					},
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "comment2"},
						}},
					Tag: &ast.BasicLit{
						Value: "tag2",
					},
				}}}}})
	if len(res) != 1 {
		t.Fatalf("Expected result")
	}
	if res[0].GetDoc() != "comment2\n" {
		t.Fatalf("Doc string should be `comment`")
	}
	if res[0].GetKind() != "test2" {
		t.Fatalf("Kind should be `test`")
	}
	if res[0].GetName() != "fieldname2" {
		t.Fatalf("Name should be `fieldname`")
	}
	if res[0].GetTag() != "tag2" {
		t.Fatalf("Tag should be `tag`")
	}
}
func TestGetFieldsNofields(t *testing.T) {
	// getFields(spec *ast.TypeSpec) []StructFieldInfo
	res := getFields(&ast.TypeSpec{
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{}}}})
	if len(res) != 0 {
		t.Fatalf("Expected result")
	}

}

func TestGetStructDescriptionWhenKind(t *testing.T) {
	test := Test{
		Field: "field",
	}
	// getStructDescriptionWhenKind(ref interface{}, declRaw ast.Decl) StructDescription
	sd := getStructDescriptionWhenKind(test, &ast.GenDecl{
		Tok: token.TYPE, // @todo this is nicer than the string match in the implementation
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: &ast.Ident{
					Name: "typespec",
				},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{{
							Type: &ast.Ident{
								Name: "test2",
							},
							Names: []*ast.Ident{
								{Name: "fieldname2"},
							},
							Doc: &ast.CommentGroup{
								List: []*ast.Comment{
									{Text: "comment2"},
								}},
							Tag: &ast.BasicLit{
								Value: "tag2",
							},
						},
						},
					},
				},
			},
		},
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: "comment",
				},
			},
		},
	})
	if sd == nil {
		t.Fatalf("Expected result")
	}
}

func TestGetStructDescriptionWhenKind2(t *testing.T) {
	test := Test{
		Field: "field",
	}
	// getStructDescriptionWhenKind(ref interface{}, declRaw ast.Decl) StructDescription
	sd := getStructDescriptionWhenKind(test, &ast.GenDecl{
		Tok: token.TYPE, // @todo this is nicer than the string match in the implementation
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: &ast.Ident{
					Name: "typespec",
				},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{{
							Type: &ast.Ident{
								Name: "tes2t2",
							},
							Names: []*ast.Ident{
								{Name: "fie2ldname2"},
							},
							Doc: &ast.CommentGroup{
								List: []*ast.Comment{
									{Text: "comm2ent2"},
								}},
							Tag: &ast.BasicLit{
								Value: "ta2g2",
							},
						},
						},
					},
				},
			},
		},
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: "com2ment",
				},
			},
		},
	})
	if sd == nil {
		t.Fatalf("Expected result")
	}
}

// func TestGetStructDescription(t *testing.T) {
// 	// sd, err := GetStructDescription()
// 	// Paths do not make sense when running from test
// }
