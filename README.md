# Go-bootstrap

This library allows you to easily generate Bootstrap in Go. Besides this it 
provides the ability to generate Bootstrap forms and tables from Golangs
structs. An example is given in `cmd/go-bootstrap/main.go` of using a `User`
struct. This project uses Go test for testing purposes.

## Technical information
### Basic HTML Elements/project core
This project at its core has a very simple Go interface, which can be found at 
`cmd/go-bootstrap/htmlwrapper/htmlwrapper.go`:
```
type Elm interface {
	AsHTML() (string, error)
}
```

All generated HTML elements in this project adhere to this interface, meaning
that it can be easily extended. However, to make it easier the following
structs are available and used to generate the HTML:
- _TextElm_, for text content, a helper/factory function called `Text` is
  available to make this element.
- _HTMLElm_, for HTML elmement allowing you to set the tag, its attributes, and
  contents of more elements following the `Elm` interface.
- _MultiElm_, provides the ability to wrap an array of Elm objects into a
  single `Elm` object. Making the interface standardized, even if you want to
  return multiple elements.

### Bootstrap
All bootstrap elements are available underneath `cmd/go-bootstrap/bootstrap`.
Basically this part of the package contains lots of factories which make
(nested) `Elm` objects. Each file in the directory focuses on a bootstrap
element. Except for, `generic.go`, which contains the constants used, so things
like bootstrap CSS class names for colors, sizes, states, etc. Which is used
in the other files. This folder allows you to basically put together most of a
page if you're using Bootstrap.

### Element generation
However, some of the pages (mostly CRUD pages) would become quite repetitive in
the code written. To reduce this, some builders are made, they can be found at
`cmd/go-bootstrap/builder`. At least the following builders can be found there:
- _Dropdown_, for making dropdown elements.
- _Form_, for making forms from Go structs, allowing renaming, non-native
  datatype handlers, reordering, removing and adding form elements.
- _Input_, for making input elements.
- _Pagination Basic_, for making Bootstrap pagination elements.
- _Table_, for generation of tables from Go struct arrays. Allowing renaming of
  columns columns, reordering, removing columns, adding custom columns, and
  allowing setting handlers of non-native datatypes.
- _Table advanced_, same as `Table` but with pagination.
 

