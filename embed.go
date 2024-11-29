package root

import (
	_ "embed"
	"io/fs"
)

//go:embed version
var AppVersion string

// https://echo.labstack.com/docs/cookbook/embed-resources
// //go:embed public
// var embededFiles embed.FS

// func getFileSystem(useOS bool) http.FileSystem {
// 	if useOS {
// 		log.Print("using live mode")
// 		return http.FS(os.DirFS("public"))
// 	}

// 	log.Print("using embed mode")
// 	fsys, err := fs.Sub(embededFiles, "public")
// 	if err != nil {
// 		panic(err)
// 	}

// 	return http.FS(fsys)
// }

var WWW fs.FS
