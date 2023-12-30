package bootstrap

import (
	"testing"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

func TestPagination(t *testing.T) {
	// Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm)
	html, err := Pagination(BsNavJustifyContentLeft, BsPaginationSizeNormal, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"pagination  \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestPaginationJustifyContentCenter(t *testing.T) {
	// Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm)
	html, err := Pagination(BsNavJustifyContentCenter, BsPaginationSizeNormal, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"pagination justify-content-center \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestPaginationJustifyContentRight(t *testing.T) {
	// Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm)
	html, err := Pagination(BsNavJustifyContentRight, BsPaginationSizeNormal, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"pagination justify-content-end \"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestPaginationJustifyContentRightSizeSmall(t *testing.T) {
	// Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm)
	html, err := Pagination(BsNavJustifyContentRight, BsPaginationSizeSmall, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"pagination justify-content-end pagination-sm\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestPaginationJustifyContentRightSizeBig(t *testing.T) {
	// Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm)
	html, err := Pagination(BsNavJustifyContentRight, BsPaginationSizeBig, []htmlwrapper.Elm{}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"pagination justify-content-end pagination-lg\"></ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestPaginationJustifyContentRightSizeBigContent(t *testing.T) {
	// Pagination(justifyContent BsNavJustifyContent, size BsPaginationSize, contents []htmlwrapper.Elm)
	html, err := Pagination(BsNavJustifyContentRight, BsPaginationSizeBig, []htmlwrapper.Elm{&htmlwrapper.TextElm{Content: "test"}}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<ul class=\"pagination justify-content-end pagination-lg\">test</ul>" {
		t.Fatalf("HTML is not as expected!")
	}
}
func TestPaginationItem(t *testing.T) {
	// PaginationItem(state BsPaginationItemState, href string, content htmlwrapper.Elm)
	html, err := PaginationItem(BsPaginationItemStateNormal, "", &htmlwrapper.TextElm{Content: ""}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<li class=\"pagination\"><a class=\"page-link \" href=\"\"></a></li>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestPaginationItemHref(t *testing.T) {
	// PaginationItem(state BsPaginationItemState, href string, content htmlwrapper.Elm)
	html, err := PaginationItem(BsPaginationItemStateNormal, "href", &htmlwrapper.TextElm{Content: ""}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<li class=\"pagination\"><a class=\"page-link \" href=\"href\"></a></li>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestPaginationItemHrefContent(t *testing.T) {
	// PaginationItem(state BsPaginationItemState, href string, content htmlwrapper.Elm)
	html, err := PaginationItem(BsPaginationItemStateNormal, "href", &htmlwrapper.TextElm{Content: "content"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<li class=\"pagination\"><a class=\"page-link \" href=\"href\">content</a></li>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestPaginationItemHrefContentActive(t *testing.T) {
	// PaginationItem(state BsPaginationItemState, href string, content htmlwrapper.Elm)
	html, err := PaginationItem(BsPaginationItemStateActive, "href", &htmlwrapper.TextElm{Content: "content"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<li class=\"pagination\"><a class=\"page-link active\" href=\"href\">content</a></li>" {
		t.Fatalf("HTML is not as expected!")
	}
}

func TestPaginationItemHrefContentDisabled(t *testing.T) {
	// PaginationItem(state BsPaginationItemState, href string, content htmlwrapper.Elm)
	html, err := PaginationItem(BsPaginationItemStateDisabled, "href", &htmlwrapper.TextElm{Content: "content"}).AsHTML()
	if err != nil {
		t.Fatalf("Failed to make HTML! %s", err)
	}
	if html != "<li class=\"pagination\"><a class=\"page-link disabled\" href=\"href\">content</a></li>" {
		t.Fatalf("HTML is not as expected!")
	}
}
