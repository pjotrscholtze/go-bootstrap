package bootstrap

import (
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type TableCaption struct {
	Content    htmlwrapper.Elm
	CaptionTop bool
}

func Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption) htmlwrapper.Elm {
	out := &htmlwrapper.HTMLElm{
		Tag: "table",
		Attrs: map[string]string{
			"class": "table",
		},
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag:      "thead",
				Contents: heading,
			},
			&htmlwrapper.HTMLElm{
				Tag:      "tbody",
				Contents: body,
			},
		},
	}
	if footer != nil {
		out.Contents = append(out.Contents, &htmlwrapper.HTMLElm{
			Tag:      "tfoot",
			Contents: *footer,
		})
	}
	if caption != nil {
		out.Contents = append([]htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag:      "caption",
				Contents: []htmlwrapper.Elm{caption.Content},
			},
		}, out.Contents...)
		if caption.CaptionTop {
			out.Attrs["class"] += " caption-top"
		}
	}
	if size != BsTablSizeDefault {
		out.Attrs["class"] += " " + string(size)
	}
	if border != BsTableBorderColorBorderedDefault {
		out.Attrs["class"] += " " + string(border)
	}
	if color != BsTableColorDefault {
		out.Attrs["class"] += " " + string(color)
	}
	if striped {
		out.Attrs["class"] += " table-striped"
	}
	if hoverable {
		out.Attrs["class"] += " table-hover"
	}
	return out
}

func TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	out := &htmlwrapper.HTMLElm{
		Tag: "tr",
		Attrs: map[string]string{
			"class": "",
		},
		Contents: contents,
	}
	if color != BsTableColorDefault {
		out.Attrs["class"] += " " + string(color)
	}
	if active {
		out.Attrs["class"] += " table-active"
	}
	return out
}

func TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	out := &htmlwrapper.HTMLElm{
		Tag: "td",
		Attrs: map[string]string{
			"class": "",
		},
		Contents: contents,
	}
	if color != BsTableColorDefault {
		out.Attrs["class"] += " " + string(color)
	}
	if active {
		out.Attrs["class"] += " table-active"
	}
	if kind == BsTableCellKindRow {
		out.Tag = "th"
		out.Attrs["scope"] = "row"
	} else if kind == BsTableCellKindCol {
		out.Tag = "th"
		out.Attrs["scope"] = "col"
	}

	if colSpan != 1 {
		out.Attrs["colspan"] = strconv.FormatUint(uint64(colSpan), 10)
	}
	return out
}

func TableResponsive(content htmlwrapper.Elm) htmlwrapper.Elm {
	out := &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "table-responsive",
		},
		Contents: []htmlwrapper.Elm{content},
	}
	return out
}
