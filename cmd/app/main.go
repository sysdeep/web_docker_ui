package main

import (
	"hdu/internal/logger"
	"hdu/internal/registry_client"
	"hdu/internal/services"
	"hdu/internal/webserver"

	"github.com/docker/docker/client"
)

func main() {
	log := logger.NewLogger()
	log.Info("start")

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
	web_server := webserver.NewWebserver(d_client, r_client, servs, log)
	web_server.Start()
}
