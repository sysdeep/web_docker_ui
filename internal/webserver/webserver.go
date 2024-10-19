package webserver

import (
	"hdu/internal/logger"
	"hdu/internal/registry_client"
	"hdu/internal/services"
	"hdu/internal/webserver/api"
	"hdu/internal/webserver/handlers"

	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Webserver struct {
	e *echo.Echo
	// docker *client.Client
}

func NewWebserver(docker *client.Client, registry_client *registry_client.RegistryClient, services *services.Services, logger *logger.Logger) *Webserver {

	e := echo.New()

	e.Static("/static", "public")

	// logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// recover
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORS())

	// setup custom renderer
	tplr := NewTemplater()

	// prev templates
	// template_files := makeTemplatesList("views")
	// t := &Template{
	// 	// templates: template.Must(template.ParseGlob("views/*.html")),
	// 	templates: template.Must(template.ParseFiles(template_files...)),
	// }

	// e.Renderer = t
	e.Renderer = tplr

	// setup custom error renderer
	e.HTTPErrorHandler = customHTTPErrorHandler

	hndls := handlers.NewHandlers(docker, services, logger)

	e.GET("/", hndls.MainPage)
	e.GET("/containers/:id", hndls.ContainerPage)
	e.GET("/containers", hndls.ContainersPage)

	// volumes
	e.GET("/volumes/:name", hndls.VolumePage)
	e.GET("/volumes/actions/prune", hndls.ActionVolumesPrune)
	e.GET("/volumes/actions/remove/:name", hndls.ActionVolumeRemove)
	e.GET("/volumes", hndls.VolumesPage)

	// images
	e.GET("/images/:id", hndls.ImagePage)
	e.GET("/images/actions/remove/:id", hndls.ActionImageRemove)
	e.GET("/images", hndls.ImagesPage)

	// networks
	e.GET("/networks/:id", hndls.NetworkPage)
	e.GET("/networks", hndls.NetworksPage)
	e.GET("/networks/actions/remove/:id", hndls.ActionNetworkRemove)

	// configs
	e.GET("/configs/:id", hndls.ConfigPage)
	e.GET("/configs/actions/remove/:id", hndls.ActionConfigRemove)
	e.GET("/configs", hndls.ConfigsPage)

	// secrets
	e.GET("/secrets/:id", hndls.SecretPage)
	e.GET("/secrets/actions/remove/:name", hndls.ActionSecretRemove)
	e.GET("/secrets", hndls.SecretsPage)
	// e.GET("/qqq", func(c echo.Context) error {

	// 	// return c.Render(200, "aaa", 0)
	// 	return c.Render(200, "aaa.html", 0)
	// })

	// api ----------------------------------------------------------------------
	api_handlers := api.NewApi(docker, registry_client, services, logger)

	// containers
	e.GET("/api/containers", api_handlers.GetContainers)
	e.GET("/api/containers/:id", api_handlers.GetContainer)

	// images
	e.GET("/api/images", api_handlers.GetImages)
	e.GET("/api/images/:id", api_handlers.GetImage)
	e.DELETE("/api/images/:id", api_handlers.RemoveImage)

	// volumes
	e.GET("/api/volumes", api_handlers.GetVolumes)
	e.GET("/api/volumes/:name", api_handlers.GetVolume)
	e.DELETE("/api/volumes/:name", api_handlers.RemoveVolume)

	// networks
	e.GET("/api/networks", api_handlers.GetNetworks)
	e.GET("/api/networks/:id", api_handlers.GetNetwork)
	e.DELETE("/api/networks/:id", api_handlers.RemoveNetwork)

	// configs
	e.GET("/api/configs", api_handlers.GetConfigs)
	e.GET("/api/configs/:id", api_handlers.GetConfig)
	e.DELETE("/api/configs/:id", api_handlers.RemoveConfig)

	// secrets
	e.GET("/api/secrets", api_handlers.GetSecrets)
	e.GET("/api/secrets/:id", api_handlers.GetSecret)
	e.DELETE("/api/secrets/:id", api_handlers.RemoveSecret)

	// info
	e.GET("/api/info", api_handlers.GetInfo)

	// registry
	e.GET("/api/registry/repositories", api_handlers.GetRegistryRepositories)
	e.GET("/api/registry/repository/:id", api_handlers.GetRegistryRepository)
	e.GET("/api/registry/repository_tag/:id/:tag", api_handlers.GetRegistryRepositoryTag)

	return &Webserver{
		e: e,
	}
}

func (w *Webserver) Start() {
	w.e.Logger.Fatal(w.e.Start("localhost:1313"))
}
