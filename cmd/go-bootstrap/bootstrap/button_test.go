package bootstrap

import "testing"

func TestButton(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondary(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmit(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmit(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" type=\"button\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" type=\"submit\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestButtonCheckbox(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-primary btn-lg\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckbox(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-lg\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxId(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-primary btn-lg\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxId(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-lg\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxValue(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value2", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-primary btn-lg\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxValue(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value2", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-lg\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdValue(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value2", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-primary btn-lg\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdValue(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value2", BsBtnPrimary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-lg\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-secondary btn-lg\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-lg\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-secondary btn-lg\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-lg\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxValueSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value2", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-secondary btn-lg\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxValueSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value2", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-lg\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdValueSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value2", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-secondary btn-lg\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdValueSecondary(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value2", BsBtnSecondary, BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-lg\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-primary btn-sm\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-sm\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-primary btn-sm\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-sm\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxValueSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value2", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-primary btn-sm\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxValueSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value2", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-sm\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdValueSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value2", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-primary btn-sm\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdValueSize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value2", BsBtnPrimary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-primary btn-sm\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-secondary btn-sm\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-sm\" for=\"id\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-secondary btn-sm\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-sm\" for=\"idOther\">value</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxValueSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "id", "value2", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"radio\"></input><label class=\"btn btn-secondary btn-sm\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxValueSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "id", "value2", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"id\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-sm\" for=\"id\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxIdValueSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(false, "idOther", "value2", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"radio\"></input><label class=\"btn btn-secondary btn-sm\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonCheckboxCheckboxIdValueSecondarySize(t *testing.T) {
	// ButtonCheckbox(checkbox bool, id string, value string, btnType BsBtnStyle, size BsBtnSize)
	html, err := ButtonCheckbox(true, "idOther", "value2", BsBtnSecondary, BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input autocomplete=\"off\" class=\"btn-check\" id=\"idOther\" type=\"checkbox\"></input><label class=\"btn btn-secondary btn-sm\" for=\"idOther\">value2</label>" {
		t.Fatalf("HTML is not as expected!")
	}
}

///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////
///////////////////////

func TestButtonAnchor(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondary(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmit(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmit(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-lg active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondaryContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindButton, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSubmitContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnPrimary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-primary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" data-content=\"content\" href=\"LinkHere\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonAnchorSecondarySubmitContentSmallActiveTooltipContent2(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := ButtonAnchor("LinkHere2", BsBtnSecondary, BsBtnKindSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<a class=\"btn btn-secondary btn-sm active\" data-content=\"content\" href=\"LinkHere2\" title=\"title\">text2</a>" {
		t.Fatalf("HTML is not as expected!")
	}
}
