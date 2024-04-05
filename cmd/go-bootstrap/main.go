package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/builder"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
)

type User struct {
	Username string
	Password *string
	ID       int
	Status   string
}

func main() {
	p := User{
		ID:       1,
		Username: "Foo",
		Status:   "Online",
	}
	ft := builder.NewAdvancedTableBuilder[User]([]User{
		p,
		{
			ID:       2,
			Username: "John",
			Status:   "Offline",
		},
		{
			ID:       3,
			Username: "Test",
			Status:   "Disabled",
		},
	})

	he := builder.NewFormBuilder[User](p).SetDefaultMappings().ActionMethod(builder.HTTPMethodGet, "/")

	tb := ft.GetTableBuilder()
	tb.GetTableMapping().RemoveByFieldName("Password").Set("ID", "Id").MoveToIndex("ID", 0)
	pagination := ft.GetPaginationBuilder()
	pagination.SetResultCount(23).SetCurrentPage(1).SetResultsPerPageOptions([]int{
		10,
		25,
		50,
		100,
	}).SetCurrentResultsPerPage(10)
	h := "#"

	nv := bootstrap.NavBar("id", bootstrap.BsColorPrimary, bootstrap.BsLocationNormal, nil, bootstrap.Nav(true, true, bootstrap.BsNavJustifyContentLeft, bootstrap.BsNavKindNormal, []*bootstrap.NavItem{
		&bootstrap.NavItem{
			Href:     &h,
			Content:  &htmlwrapper.TextElm{Content: "a"},
			NavState: bootstrap.BsNavStateNormal,
			DropDownItems: &[]*bootstrap.NavItem{
				&bootstrap.NavItem{
					Href:          &h,
					Content:       &htmlwrapper.TextElm{"asdfadfs"},
					DropDownItems: &[]*bootstrap.NavItem{},
					NavState:      bootstrap.BsNavStateNormal,
				},
			},
		},
	}))
	nvHTML, _ := nv.AsHTML()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, err := he.AsElm().AsHTML()
		html2, err := ft.AsElm().AsHTML()
		// html3, err := pagination.AsElm().AsHTML()
		_ = err
		fmt.Fprintf(w, `<!doctype html>
		<html lang="en">
		  <head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<title>Bootstrap demo</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
		  </head>
		  <body>
			`+nvHTML+html+html2+`
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>
		  </body>
		</html>`)
	})

	fmt.Printf("Starting server at port 8085\n")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}
