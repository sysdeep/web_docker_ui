package webserver

import (
	"context"
	"hdu/internal/logger"
	"hdu/internal/registry_client"
	"hdu/internal/services"
	"hdu/internal/webserver/api"
	"hdu/internal/webserver/pages"
	"hdu/internal/webserver/registry_handler"
	"io/fs"

	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Webserver struct {
	e *echo.Echo
	// docker *client.Client
}

func NewWebserver(
	docker *client.Client,
	registry_client *registry_client.RegistryClient,
	services *services.Services,
	logger *logger.Logger,
	www_fs fs.FS,
	version string,
) *Webserver {

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

	// mount embed fs
	// e.StaticFS("/embed", www_fs)
	e.StaticFS("/", www_fs)

	// static pages
	e.GET("/", func(c echo.Context) error {

		config := pages.MainPageConfiguration{
			BaseURL:  "/",
			Registry: registry_client.IsEnabled(),
			Version:  version,
		}

		return pages.MainPage(c, config, www_fs)
	})

	// api ----------------------------------------------------------------------
	api_handlers := api.NewApi(docker, registry_client, services, logger)

	// containers
	e.GET("/api/containers", api_handlers.GetContainers)
	e.GET("/api/containers/:id", api_handlers.GetContainer)
	e.GET("/api/container_inspect/:id", api_handlers.GetContainerInspect)
	e.GET("/api/container_top/:id", api_handlers.GetContainerTop)
	e.GET("/api/container_stats/:id", api_handlers.GetContainerStats)
	e.POST("/api/container_action", api_handlers.ContainerAction)

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

	// services
	e.GET("/api/services", api_handlers.GetServices)
	e.GET("/api/service/:id", api_handlers.GetService)

	// info
	e.GET("/api/info", api_handlers.GetInfo)

	// registry
	reg_handler := registry_handler.NewRegistryHandler(docker, registry_client)
	e.GET("/api/registry/repositories", reg_handler.GetRegistryRepositories)
	e.GET("/api/registry/repositories_smart", reg_handler.GetRegistryRepositoriesSmart)
	e.GET("/api/registry/repository/:id", reg_handler.GetRegistryRepository)
	e.GET("/api/registry/repository_tag/:id/:tag", reg_handler.GetRegistryRepositoryTag)
	e.DELETE("/api/registry/repository_tag/:id/:tag", reg_handler.RemoveRegistryRepositoryTag)
	e.POST("/api/registry/action/:action", reg_handler.RegistryAction)

	return &Webserver{
		e: e,
	}
}

func (w *Webserver) Start() error {
	return w.e.Start("localhost:1313")
}

func (w *Webserver) Shutdown(ctx context.Context) error {
	return w.e.Shutdown(ctx)
}
