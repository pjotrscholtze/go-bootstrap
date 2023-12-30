package bootstrap

import (
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type ProgressBar struct {
	ProgressValue float64
	Content       htmlwrapper.Elm
	Bg            BsBg
	Strip         bool
	Animated      bool
}

func Progress(bars []ProgressBar) htmlwrapper.Elm {
	content := make([]htmlwrapper.Elm, len(bars))
	for i, pb := range bars {
		progressValueStr := strconv.FormatFloat(pb.ProgressValue, 'f', 2, 64)
		class := ""
		if pb.Strip {
			class += " progress-bar-striped"
		}
		if pb.Animated {
			class += " progress-bar-animated"
		}

		content[i] = &htmlwrapper.HTMLElm{
			Tag: "div",
			Attrs: map[string]string{
				"class":         "progress-bar " + string(pb.Bg),
				"role":          "progressbar",
				"aria-valuenow": "0",
				"aria-valuemin": "0",
				"aria-valuemax": "100",
				"style":         "width: " + progressValueStr + "%",
			},
			Contents: []htmlwrapper.Elm{pb.Content},
		}
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "progress",
		},
		Contents: content,
	}
}
