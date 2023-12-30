package bootstrap

import "testing"

// type DrowndownItem struct {
// 	// Divider bool
// 	Kind DrowndownItemKind
// 	Href string
// 	Text string
// }

func TestDrowndown(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownDown, BsBtnPrimary,
		BsBtnNormal, "text", BsBtnStateNormal, BsBtnStateDisabled,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"dropdown\"><button aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-primary  dropdown-toggle\" data-bs-toggle=\"dropdown\" disabled=\"disabled\" id=\"dropdownMenuButton\" type=\"button\" value=\"text\">text</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestDrowndownDirectionUp(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownUp, BsBtnPrimary,
		BsBtnNormal, "text", BsBtnStateNormal, BsBtnStateDisabled,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"btn-group dropleft\"><div class=\"dropdown\"><button aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-primary  dropdown-toggle\" data-bs-toggle=\"dropdown\" disabled=\"disabled\" id=\"dropdownMenuButton\" type=\"button\" value=\"text\">text</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestDrowndownDirectionUpSecondary(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownUp, BsBtnSecondary,
		BsBtnNormal, "text", BsBtnStateNormal, BsBtnStateDisabled,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"btn-group dropleft\"><div class=\"dropdown\"><button aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-secondary  dropdown-toggle\" data-bs-toggle=\"dropdown\" disabled=\"disabled\" id=\"dropdownMenuButton\" type=\"button\" value=\"text\">text</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestDrowndownDirectionUpSecondaryBtnLg(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownUp, BsBtnSecondary,
		BsBtnLg, "text", BsBtnStateNormal, BsBtnStateDisabled,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"btn-group dropleft\"><div class=\"dropdown\"><button aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-secondary  dropdown-toggle\" data-bs-toggle=\"dropdown\" disabled=\"disabled\" id=\"dropdownMenuButton\" type=\"button\" value=\"text\">text</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestDrowndownDirectionUpSecondaryBtnLgText(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownUp, BsBtnSecondary,
		BsBtnLg, "click me", BsBtnStateNormal, BsBtnStateDisabled,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"btn-group dropleft\"><div class=\"dropdown\"><button aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-secondary  dropdown-toggle\" data-bs-toggle=\"dropdown\" disabled=\"disabled\" id=\"dropdownMenuButton\" type=\"button\" value=\"click me\">click me</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestDrowndownDirectionUpSecondaryBtnLgTextActive(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownUp, BsBtnSecondary,
		BsBtnLg, "click me", BsBtnStateActive, BsBtnStateDisabled,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"btn-group dropleft\"><div class=\"dropdown\"><button aria-disabled=\"true\" aria-expanded=\"false\" class=\"btn btn-secondary active dropdown-toggle\" data-bs-toggle=\"dropdown\" disabled=\"disabled\" id=\"dropdownMenuButton\" type=\"button\" value=\"click me\">click me</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestDrowndownDirectionUpSecondaryBtnLgTextActiveActive(t *testing.T) {
	// Drowndown(direction BsDropdownDirection, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, items []DrowndownItem)
	html, err := Drowndown(BsDropdownDirectionDownUp, BsBtnSecondary,
		BsBtnLg, "click me", BsBtnStateActive, BsBtnStateActive,
		[]DrowndownItem{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"btn-group dropleft\"><div class=\"dropdown\"><button aria-expanded=\"false\" class=\"btn btn-secondary active active dropdown-toggle\" data-bs-toggle=\"dropdown\" id=\"dropdownMenuButton\" type=\"button\" value=\"click me\">click me</button><div aria-labelledby=\"dropdownMenuButton\" class=\"dropdown-menu\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
