package bootstrap

import "testing"

func TestButton(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondary(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmit(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmit(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmall(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateNormal, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallActive(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateActive, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallActiveTooltip(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateNormal, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnLg, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-lg active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondaryContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnNormal, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSubmitContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnPrimary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-primary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonSecondarySubmitContentSmallActiveTooltipContent(t *testing.T) {
	// Button(btnType BsBtnStyle, kind BsBtnKind, text string, size BsBtnSize, state BsBtnState, popover *TooltipPopover)
	content := "content"
	html, err := Button(BsBtnSecondary, BsBtnSubmit, "text2", BsBtnSm, BsBtnStateActive, &TooltipPopover{Title: "title", Content: &content}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn btn-secondary btn-sm active\" data-content=\"content\" title=\"title\" value=\"text2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
// // // // // // // // // // // // // // // // // // // // // // // // // // //
// // // // // // // // // // // // // // // // // // // // // // // // // // //
// // // // // // // // // // // // // // // // // // // // // // // // // // //
// // // // // // // // // // // // // // // // // // // // // // // // // // //
// // // // // // // // // // // // // // // // // // // // // // // // // // //
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
