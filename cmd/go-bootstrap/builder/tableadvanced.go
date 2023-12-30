package builder

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type advancedTableBuilder[T interface{}] struct {
	paginationBuilder PaginationBasicBuilder
	tableBuilder      TableBuilder
}

type AdvancedTableBuilder interface {
	AsElm() htmlwrapper.Elm
	GetPaginationBuilder() PaginationBasicBuilder
	GetTableBuilder() TableBuilder
}

func (tb *advancedTableBuilder[T]) GetPaginationBuilder() PaginationBasicBuilder {
	return tb.paginationBuilder
}
func (tb *advancedTableBuilder[T]) GetTableBuilder() TableBuilder {
	return tb.tableBuilder
}

func (tb *advancedTableBuilder[T]) AsElm() htmlwrapper.Elm {
	pagination := tb.paginationBuilder.AsElm()
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "container",
		},
		Contents: []htmlwrapper.Elm{
			pagination,
			tb.tableBuilder.AsElm(),
			pagination,
		},
	}
}

func NewAdvancedTableBuilder[T interface{}](ref []T) AdvancedTableBuilder {
	return &advancedTableBuilder[T]{
		paginationBuilder: NewPaginationBasicBuilder().SetResultCount(len(ref)),
		tableBuilder:      NewTableBuilder[T](ref),
	}
}
