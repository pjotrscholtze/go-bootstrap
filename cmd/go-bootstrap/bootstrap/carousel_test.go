package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestCarousel(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, false, false, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselIndicators(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", true, false, false, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselCaptions(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, true, false, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"></div><button class=\"carousel-control-prev\" data-bs-slide=\"prev\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-prev-icon\"></span><span class=\"visually-hidden\">Previous</span></button><button class=\"carousel-control-next\" data-bs-slide=\"next\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-next-icon\"></span><span class=\"visually-hidden\">Next</span></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselIndicatorsCaptions(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", true, true, false, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"></div><button class=\"carousel-control-prev\" data-bs-slide=\"prev\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-prev-icon\"></span><span class=\"visually-hidden\">Previous</span></button><button class=\"carousel-control-next\" data-bs-slide=\"next\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-next-icon\"></span><span class=\"visually-hidden\">Next</span></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

////////

func TestCarouselCarouselFade(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, false, true, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide carousel-fade\" id=\"id\"><div class=\"carousel-inner\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselIndicatorsCarouselFade(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", true, false, true, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide carousel-fade\" id=\"id\"><div class=\"carousel-inner\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselCaptionsCarouselFade(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, true, true, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide carousel-fade\" id=\"id\"><div class=\"carousel-inner\"></div><button class=\"carousel-control-prev\" data-bs-slide=\"prev\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-prev-icon\"></span><span class=\"visually-hidden\">Previous</span></button><button class=\"carousel-control-next\" data-bs-slide=\"next\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-next-icon\"></span><span class=\"visually-hidden\">Next</span></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselIndicatorsCaptionsCarouselFade(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", true, true, true, BsCarouselAutoplayNo, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide carousel-fade\" id=\"id\"><div class=\"carousel-inner\"></div><button class=\"carousel-control-prev\" data-bs-slide=\"prev\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-prev-icon\"></span><span class=\"visually-hidden\">Previous</span></button><button class=\"carousel-control-next\" data-bs-slide=\"next\" data-bs-target=\"#id\" type=\"button\"><span aria-hidden=\"true\" class=\"carousel-control-next-icon\"></span><span class=\"visually-hidden\">Next</span></button></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

/////

func TestCarouselAutoplayImmediate(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, false, false, BsCarouselAutoplayImmediate, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselAutoplayOnFirstInteraction(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, false, false, BsCarouselAutoplayOnFirstInteraction, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestCarouselContent(t *testing.T) {
	// Carousel(id string, indicators, captions, carouselFade bool, autoplayStyle BsCarouselAutoplayStyle, contents []htmlwrapper.Elm)
	html, err := Carousel("id", false, false, false, BsCarouselAutoplayOnFirstInteraction, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "asdf"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"carousel slide\" id=\"id\"><div class=\"carousel-inner\"><div class=\"carousel-item active\">asdf</div></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
