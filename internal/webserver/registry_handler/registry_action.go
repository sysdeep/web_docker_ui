package registry_handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

// handler
func (h *RegistryHandler) RegistryAction(c echo.Context) error {

	action := c.Param("action")
	fmt.Println(action)

	if action == "1" {
		err := runGarbageCollector(h.docker_client, "registry")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	if action == "2" {
		err := runRestartRegistry(h.docker_client, "registry")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	slog.Info("action ok")
	return c.JSON(http.StatusOK, nullResponse{})
}

func runGarbageCollector(docker_client *client.Client, registry_name string) error {
	slog.Info("run garbage collextor action")
	// cmd := []string{"ls"}
	// cmd := []string{"uptime"}
	cmd := []string{registry_name, "garbage-collect", "/etc/docker/registry/config.yml"}
	exec_options := container.ExecOptions{
		Cmd:          cmd,
		Detach:       false,
		AttachStdout: true,
		Tty:          false,
		AttachStderr: true,
	}

	// prepare exec
	res, err := docker_client.ContainerExecCreate(context.Background(), registry_name, exec_options)
	if err != nil {
		return err
	}
	// fmt.Println(res)

	// exec_start_options := container.ExecStartOptions{}
	// err = h.docker_client.ContainerExecStart(context.Background(), res.ID, exec_start_options)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err)
	// }

	// attach exec
	// exec_attach_options := container.ExecAttachOptions{}
	// attach_result, err := docker_client.ContainerExecAttach(context.Background(), res.ID, exec_attach_options)
	// if err != nil {
	// 	return err
	// }
	//
	// for {
	// 	out_str, err := attach_result.Reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		fmt.Println("end of reading, break")
	// 		break
	// 	}
	//
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(string(out_str))
	//
	// }

	// inspect exec
	inspect_result, err := docker_client.ContainerExecInspect(context.Background(), res.ID)
	if err != nil {
		return err
	}

	// utils.PrintAsJson(inspect_result)

	if inspect_result.ExitCode != 0 {

		return errors.New(fmt.Sprintf("error to exec garbage, exit code: %d", inspect_result.ExitCode))
	}

	return nil
}

func runRestartRegistry(docker_client *client.Client, registry_name string) error {
	slog.Info("restart registry action")
	restart_options := container.StopOptions{}
	return docker_client.ContainerRestart(context.Background(), registry_name, restart_options)

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
