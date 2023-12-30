package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "ul",
		Attrs: map[string]string{
			"class": "pagination " + string(justifyContent) + " " + string(size),
		},
		Contents: contents,
	}
}
func PaginationItem(state BsPaginationItemState, href string, content htmlwrapper.Elm) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "li",
		Attrs: map[string]string{
			"class": "pagination",
		},
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "a",
				Attrs: map[string]string{
					"class": "page-link " + string(state),
					"href":  href,
				},
				Contents: []htmlwrapper.Elm{content},
			},
		},
	}
}
