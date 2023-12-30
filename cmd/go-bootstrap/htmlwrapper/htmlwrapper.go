package htmlwrapper

import (
	"bytes"
	htmltemplate "html/template"
	"text/template"
)

func cleanHTML(inpt string) (string, error) {
	htmlTmpl, err := htmltemplate.New("").Parse("{{.}}")
	var out bytes.Buffer
	err = htmlTmpl.Execute(&out, inpt)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

type Elm interface {
	AsHTML() (string, error)
}

type TextElm struct {
	Content string
}

func (te *TextElm) AsHTML() (string, error) {
	return cleanHTML(te.Content)
}

type MultiElm struct {
	Contents []Elm
}

func (me *MultiElm) AsHTML() (string, error) {
	out := ""
	for _, elem := range me.Contents {
		temp, err := elem.AsHTML()
		if err != nil {
			return "", err
		}
		out += temp
	}
	return out, nil
}

type HTMLElm struct {
	Tag      string
	Attrs    map[string]string
	Contents []Elm
}

func (he *HTMLElm) contentsAsHTML() (string, error) {
	tmp := make([]string, len(he.Contents))
	for i, v := range he.Contents {
		html, err := v.AsHTML()
		if err != nil {
			return "", err
		}
		tmp[i] = html
	}
	tmpl, err := template.New("").Parse("{{range $value := .}}{{$value}}{{end}}")
	if err != nil {
		return "", err
	}
	var tplKey bytes.Buffer
	err = tmpl.Execute(&tplKey, tmp)
	if err != nil {
		return "", err
	}
	return tplKey.String(), nil
}
func (he *HTMLElm) attrsAsHTML() (string, error) {
	tmpl, err := template.New("").Parse("{{range $key, $value := .}} {{$key}}=\"{{$value}}\"{{end}}")
	if err != nil {
		return "", err
	}
	attrs := make(map[string]string)
	for k, v := range he.Attrs {
		key, err := cleanHTML(k)
		if err != nil {
			return "", err
		}
		value, err := cleanHTML(v)
		if err != nil {
			return "", err
		}
		attrs[key] = value
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, attrs)
	return tpl.String(), err
}

func (he *HTMLElm) AsHTML() (string, error) {
	tmpl, err := template.New("").Parse("<{{.tag}}{{.attrs}}>{{.content}}</{{.tag}}>")
	if err != nil {
		return "", err
	}
	attrs, err := he.attrsAsHTML()
	if err != nil {
		return "", err
	}
	content, err := he.contentsAsHTML()
	if err != nil {
		return "", err
	}
	tag, err := cleanHTML(he.Tag)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, map[string]string{
		"tag":     tag,
		"attrs":   attrs,
		"content": content,
	})
	return tpl.String(), err
}
