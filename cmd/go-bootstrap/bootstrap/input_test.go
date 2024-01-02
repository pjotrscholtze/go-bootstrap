package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestInputGroup(t *testing.T) {
	// InputGroup(content []htmlwrapper.Elm)
	html, err := InputGroup([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputGroupContent(t *testing.T) {
	// InputGroup(content []htmlwrapper.Elm)
	html, err := InputGroup([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputGroupPrepend(t *testing.T) {
	// InputGroupPrepend(content []htmlwrapper.Elm)
	html, err := InputGroupPrepend([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group-prepend\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputGroupPrependContent(t *testing.T) {
	// InputGroupPrepend(content []htmlwrapper.Elm)
	html, err := InputGroupPrepend([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group-prepend\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputGroupAppend(t *testing.T) {
	// InputGroupAppend(content []htmlwrapper.Elm)
	html, err := InputGroupAppend([]htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group-append\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestInputGroupAppendContent(t *testing.T) {
	// InputGroupAppend(content []htmlwrapper.Elm)
	html, err := InputGroupAppend([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group-append\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputGroupText(t *testing.T) {
	// InputGroupText(content []htmlwrapper.Elm, id string)
	html, err := InputGroupText([]htmlwrapper.Elm{}, "idString").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group-text\" id=\"idString\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputGroupTextContent(t *testing.T) {
	// InputGroupText(content []htmlwrapper.Elm, id string)
	html, err := InputGroupText([]htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "content"}}, "idString").AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"input-group-text\" id=\"idString\">content</div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInput(t *testing.T) {
	// Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool)
	placeholder := "placeholder"
	value := "value"
	html, err := Input(BsInputTypeText, "id", "name", &placeholder, &value, BsInputSizeNormal, false, false, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"form-control \" id=\"id\" name=\"name\" placeholder=\"placeholder\" type=\"text\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputContent(t *testing.T) {
	// Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool)
	placeholder := "text"
	value := "asfasdfsadf"
	html, err := Input(BsInputTypeEmail, "identifier", "email", &placeholder, &value, BsInputSizeSmall, false, false, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"form-control  form-control-sm\" id=\"identifier\" name=\"email\" placeholder=\"text\" type=\"email\" value=\"asfasdfsadf\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputContentReadonly(t *testing.T) {
	// Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool)
	placeholder := "text"
	value := "asfasdfsadf"
	html, err := Input(BsInputTypeEmail, "identifier", "email", &placeholder, &value, BsInputSizeSmall, true, false, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"form-control  form-control-sm\" id=\"identifier\" name=\"email\" placeholder=\"text\" readonly=\"readonly\" type=\"email\" value=\"asfasdfsadf\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputContentPlaintext(t *testing.T) {
	// Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool)
	placeholder := "text"
	value := "asfasdfsadf"
	html, err := Input(BsInputTypeEmail, "identifier", "email", &placeholder, &value, BsInputSizeSmall, false, true, false).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"form-control  form-control-sm form-control-plaintext\" id=\"identifier\" name=\"email\" placeholder=\"text\" type=\"email\" value=\"asfasdfsadf\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputContentChecked(t *testing.T) {
	// Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool)
	placeholder := "text"
	value := "asfasdfsadf"
	html, err := Input(BsInputTypeEmail, "identifier", "email", &placeholder, &value, BsInputSizeSmall, false, false, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input checked=\"checked\" class=\"form-control  form-control-sm\" id=\"identifier\" name=\"email\" placeholder=\"text\" type=\"email\" value=\"asfasdfsadf\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputContentCheckedCheckbox(t *testing.T) {
	// Input(kind BsInputType, id, name string, placeholder, value *string, size BsInputSize, readonly, plaintext, checked bool)
	placeholder := "text"
	value := "asfasdfsadf"
	html, err := Input(BsInputTypeCheckbox, "identifier", "email", &placeholder, &value, BsInputSizeSmall, false, false, true).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"form-check\"><input checked=\"checked\" class=\"form-check-input  form-control-sm\" id=\"identifier\" name=\"email\" type=\"checkbox\" value=\"asfasdfsadf\"></input><label class=\"form-check-label\" for=\"identifier\">text</label></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButton(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value", BsInputSizeNormal, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-primary\" type=\"button\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmit(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeNormal, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-primary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButton2(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value2", BsInputSizeNormal, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-primary\" type=\"button\" value=\"value2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmit2(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeNormal, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-primary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonLarge(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value", BsInputSizeLarge, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-primary\" type=\"button\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmitLarge(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeLarge, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-primary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButton2Large(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value2", BsInputSizeLarge, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-primary\" type=\"button\" value=\"value2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmit2Large(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeLarge, BsBtnStylePrimary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-primary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSecondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value", BsInputSizeNormal, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-secondary\" type=\"button\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmitSecondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeNormal, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-secondary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButton2Secondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value2", BsInputSizeNormal, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-secondary\" type=\"button\" value=\"value2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmit2Secondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeNormal, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  btn-secondary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonLargeSecondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value", BsInputSizeLarge, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-secondary\" type=\"button\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmitLargeSecondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeLarge, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-secondary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButton2LargeSecondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnType(BsInputBtnTypeButton), "value2", BsInputSizeLarge, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-secondary\" type=\"button\" value=\"value2\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestInputButtonSubmit2LargeSecondary(t *testing.T) {
	// InputButton(kind BsInputBtnType, value string, size BsInputSize, btnStyle BsBtnStyle)
	html, err := InputButton(BsInputBtnTypeSubmit, "value", BsInputSizeLarge, BsBtnStyleSecondary).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<input class=\"btn  form-control-lg btn-secondary\" type=\"submit\" value=\"value\"></input>" {
		t.Fatalf("HTML is not as expected!")
	}
}
