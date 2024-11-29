package main

import (
	root "hdu"
	"hdu/internal/logger"
	"hdu/internal/registry_client"
	"hdu/internal/services"
	"hdu/internal/webserver"
	"log/slog"

	"github.com/docker/docker/client"
)

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

func main() {

	slog.SetLogLoggerLevel(slog.LevelDebug)

	log := logger.NewLogger()
	slog.Info("start")

	// test docker

	d_client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer d_client.Close()

	// // Получение списка запуцщенных контейнеров(docker ps)
	// containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	// if err != nil {
	// 	panic(err)
	// }

	// // Вывод всех идентификаторов контейнеров
	// for _, container := range containers {
	// 	// fmt.Println(container.ID)
	// 	fmt.Printf("%s %s (status: %s)\n", container.ID, container.Image, container.Status)
	// }

	r_client := registry_client.NewRegistryClient("https://localhost:5000")

	// core
	servs := services.NewServices(d_client)

	// web server
	web_server := webserver.NewWebserver(d_client, r_client, servs, log, root.WWW)
	web_server.Start()
}
