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
