//go:build !prod
// +build !prod

package root

import (
	"os"
)

func init() {
	WWW = os.DirFS("frontend/dist")
}
