//go:build prod
// +build prod

package root

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed frontend/dist
var www_real embed.FS

func init() {

	// strip
	xxx, err := fs.Sub(www_real, "frontend/dist")
	if err != nil {
		log.Panic(err)
	}

	WWW = xxx
}
