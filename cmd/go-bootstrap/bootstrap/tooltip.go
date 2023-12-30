package bootstrap

import "github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"

func TooltipPopoverPre(content string, popover *TooltipPopover) htmlwrapper.Elm {
	attrs := map[string]string{}
	if popover != nil {
		attrs["title"] = popover.Title
		if popover.Content != nil {
			attrs["data-content"] = *popover.Content
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:      "pre",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: content}},
	}
}

func TooltipPopoverA(href, content string, popover *TooltipPopover) htmlwrapper.Elm {
	attrs := map[string]string{
		"href": href,
	}
	if popover != nil {
		attrs["title"] = popover.Title
		if popover.Content != nil {
			attrs["data-content"] = *popover.Content
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag:      "a",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: content}},
	}
}
