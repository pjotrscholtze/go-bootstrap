package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestModalHeader(t *testing.T) {
	// ModalHeader(headerText string, closeBtn bool)
	html, err := ModalHeader("headerText", false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-header\"><h5 class=\"modal-title\">headerText</h5></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestModalHeaderContent(t *testing.T) {
	// ModalHeader(headerText string, closeBtn bool)
	html, err := ModalHeader("ASDFasdf", false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-header\"><h5 class=\"modal-title\">ASDFasdf</h5></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestModalHeaderClose(t *testing.T) {
	// ModalHeader(headerText string, closeBtn bool)
	html, err := ModalHeader("headerText", true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-header\"><h5 class=\"modal-title\">headerText</h5><button aria-label=\"Close\" class=\"close\" data-dismiss=\"modal\" type=\"button\"><span aria-hidden=\"true\">&amp;times;</span></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModalBody(t *testing.T) {
	// ModalBody(content []htmlwrapper.Elm)
	html, err := ModalBody([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-body\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestModalBodyContent(t *testing.T) {
	// ModalBody(content []htmlwrapper.Elm)
	html, err := ModalBody([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "asdfasdf"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-body\">asdfasdf</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestModalFooter(t *testing.T) {
	// ModalFooter(content []htmlwrapper.Elm, closeBtn *string)
	html, err := ModalFooter([]htmlwrapper.Elm{}, nil).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-footer\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestModalFooterCloseBtn(t *testing.T) {
	// ModalFooter(content []htmlwrapper.Elm, closeBtn *string)
	closeBtn := "closeBtn"
	html, err := ModalFooter([]htmlwrapper.Elm{}, &closeBtn).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-footer\"><button class=\"btn btn-secondary\" data-dismiss=\"modal\" type=\"button\">closeBtn</button></div>" {
		t.Fatalf("HTML is not as expected!")
	}

}

func TestModalFooterCloseBtnContent(t *testing.T) {
	// ModalFooter(content []htmlwrapper.Elm, closeBtn *string)
	closeBtn := "closeBtn"
	html, err := ModalFooter([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, &closeBtn).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal-footer\">content<button class=\"btn btn-secondary\" data-dismiss=\"modal\" type=\"button\">closeBtn</button></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModal(t *testing.T) {
	// Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm)
	html, err := Modal(ModelSizeNormal, false, false, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialog\" role=\"document\"><div class=\"modal-content\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModalLarge(t *testing.T) {
	// Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm)
	html, err := Modal(ModelSizeLarge, false, false, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialogbd-example-modal-lg\" role=\"document\"><div class=\"modal-content\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModalSmall(t *testing.T) {
	// Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm)
	html, err := Modal(ModelSizeSmall, false, false, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialogbd-example-modal-sm\" role=\"document\"><div class=\"modal-content\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModalFade(t *testing.T) {
	// Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm)
	html, err := Modal(ModelSizeNormal, true, false, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialog fade\" role=\"document\"><div class=\"modal-content\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModalCentered(t *testing.T) {
	// Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm)
	html, err := Modal(ModelSizeNormal, false, true, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialog modal-dialog-centered\" role=\"document\"><div class=\"modal-content\"></div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModalContent(t *testing.T) {
	// Modal(modelSize ModelSize, fade, modalDialogCentered bool, modalContent []htmlwrapper.Elm)
	html, err := Modal(ModelSizeNormal, false, false, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"modal\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialog\" role=\"document\"><div class=\"modal-content\">content</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestModelWithButton(t *testing.T) {
	// ModelWithButton(modelSize ModelSize, fade bool, id string, modalDialogCentered bool, btnType BsBtnStyle, btnKind BsBtnKind, btnText string, btnSize BsBtnSize, btnState BsBtnState, modalContent []htmlwrapper.Elm)
	html, err := ModelWithButton(ModelSizeNormal, false, "idString", false, BsBtnPrimary, BsBtnKindSubmit, "text", BsBtnSm, BsBtnStateNormal, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div aria-hidden=\"true\" class=\"modal\" id=\"idString\" role=\"dialog\" tabindex=\"-1\"><div class=\"modal-dialog\" role=\"document\"><div class=\"modal-content\"></div></div></div><input class=\"btn btn-primary btn-sm\" data-target=\"#idString\" data-toggle=\"modal\" type=\"submit\" value=\"text\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}
