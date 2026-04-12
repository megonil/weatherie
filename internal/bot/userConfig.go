package bot

type FormatType uint

const (
	FormatDefault FormatType = iota
	FormatTable              // table
	FormatPlain              // just like default but without "code" block
	FormatSmall              // breaf 1-line info
	FormatCustom             // TODO:
)

type InfoNodeType uint

const (
	InfoTemp InfoNodeType = iota
	InfoLocation
	InfoPressure
	InfoUV
	InfoWind
	InfoCloud
	InfoIcon
)

type UserConfig struct {
	FormatType    FormatType
	InfoNodes     []InfoNodeType
	SaveLocation  bool
	SavedLocation string
}
