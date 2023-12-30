package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestTable(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(false, false, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{}, []htmlwrapper.Elm{}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table table-lg \"><thead></thead><tbody></tbody></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableStriped(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(true, false, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{}, []htmlwrapper.Elm{}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table table-lg  table-striped\"><thead></thead><tbody></tbody></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableHoverable(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(false, true, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{}, []htmlwrapper.Elm{}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table table-lg  table-hover\"><thead></thead><tbody></tbody></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableStripedHoverable(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(true, true, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{}, []htmlwrapper.Elm{}, nil, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table table-lg  table-striped table-hover\"><thead></thead><tbody></tbody></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableStripedHoverableContent(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(true, true, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "heading"}}, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "body"}}, &[]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "footer"}}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table table-lg  table-striped table-hover\"><thead>heading</thead><tbody>body</tbody><tfoot>footer</tfoot></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableStripedHoverableContentCaption(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(true, true, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "heading"}}, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "body"}}, &[]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "footer"}}, &TableCaption{
		Content: &htmlwrapper.TextElm{Content: "content"},
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table table-lg  table-striped table-hover\"><caption>content</caption><thead>heading</thead><tbody>body</tbody><tfoot>footer</tfoot></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableStripedHoverableContentCaptionTop(t *testing.T) {
	// Table(striped, hoverable bool, color BsTableColor, border BsTableBorderColor, size BsTableSize, heading, body []htmlwrapper.Elm, footer *[]htmlwrapper.Elm, caption *TableCaption)
	html, err := Table(true, true, BsTableColorDefault, BsTableBorderColorDefault, BsTablSizeLarge, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "heading"}}, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "body"}}, &[]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "footer"}}, &TableCaption{
		Content:    &htmlwrapper.TextElm{Content: "content"},
		CaptionTop: true,
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<table class=\"table caption-top table-lg  table-striped table-hover\"><caption>content</caption><thead>heading</thead><tbody>body</tbody><tfoot>footer</tfoot></table>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableRow(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(false, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\"\"></tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableRowActive(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(true, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\" table-active\"></tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableRowColor(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(false, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\" table-primary\"></tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableRowActiveColor(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(true, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\" table-primary table-active\"></tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableRowContent(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(false, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\"\">content</tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableRowActiveContent(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(true, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\" table-active\">content</tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableRowColorContent(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(false, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\" table-primary\">content</tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableRowActiveColorContent(t *testing.T) {
	// TableRow(active bool, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableRow(true, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<tr class=\" table-primary table-active\">content</tr>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCell(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 1, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\"\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActive(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 1, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-active\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKind(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 1, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKind(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 1, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKind(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 1, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKind(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 1, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellSpan(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 3, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\"\" colspan=\"3\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveSpan(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 3, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-active\" colspan=\"3\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindSpan(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 3, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" colspan=\"3\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindSpan(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 3, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" colspan=\"3\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindSpan(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 3, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" colspan=\"3\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindSpan(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 3, BsTableColorDefault, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" colspan=\"3\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 1, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 1, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary table-active\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 1, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 1, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 1, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 1, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellSpanColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 3, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary\" colspan=\"3\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveSpanColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 3, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary table-active\" colspan=\"3\"></td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindSpanColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 3, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" colspan=\"3\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindSpanColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 3, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" colspan=\"3\" scope=\"row\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindSpanColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 3, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" colspan=\"3\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindSpanColor(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 3, BsTableColorPrimary, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" colspan=\"3\" scope=\"col\"></th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 1, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\"\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 1, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-active\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 1, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 1, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 1, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 1, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellSpanContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 3, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\"\" colspan=\"3\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveSpanContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 3, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-active\" colspan=\"3\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindSpanContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 3, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" colspan=\"3\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindSpanContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 3, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" colspan=\"3\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindSpanContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 3, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\"\" colspan=\"3\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindSpanContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 3, BsTableColorDefault, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-active\" colspan=\"3\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 1, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 1, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary table-active\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 1, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 1, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 1, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 1, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellSpanColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindNormal, 3, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary\" colspan=\"3\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveSpanColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindNormal, 3, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<td class=\" table-primary table-active\" colspan=\"3\">content</td>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellRowKindSpanColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindRow, 3, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" colspan=\"3\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveRowKindSpanColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindRow, 3, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" colspan=\"3\" scope=\"row\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableCellColKindSpanColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(false, BsTableCellKindCol, 3, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary\" colspan=\"3\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestTableCellActiveColKindSpanColorContent(t *testing.T) {
	// TableCell(active bool, kind BsTableCellKind, colSpan uint, color BsTableColor, contents []htmlwrapper.Elm)
	html, err := TableCell(true, BsTableCellKindCol, 3, BsTableColorPrimary, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<th class=\" table-primary table-active\" colspan=\"3\" scope=\"col\">content</th>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableResponsive(t *testing.T) {
	// TableResponsive(content htmlwrapper.Elm)
	html, err := TableResponsive(&htmlwrapper.TextElm{Content: "test"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"table-responsive\">test</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTableResponsiveContent(t *testing.T) {
	// TableResponsive(content htmlwrapper.Elm)
	html, err := TableResponsive(&htmlwrapper.TextElm{Content: "content"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"table-responsive\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
