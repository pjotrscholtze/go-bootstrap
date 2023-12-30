package bootstrap

import (
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm) htmlwrapper.Elm {
	// autoplayStyle // @todo
	// carousel-item interval
	carouselItems := make([]htmlwrapper.Elm, len(contents))
	carouselIndicatorsLen := 0
	if indicators {
		carouselIndicatorsLen = len(contents)
	}
	carouselIndicators := make([]htmlwrapper.Elm, carouselIndicatorsLen)
	for i, content := range contents {
		activeStr := ""
		carouselIndicatorAttrs := map[string]string{
			"type":             "button",
			"data-bs-target":   "#" + id,
			"data-bs-slide-to": strconv.Itoa(i),
			"class":            activeStr,
			"aria-label":       "Slide " + strconv.Itoa(i),
		}
		if i == 0 {
			activeStr = " active"
			carouselIndicatorAttrs["aria-current"] = "true"
		}
		carouselItems[i] = &htmlwrapper.HTMLElm{
			Tag: "div",
			Attrs: map[string]string{
				"class": "carousel-item" + activeStr,
			},
			Contents: []htmlwrapper.Elm{
				content,
			},
		}
		if indicators {
			carouselIndicators[i] = &htmlwrapper.HTMLElm{
				Tag:   "button",
				Attrs: carouselIndicatorAttrs,
			}
		}
	}
	carouselContents := append(carouselIndicators, &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "carousel-inner",
		},
		Contents: carouselItems,
		// @todo contents
	})

	if captions {
		carouselContents = append(carouselContents,
			&htmlwrapper.HTMLElm{
				Tag: "button",
				Attrs: map[string]string{
					"class":          "carousel-control-prev",
					"type":           "button",
					"data-bs-target": "#" + id,
					"data-bs-slide":  "prev",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "span",
						Attrs: map[string]string{
							"class":       "carousel-control-prev-icon",
							"aria-hidden": "true",
						},
					},
					&htmlwrapper.HTMLElm{
						Tag: "span",
						Attrs: map[string]string{
							"class": "visually-hidden",
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.TextElm{Content: "Previous"}, // @todo Previous should be dynamic
						},
					},
				},
			},

			&htmlwrapper.HTMLElm{
				Tag: "button",
				Attrs: map[string]string{
					"class":          "carousel-control-next",
					"type":           "button",
					"data-bs-target": "#" + id,
					"data-bs-slide":  "next",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "span",
						Attrs: map[string]string{
							"class":       "carousel-control-next-icon",
							"aria-hidden": "true",
						},
					},
					&htmlwrapper.HTMLElm{
						Tag: "span",
						Attrs: map[string]string{
							"class": "visually-hidden",
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.TextElm{Content: "Next"}, // @todo Previous should be dynamic
						},
					},
				},
			})
	}

	carouselFadeStr := ""
	if carouselFade {
		carouselFadeStr = " carousel-fade"
	}
	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "carousel slide" + carouselFadeStr,
			"id":    id,
		},
		Contents: carouselContents,
	}
}
