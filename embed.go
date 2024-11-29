package root

import (
	_ "embed"
	"io/fs"
)

//go:embed version
var AppVersion string

var WWW fs.FS
