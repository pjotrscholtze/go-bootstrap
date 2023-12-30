package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestButtonGroup(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{}, "ariaLabel", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group btn-lg\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVertical(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{}, "ariaLabel", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group-vertical btn-lg\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupContent(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group btn-lg\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalContent(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group-vertical btn-lg\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupAriaLabel(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{}, "ariaLabel2", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group btn-lg\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalAriaLabel(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{}, "ariaLabel2", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group-vertical btn-lg\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupContentAriaLabel(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel2", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group btn-lg\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalContentAriaLabel(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel2", BsBtnLg).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group-vertical btn-lg\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{}, "ariaLabel", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group btn-sm\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{}, "ariaLabel", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group-vertical btn-sm\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupContentSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group btn-sm\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalContentSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-group-vertical btn-sm\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupAriaLabelSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{}, "ariaLabel2", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group btn-sm\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalAriaLabelSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{}, "ariaLabel2", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group-vertical btn-sm\" role=\"group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupContentAriaLabelSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(false, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel2", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group btn-sm\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonGroupVerticalContentAriaLabelSmall(t *testing.T) {
	// ButtonGroup(vertical bool, content []htmlwrapper.Elm, ariaLabel string, size BsBtnGroupSize)
	html, err := ButtonGroup(true, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel2", BsBtnSm).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel2\" class=\"btn-group-vertical btn-sm\" role=\"group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestButtonToolbar(t *testing.T) {
	// ButtonToolbar(content []htmlwrapper.Elm, ariaLabel string)
	html, err := ButtonToolbar([]htmlwrapper.Elm{}, "ariaLabel").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-toolbar\" role=\"toolbar\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonToolbarLabel(t *testing.T) {
	// ButtonToolbar(content []htmlwrapper.Elm, ariaLabel string)
	html, err := ButtonToolbar([]htmlwrapper.Elm{}, "ariaLabe2l").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabe2l\" class=\"btn-toolbar\" role=\"toolbar\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonToolbarContent(t *testing.T) {
	// ButtonToolbar(content []htmlwrapper.Elm, ariaLabel string)
	html, err := ButtonToolbar([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabel").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabel\" class=\"btn-toolbar\" role=\"toolbar\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestButtonToolbarLabelContent(t *testing.T) {
	// ButtonToolbar(content []htmlwrapper.Elm, ariaLabel string)
	html, err := ButtonToolbar([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "ariaLabe2l").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-label=\"ariaLabe2l\" class=\"btn-toolbar\" role=\"toolbar\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
