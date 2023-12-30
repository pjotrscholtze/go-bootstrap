package bootstrap

import "testing"

func TestTooltipPopoverPre(t *testing.T) {
	html, err := TooltipPopoverPre("", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != `<pre></pre>` {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverPreContent(t *testing.T) {
	html, err := TooltipPopoverPre("asdf", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != `<pre>asdf</pre>` {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverPreTitle(t *testing.T) {
	html, err := TooltipPopoverPre("", &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<pre title=\"title\"></pre>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverPreTitleContent(t *testing.T) {
	content := "content"
	html, err := TooltipPopoverPre("", &TooltipPopover{
		Title:   "title",
		Content: &content,
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<pre data-content=\"content\" title=\"title\"></pre>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverA(t *testing.T) {
	html, err := TooltipPopoverA("", "", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a href=\"\"></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverAHref(t *testing.T) {
	html, err := TooltipPopoverA("href", "", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a href=\"href\"></a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverAContent(t *testing.T) {
	html, err := TooltipPopoverA("", "text", nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a href=\"\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverAContentPopoverTitle(t *testing.T) {
	html, err := TooltipPopoverA("", "text", &TooltipPopover{
		Title: "title",
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a href=\"\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestTooltipPopoverAContentPopoverTitleContent(t *testing.T) {
	content := "content"
	html, err := TooltipPopoverA("", "text", &TooltipPopover{
		Title:   "title",
		Content: &content,
	}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a data-content=\"content\" href=\"\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}
