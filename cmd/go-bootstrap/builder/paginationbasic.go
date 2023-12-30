package builder

import (
	"math"
	"strconv"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type paginationBasicBuilder struct {
	justifyContent             bootstrap.BsNavJustifyContent
	size                       bootstrap.BsPaginationSize
	resultsPerPagePrependLabel string
	resultCount                int
	currentPage                int
	currentResultsPerPage      int
	resultsPerPageOptions      []int
	hrefGenerator              func(pageNumber int, numberOfResultsPerPage int) string
}
type PaginationBasicBuilder interface {
	AsElm() htmlwrapper.Elm

	GetJustifyContent(justifyContent bootstrap.BsNavJustifyContent) PaginationBasicBuilder
	GetSize(size bootstrap.BsPaginationSize) PaginationBasicBuilder

	SetResultCount(resultCount int) PaginationBasicBuilder
	SetResultsPerPageOptions(resultsPerPageOptions []int) PaginationBasicBuilder
	SetCurrentPage(currentPage int) PaginationBasicBuilder
	SetCurrentResultsPerPage(currentResultsPerPage int) PaginationBasicBuilder

	SetResultsPerPagePrependLabel(resultsPerPagePrependLabel string) PaginationBasicBuilder
	SetHrefGenerator(hrefGenerator func(pageNumber int, numberOfResultsPerPage int) string) PaginationBasicBuilder
}

func (pbb *paginationBasicBuilder) SetResultsPerPagePrependLabel(resultsPerPagePrependLabel string) PaginationBasicBuilder {
	pbb.resultsPerPagePrependLabel = resultsPerPagePrependLabel
	return pbb
}

func (pbb *paginationBasicBuilder) SetHrefGenerator(hrefGenerator func(pageNumber int, numberOfResultsPerPage int) string) PaginationBasicBuilder {
	pbb.hrefGenerator = hrefGenerator
	return pbb
}
func (pbb *paginationBasicBuilder) SetCurrentResultsPerPage(currentResultsPerPage int) PaginationBasicBuilder {
	pbb.currentResultsPerPage = currentResultsPerPage
	return pbb
}

func (pbb *paginationBasicBuilder) SetCurrentPage(currentPage int) PaginationBasicBuilder {
	pbb.currentPage = currentPage
	return pbb
}

func (pbb *paginationBasicBuilder) GetJustifyContent(justifyContent bootstrap.BsNavJustifyContent) PaginationBasicBuilder {
	pbb.justifyContent = justifyContent
	return pbb
}
func (pbb *paginationBasicBuilder) GetSize(size bootstrap.BsPaginationSize) PaginationBasicBuilder {
	pbb.size = size
	return pbb
}

func (pbb *paginationBasicBuilder) SetResultCount(resultCount int) PaginationBasicBuilder {
	pbb.resultCount = resultCount
	return pbb
}
func (pbb *paginationBasicBuilder) SetResultsPerPageOptions(resultsPerPageOptions []int) PaginationBasicBuilder {
	pbb.resultsPerPageOptions = resultsPerPageOptions
	return pbb
}

func (pbb *paginationBasicBuilder) AsElm() htmlwrapper.Elm {
	contents := []htmlwrapper.Elm{}
	// for _, entry := range pbb.resultsPerPageOptions {
	// state := bootstrap.BsPaginationItemStateNormal
	var state bootstrap.BsPaginationItemState
	currentPage := 1
	maxPageNumber := int(math.Ceil(float64(pbb.resultCount) / float64(pbb.currentResultsPerPage)))
	for true {
		state = bootstrap.BsPaginationItemStateNormal
		if currentPage == pbb.currentPage {
			state = bootstrap.BsPaginationItemStateActive
		}
		if maxPageNumber < 10 || currentPage < 4 || currentPage >= maxPageNumber-3 || (currentPage >= pbb.currentPage-3 && currentPage < pbb.currentPage+3) {
			contents = append(contents, bootstrap.PaginationItem(
				state,
				pbb.hrefGenerator(currentPage, pbb.currentResultsPerPage),
				&htmlwrapper.TextElm{
					Content: strconv.FormatInt(int64(currentPage), 10),
				},
			))
		}
		if currentPage*pbb.currentResultsPerPage > pbb.resultCount {
			break
		}
		currentPage += 1
	}
	dropdownItems := []bootstrap.DrowndownItem{}
	for _, option := range pbb.resultsPerPageOptions {
		dropdownItems = append(dropdownItems, bootstrap.DrowndownItem{
			Text: strconv.FormatInt(int64(option), 10),
			Href: pbb.hrefGenerator(pbb.currentPage, option),
			Kind: bootstrap.BsDrowndownItemKindItem,
		})
	}
	return &htmlwrapper.MultiElm{
		Contents: []htmlwrapper.Elm{
			bootstrap.Drowndown(
				bootstrap.BsDropdownDirectionDownDown,
				bootstrap.BsBtnOutlinePrimary,
				bootstrap.BsBtnSubmit,
				pbb.resultsPerPagePrependLabel+strconv.FormatInt(int64(pbb.currentResultsPerPage), 10),
				bootstrap.BsBtnSize(""),
				bootstrap.BsBtnStateNormal,
				dropdownItems,
			),
			bootstrap.Pagination(
				pbb.justifyContent,
				pbb.size,
				contents,
			),
		},
	}
}

func NewPaginationBasicBuilder() PaginationBasicBuilder {
	return &paginationBasicBuilder{
		justifyContent:             bootstrap.BsNavJustifyContentLeft,
		size:                       bootstrap.BsPaginationSizeNormal,
		resultCount:                0,
		currentResultsPerPage:      0,
		resultsPerPageOptions:      []int{},
		resultsPerPagePrependLabel: "Results per page: ",
		hrefGenerator: func(pageNumber int, numberOfResultsPerPage int) string {
			return "?pageNumber=" + strconv.FormatInt(int64(pageNumber), 10) +
				"&numberOfResultsPerPage=" +
				strconv.FormatInt(int64(numberOfResultsPerPage), 10)
		},
	}
}
