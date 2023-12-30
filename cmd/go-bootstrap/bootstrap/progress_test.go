package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestProgress(t *testing.T) {
	html, err := Progress([]ProgressBar{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressSingleBar(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 0,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgPrimary,
		Strip:         false,
		Animated:      false,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-primary\" role=\"progressbar\" style=\"width: 0.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressSingleBar10(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 10,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgPrimary,
		Strip:         false,
		Animated:      false,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-primary\" role=\"progressbar\" style=\"width: 10.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressSingleBar100(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 100,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgPrimary,
		Strip:         false,
		Animated:      false,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-primary\" role=\"progressbar\" style=\"width: 100.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressSingleBar100DifferentBackground(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 100,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgSecondary,
		Strip:         false,
		Animated:      false,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-secondary\" role=\"progressbar\" style=\"width: 100.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressSingleBar100DifferentBackgroundStrip(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 100,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgSecondary,
		Strip:         true,
		Animated:      false,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-secondary\" role=\"progressbar\" style=\"width: 100.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressSingleBar100DifferentBackgroundStripAnimated(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 100,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgSecondary,
		Strip:         true,
		Animated:      true,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-secondary\" role=\"progressbar\" style=\"width: 100.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestProgressMultibar(t *testing.T) {
	html, err := Progress([]ProgressBar{{
		ProgressValue: 100,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgSecondary,
		Strip:         true,
		Animated:      true,
	}, {
		ProgressValue: 0,
		Content:       &htmlwrapper.TextElm{Content: ""},
		Bg:            BsBgPrimary,
		Strip:         false,
		Animated:      false,
	}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<div class=\"progress\"><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-secondary\" role=\"progressbar\" style=\"width: 100.00%\"></div><div aria-valuemax=\"100\" aria-valuemin=\"0\" aria-valuenow=\"0\" class=\"progress-bar bg-primary\" role=\"progressbar\" style=\"width: 0.00%\"></div></div>" {
		t.Fatalf("HTML is not as expected!")
	}
}
