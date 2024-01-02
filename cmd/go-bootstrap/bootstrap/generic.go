package bootstrap

const (
	BsFade        string = "fade"
	BsShow               = "show"
	BsRoundedPill        = "rounded-pill"
)

type BsTextBg string

const (
	BsTextBgPrimary   BsTextBg = "text-bg-primary"
	BsTextBgSecondary          = "text-bg-secondary"
	BsTextBgSuccess            = "text-bg-success"
	BsTextBgDanger             = "text-bg-danger"
	BsTextBgWarning            = "text-bg-warning"
	BsTextBgInfo               = "text-bg-info"
	BsTextBgLight              = "text-bg-light"
	BsTextBgDark               = "text-bg-dark"
)

type BsBtnStyle string

const (
	BsBtnStylePrimary   BsBtnStyle = "btn-primary"
	BsBtnStyleSecondary            = "btn-secondary"
	BsBtnStyleSuccess              = "btn-success"
	BsBtnStyleDanger               = "btn-danger"
	BsBtnStyleWarning              = "btn-warning"
	BsBtnStyleInfo                 = "btn-info"
	BsBtnStyleLight                = "btn-light"
	BsBtnStyleDark                 = "btn-dark"
	BsBtnStyleLink                 = "btn-link"

	BsBtnStyleOutlinePrimary   = "btn-outline-primary"
	BsBtnStyleOutlineSecondary = "btn-outline-secondary"
	BsBtnStyleOutlineSuccess   = "btn-outline-success"
	BsBtnStyleOutlineDanger    = "btn-outline-danger"
	BsBtnStyleOutlineWarning   = "btn-outline-warning"
	BsBtnStyleOutlineInfo      = "btn-outline-info"
	BsBtnStyleOutlineLight     = "btn-outline-light"
	BsBtnStyleOutlineDark      = "btn-outline-dark"
)

type BsBtnSize string

const (
	BsBtnNormal BsBtnSize = ""
	BsBtnLg               = "btn-lg"
	BsBtnSm               = "btn-sm"
)

type BsBtnKind string

const (
	BsBtnKindSubmit BsBtnKind = "submit"
	BsBtnKindButton           = "button"
	BsBtnKindReset            = "reset"
)

type BsBtnState string

const (
	BsBtnStateDisabled BsBtnState = "disabled"
	BsBtnStateActive              = "active"
	BsBtnStateNormal              = ""
)

type BsBtnGroupSize string

const (
	BtnGroupSizeNormal BsBtnGroupSize = ""
	BtnGroupSizeLg                    = "btn-group-lg"
	BtnGroupSizeSm                    = "btn-group-sm"
)

const (
	BsBorderPrimary   string = "border-primary"
	BsBorderSecondary        = "border-secondary"
	BsBorderSuccess          = "border-success"
	BsBorderDanger           = "border-danger"
	BsBorderWarning          = "border-warning"
	BsBorderInfo             = "border-info"
	BsBorderLight            = "border-light"
	BsBorderDark             = "border-dark"
)

const (
	BsTextPrimary   string = "text-primary"
	BsTextSecondary        = "text-secondary"
	BsTextSuccess          = "text-success"
	BsTextDanger           = "text-danger"
	BsTextWarning          = "text-warning"
	BsTextInfo             = "text-info"
	BsTextLight            = "text-light"
	BsTextDark             = "text-dark"
)

const BsBgTransparent string = "bg-transparent"

type BsCarouselAutoplayStyle string

const (
	BsCarouselAutoplayNo                 BsCarouselAutoplayStyle = ""
	BsCarouselAutoplayImmediate                                  = "carousel"
	BsCarouselAutoplayOnFirstInteraction                         = "true"
)

type BsDropdownDirection string

const (
	BsDropdownDirectionDownDown  BsDropdownDirection = ""
	BsDropdownDirectionDownUp                        = "dropup"
	BsDropdownDirectionDownLeft                      = "dropleft"
	BsDropdownDirectionDownRight                     = "dropright"
)

type DrowndownItemKind string

const (
	BsDrowndownItemKindItem         DrowndownItemKind = "item"
	BsDrowndownItemKindItemActive                     = "itemActive"
	BsDrowndownItemKindItemDisabled                   = "itemDisabled"
	BsDrowndownItemKindDivider                        = "divider"
	BsDrowndownItemKindHeader                         = "header"
)

type BsInputType string

const (
	BsInputTypeButton   BsInputType = "button"
	BsInputTypeCheckbox             = "checkbox"
	BsInputTypeColor                = "color"
	BsInputTypeDate                 = "date"
	BsInputTypeDatetime             = "datetime-local"
	BsInputTypeEmail                = "email"
	BsInputTypeFile                 = "file"
	BsInputTypeHidden               = "hidden"
	BsInputTypeImage                = "image"
	BsInputTypeMonth                = "month"
	BsInputTypeNumber               = "number"
	BsInputTypePassword             = "password"
	BsInputTypeRadio                = "radio"
	BsInputTypeRange                = "range"
	BsInputTypeReset                = "reset"
	BsInputTypeSearch               = "search"
	BsInputTypeSubmit               = "submit"
	BsInputTypeTel                  = "tel"
	BsInputTypeText                 = "text"
	BsInputTypeTime                 = "time"
	BsInputTypeUrl                  = "url"
	BsInputTypeWeek                 = "week"
)

type BsInputBtnType string

const (
	BsInputBtnTypeButton BsInputType = "button"
	BsInputBtnTypeReset              = "reset"
	BsInputBtnTypeSubmit             = "submit"
)

type BsInputSize string

const (
	BsInputSizeNormal BsInputSize = ""
	BsInputSizeLarge              = " form-control-lg"
	BsInputSizeSmall              = " form-control-sm"
)

type BsColSize string

const (
	BsColSizeAuto BsColSize = ""
	BsColSize1              = "col-1"
	BsColSize2              = "col-2"
	BsColSize3              = "col-3"
	BsColSize4              = "col-4"
	BsColSize5              = "col-5"
	BsColSize6              = "col-6"
	BsColSize7              = "col-7"
	BsColSize8              = "col-8"
	BsColSize9              = "col-9"
	BsColSize10             = "col-10"
	BsColSize11             = "col-11"
	BsColSize12             = "col-12"
	BsColSizeSm             = "col-sm"
	BsColSizeMd             = "col-md"
	BsColSizeLg             = "col-lg"
	BsColSizeXl             = "col-xl"
)

type ListGroupItemKind string

const (
	ListGroupItemKindPrimary   ListGroupItemKind = "list-group-item-primary"
	ListGroupItemKindSecondary                   = "list-group-item-secondary"
	ListGroupItemKindSuccess                     = "list-group-item-success"
	ListGroupItemKindDanger                      = "list-group-item-danger"
	ListGroupItemKindWarning                     = "list-group-item-warning"
	ListGroupItemKindInfo                        = "list-group-item-info"
	ListGroupItemKindLight                       = "list-group-item-light"
	ListGroupItemKindDark                        = "list-group-item-dark"
)

type ModelSize string

const (
	ModelSizeNormal ModelSize = ""
	ModelSizeSmall            = "bd-example-modal-sm"
	ModelSizeLarge            = "bd-example-modal-lg"
)

type TooltipPopover struct {
	Title   string
	Content *string
}

type BsNavState string

const (
	BsNavStateDisabled BsNavState = "disabled"
	BsNavStateActive              = "active"
	BsNavStateNormal              = ""
)

type BsNavKind string

const (
	BsNavKindNormal BsNavKind = ""
	BsNavKindPills            = "nav-pills"
	BsNavKindTabs             = "nav-tabs"
)

type BsNavJustifyContent string

const (
	BsNavJustifyContentLeft   BsNavJustifyContent = ""
	BsNavJustifyContentCenter                     = "justify-content-center"
	BsNavJustifyContentRight                      = "justify-content-end"
)

type BsColor string

const (
	BsColorPrimary   BsColor = "primary"
	BsColorSecondary         = "secondary"
	BsColorSuccess           = "success"
	BsColorDanger            = "danger"
	BsColorWarning           = "warning"
	BsColorInfo              = "info"
	BsColorLight             = "light"
	BsColorDark              = "dark"
	BsColorWhite             = "white"
)

type BsLocation string

const (
	BsLocationNormal      BsLocation = ""
	BsLocationFixedTop               = "fixed-top"
	BsLocationFixedBottom            = "fixed-bottom"
	BsLocationStickyTop              = "sticky-top"
)

type BsPaginationItemState string

const (
	BsPaginationItemStateDisabled BsPaginationItemState = "disabled"
	BsPaginationItemStateActive                         = "active"
	BsPaginationItemStateNormal                         = ""
)

type BsPaginationSize string

const (
	BsPaginationSizeNormal BsPaginationSize = ""
	BsPaginationSizeSmall                   = "pagination-sm"
	BsPaginationSizeBig                     = "pagination-lg"
)

type BsBg string

const (
	BsBgPrimary   BsBg = "bg-primary"
	BsBgSecondary      = "bg-secondary"
	BsBgSuccess        = "bg-success"
	BsBgDanger         = "bg-danger"
	BsBgWarning        = "bg-warning"
	BsBgInfo           = "bg-info"
	BsBgLight          = "bg-light"
	BsBgDark           = "bg-dark"
)

type BsTableCellKind string

const (
	BsTableCellKindRow    BsTableCellKind = "row"
	BsTableCellKindCol                    = "col"
	BsTableCellKindNormal                 = ""
)

type BsTableColor string

const (
	BsTableColorPrimary   BsTableColor = "table-primary"
	BsTableColorSecondary              = "table-secondary"
	BsTableColorSuccess                = "table-success"
	BsTableColorDanger                 = "table-danger"
	BsTableColorWarning                = "table-warning"
	BsTableColorInfo                   = "table-info"
	BsTableColorLight                  = "table-light"
	BsTableColorDark                   = "table-dark"
	BsTableColorDefault                = ""
)

type BsTableBorderColor string

const (
	BsTableBorderColorPrimary         BsTableBorderColor = "table-bordered border-primary"
	BsTableBorderColorSecondary                          = "table-bordered border-secondary"
	BsTableBorderColorSuccess                            = "table-bordered border-success"
	BsTableBorderColorDanger                             = "table-bordered border-danger"
	BsTableBorderColorWarning                            = "table-bordered border-warning"
	BsTableBorderColorInfo                               = "table-bordered border-info"
	BsTableBorderColorLight                              = "table-bordered border-light"
	BsTableBorderColorDark                               = "table-bordered border-dark"
	BsTableBorderColorBorderedDefault                    = "table-bordered"
	BsTableBorderColorDefault                            = ""
	BsTableBorderColorBorderless                         = "table-borderless"
)

type BsTableSize string

const (
	BsTablSizeLarge   BsTableSize = "table-lg"
	BsTablSizeSmall               = "table-sm"
	BsTablSizeDefault             = ""
)
