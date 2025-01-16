package constants

const (
	Domain     = "copilot.meshapi.org"
	JwtSecret  = "DtYPV1pyvJZ3mif1zavOcrSpC"
	ConfigPath = "etc/config.yaml"
)

const (
	UserDataHeader      = "user-data"
	ClientMapUserHeader = "client-user-map"
)

type CompletionType int

const (
	CompletionTypeCode CompletionType = iota
	CompletionTypeChat
)
