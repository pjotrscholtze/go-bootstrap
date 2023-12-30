package builder

import (
	"reflect"
	"slices"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/typetohtml"
)

type tableMappingEntry struct {
	FieldName string
	LabelName string
}

type tableMapping struct {
	entries []tableMappingEntry
}
type TableMapping interface {
	Set(fieldName string, labelName string) TableMapping
	RemoveByFieldName(fieldName string) TableMapping
	MoveToIndex(fieldName string, idx uint) TableMapping
	OnlyInclude(fieldNames []string) TableMapping
	GetEntries() []tableMappingEntry
}

func (tm *tableMapping) GetEntries() []tableMappingEntry {
	return tm.entries
}
func (tm *tableMapping) Set(fieldName string, labelName string) TableMapping {
	for i := range tm.entries {
		if tm.entries[i].FieldName == fieldName {
			tm.entries[i].LabelName = labelName
			return tm
		}
	}
	tm.entries = append(tm.entries, tableMappingEntry{
		FieldName: fieldName,
		LabelName: labelName,
	})
	return tm
}
func (tm *tableMapping) RemoveByFieldName(fieldName string) TableMapping {
	for i := range tm.entries {
		if tm.entries[i].FieldName == fieldName {
			tm.entries = append(tm.entries[:i], tm.entries[i+1:]...)
			break
		}
	}
	return tm
}
func (tm *tableMapping) MoveToIndex(fieldName string, idx uint) TableMapping {
	if idx >= uint(len(tm.entries)) {
		return tm
	}
	labelName := fieldName
	idxToRemove := -1
	for i, v := range tm.entries {
		if v.FieldName == fieldName {
			idxToRemove = i
			labelName = v.LabelName
			break
		}
	}
	if idxToRemove != -1 {
		tm.RemoveByFieldName(fieldName)

		// func insert(a []int, index int, value int) []int {
		if uint(len(tm.entries)) == idx { // nil or empty slice or after last element
			tm.entries = append(tm.entries, tableMappingEntry{
				FieldName: fieldName,
				LabelName: labelName,
			})
			return tm
		}
		tm.entries = append(tm.entries[:idx+1], tm.entries[idx:]...) // idx < len(a)
		tm.entries[idx] = tableMappingEntry{
			FieldName: fieldName,
			LabelName: labelName,
		}

	} else {
		tm.entries = append(append(tm.entries[:idx], tableMappingEntry{
			FieldName: fieldName,
			LabelName: labelName,
		}), tm.entries[idx+1:]...)

	}

	return tm
}
func (tm *tableMapping) OnlyInclude(fieldNames []string) TableMapping {
	toRemove := []string{}
	for _, e := range tm.entries {
		if !slices.Contains[[]string, string](fieldNames, e.FieldName) {
			toRemove = append(toRemove, e.FieldName)
		}
	}
	for _, remove := range toRemove {
		tm.RemoveByFieldName(remove)
	}
	return tm
}

type tableBuilder[T interface{}] struct {
	ref            []T
	typeMapperConv typetohtml.TypeMapperConv
	striped        bool
	hoverable      bool
	color          bootstrap.BsTableColor
	border         bootstrap.BsTableBorderColor
	size           bootstrap.BsTableSize
	footer         *[]htmlwrapper.Elm
	caption        *bootstrap.TableCaption
	tableMapping   *tableMapping
}

type TableBuilder interface {
	AsElm() htmlwrapper.Elm

	GetTableMapping() TableMapping

	SetStriped(striped bool) TableBuilder
	SetHoverable(hoverable bool) TableBuilder
	SetColor(color bootstrap.BsTableColor) TableBuilder
	SetBorder(border bootstrap.BsTableBorderColor) TableBuilder
	SetSize(size bootstrap.BsTableSize) TableBuilder
	SetFooter(footer *[]htmlwrapper.Elm) TableBuilder
	SetCaption(caption *bootstrap.TableCaption) TableBuilder
}

func (tb *tableBuilder[T]) GetTableMapping() TableMapping {
	if tb.tableMapping == nil {
		entries := []tableMappingEntry{}
		if len(tb.ref) > 0 {
			t := reflect.TypeOf(tb.ref[0])
			for i := 0; i < t.NumField(); i++ {
				entries = append(entries, tableMappingEntry{
					FieldName: t.Field(i).Name,
					LabelName: t.Field(i).Name,
				})
			}
		}

		tb.tableMapping = &tableMapping{
			entries: entries,
		}
	}
	return tb.tableMapping
}

func (tb *tableBuilder[T]) SetStriped(striped bool) TableBuilder {
	tb.striped = striped
	return tb
}
func (tb *tableBuilder[T]) SetHoverable(hoverable bool) TableBuilder {
	tb.hoverable = hoverable
	return tb
}
func (tb *tableBuilder[T]) SetColor(color bootstrap.BsTableColor) TableBuilder {
	tb.color = color
	return tb
}
func (tb *tableBuilder[T]) SetBorder(border bootstrap.BsTableBorderColor) TableBuilder {
	tb.border = border
	return tb
}
func (tb *tableBuilder[T]) SetSize(size bootstrap.BsTableSize) TableBuilder {
	tb.size = size
	return tb
}
func (tb *tableBuilder[T]) SetFooter(footer *[]htmlwrapper.Elm) TableBuilder {
	tb.footer = footer
	return tb
}
func (tb *tableBuilder[T]) SetCaption(caption *bootstrap.TableCaption) TableBuilder {
	tb.caption = caption
	return tb
}
func (tb *tableBuilder[T]) getRowContent(ref T) ([]htmlwrapper.Elm, error) {
	t := reflect.TypeOf(ref)
	rv := reflect.Indirect(reflect.ValueOf(&ref))
	content := []htmlwrapper.Elm{}
	for _, entry := range tb.GetTableMapping().GetEntries() {
		f, ok := t.FieldByName(entry.FieldName)
		if ok && tb.typeMapperConv.HasMapping(f.Type.String()) {
			elm, err := tb.typeMapperConv.MapToHTML(f, rv.FieldByName(f.Name), ref)
			if err != nil {
				return nil, err
			}
			content = append(content, bootstrap.TableCell(
				false,
				bootstrap.BsTableCellKindNormal,
				1,
				bootstrap.BsTableColorDefault,
				[]htmlwrapper.Elm{elm},
			))
		} else {
			content = append(content, bootstrap.TableCell(
				false,
				bootstrap.BsTableCellKindNormal,
				1,
				bootstrap.BsTableColorDefault,
				[]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: ""}},
			))
		}

	}

	return content, nil

}

func (tb *tableBuilder[T]) getHeading() []htmlwrapper.Elm {
	out := []htmlwrapper.Elm{}
	for _, entry := range tb.GetTableMapping().GetEntries() {
		out = append(out, bootstrap.TableCell(
			false,
			bootstrap.BsTableCellKindNormal,
			1,
			bootstrap.BsTableColorDark,
			[]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: entry.LabelName}},
		))
	}
	return out
}

func (tb *tableBuilder[T]) AsElm() htmlwrapper.Elm {
	body := []htmlwrapper.Elm{}
	for _, row := range tb.ref {
		rowElm, _ := tb.getRowContent(row)
		body = append(body, bootstrap.TableRow(false, bootstrap.BsTableColorDefault, rowElm))
	}
	heading := tb.getHeading()
	return bootstrap.Table(tb.striped, tb.hoverable, tb.color, tb.border,
		tb.size, heading, body, tb.footer, tb.caption)
}

func NewTableBuilder[T interface{}](ref []T) TableBuilder {
	return &tableBuilder[T]{
		ref:            ref,
		typeMapperConv: typetohtml.NewTypeMapperConv(),

		striped:   true,
		hoverable: true,
		color:     bootstrap.BsTableColorDefault,
		border:    bootstrap.BsTableBorderColorBorderedDefault,
		size:      bootstrap.BsTablSizeDefault,
		// heading:   []htmlwrapper.Elm{},
		// body:      []htmlwrapper.Elm{},
		footer:  nil,
		caption: nil,
	}
}
