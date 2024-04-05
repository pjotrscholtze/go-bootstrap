package htmlhelper

import (
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func Heading(level uint8, inpt string) htmlwrapper.Elm {
	tag := "h"
	if level < 1 || level > 7 {
		level = 1
	}
	tag += strconv.FormatUint(uint64(level), 10)
	return &htmlwrapper.HTMLElm{
		Tag:      tag,
		Contents: []htmlwrapper.Elm{htmlwrapper.Text(inpt)},
	}
}
