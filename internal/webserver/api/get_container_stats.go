package api

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetContainerStats(c echo.Context) error {

	id := c.Param("id")

	stats_result, err := h.docker_client.ContainerStatsOneShot(context.Background(), id)

	if err != nil {
		slog.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	defer stats_result.Body.Close()

	result, _ := io.ReadAll(stats_result.Body)

	// TODO: возвращается строка, которую надо преобразовать в структуры и пр...
	fmt.Println("===")
	// fmt.Println(stats_result)

	/*
		{"read":"2024-10-23T07:40:12.013599427Z","preread":"0001-01-01T00:00:00Z","pids_stats":{"current":11,"limit":28484},"blkio_stats":{"io_service_bytes_recursive":[{"major":259,"minor":0,"op":"read","value":40738816},{"major":259,"minor":0,"op":"write","value":1392640}],"io_serviced_recursive":null,"io_queue_recursive":null,"io_service_time_recursive":null,"io_wait_time_recursive":null,"io_merged_recursive":null,"io_time_recursive":null,"sectors_recursive":null},"num_procs":0,"storage_stats":{},"cpu_stats":{"cpu_usage":{"total_usage":329660000,"usage_in_kernelmode":88219000,"usage_in_usermode":241441000},"system_cpu_usage":20575320000000,"online_cpus":8,"throttling_data":{"periods":0,"throttled_periods":0,"throttled_time":0}},"precpu_stats":{"cpu_usage":{"total_usage":0,"usage_in_kernelmode":0,"usage_in_usermode":0},"throttling_data":{"periods":0,"throttled_periods":0,"throttled_time":0}},"memory_stats":{"usage":52981760,"stats":{"active_anon":11100160,"active_file":40472576,"anon":11100160,"anon_thp":0,"file":40898560,"file_dirty":0,"file_mapped":35635200,"file_writeback":0,"inactive_anon":0,"inactive_file":425984,"kernel_stack":180224,"pgactivate":0,"pgdeactivate":0,"pgfault":6008,"pglazyfree":0,"pglazyfreed":0,"pgmajfault":349,"pgrefill":0,"pgscan":0,"pgsteal":0,"shmem":0,"slab":551600,"slab_reclaimable":272376,"slab_unreclaimable":279224,"sock":0,"thp_collapse_alloc":0,"thp_fault_alloc":0,"unevictable":0,"workingset_activate":0,"workingset_nodereclaim":0,"workingset_refault":0},"limit":24973926400},"networks":{"eth0":{"rx_bytes":9996,"rx_packets":84,"rx_errors":0,"rx_dropped":0,"tx_bytes":0,"tx_packets":0,"tx_errors":0,"tx_dropped":0}}}
	*/
	fmt.Println(string(result))
	fmt.Println("===")

	return c.JSON(http.StatusOK, stats_result)
	// return c.JSON(http.StatusOK, newGetContainerTopResponse(top_result))
}

// type getContainerTopResponse struct {
// 	Top containerTopResult `json:"top"`
// }
//
// func newGetContainerTopResponse(model container.ContainerTopOKBody) getContainerTopResponse {
// 	top := containerTopResult{
// 		Processes: model.Processes,
// 		Titles:    model.Titles,
// 	}
//
// 	return getContainerTopResponse{
// 		Top: top,
// 	}
// }
//
// type containerTopResult struct {
// 	// Required: true
// 	Processes [][]string `json:"processes"`
//
// 	// The ps column titles
// 	// Required: true
// 	Titles []string `json:"titles"`
// }
