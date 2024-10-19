package api

import (
	"context"
	"fmt"
	"hdu/internal/utils"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) RegistryAction(c echo.Context) error {

	action := c.Param("action")

	// cmd := []string{"ls"}
	// cmd := []string{"uptime"}
	cmd := []string{"registry", "garbage-collect", "/etc/docker/registry/config.yml"}
	exec_options := container.ExecOptions{
		Cmd:          cmd,
		Detach:       false,
		AttachStdout: true,
		Tty:          false,
		AttachStderr: true,
	}

	res, err := h.docker_client.ContainerExecCreate(context.Background(), "registry", exec_options)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	fmt.Println(res)

	// exec_start_options := container.ExecStartOptions{}
	// err = h.docker_client.ContainerExecStart(context.Background(), res.ID, exec_start_options)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err)
	// }

	exec_attach_options := container.ExecAttachOptions{}
	ass, err := h.docker_client.ContainerExecAttach(context.Background(), res.ID, exec_attach_options)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	utils.PrintAsJson(ass)

	// line, _, err := ass.Reader.ReadLine()
	out_str, err := ass.Reader.ReadString('\n')
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// fmt.Println(string(line))
	fmt.Println(string(out_str))

	zzz, err := h.docker_client.ContainerExecInspect(context.Background(), res.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	utils.PrintAsJson(zzz)

	fmt.Println(action)
	return c.JSON(http.StatusOK, stdResponse{status: true})
}

/*
@classmethod
    def _run_gc(cls, cid: str):
        cmd = f'docker exec -it {cid} registry garbage-collect /etc/docker/registry/config.yml'

        result = subprocess.run(shlex.split(cmd),
                                universal_newlines=True,
                                capture_output=True)
        print(result.stdout)

    @classmethod
    def _run_restart(cls, cid: str):
        cmd = f'docker container restart {cid}'

        result = subprocess.run(shlex.split(cmd),
                                universal_newlines=True,
                                capture_output=True)
        print(result.stdout)

*/
