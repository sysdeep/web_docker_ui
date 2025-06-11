import { join_url } from "@src/routes";

interface ApiSwarmInfo {
  node_id: string;
  node_addr: string;
}

interface ApiNetworkPool {
  base: string;
  size: number;
}

interface ApiSystemModel {
  id: string;
  name: string;
  containers: number;
  containers_running: number;
  containers_paused: number;
  containers_stopped: number;
  images: number;
  driver: string;
  kernel_version: string;
  operating_system: string;
  os_version: string;
  os_type: string;
  swarm: ApiSwarmInfo;
  default_addresses_pools: ApiNetworkPool[] | null;
  server_version: string;
  default_runtime: string;
  ncpu: number;
  mem_total: number;
}

export interface ApiInfoModel {
  daemon_host: string;
  client_version: string;
  system: ApiSystemModel;
}

export function useInfoService(base_url: string) {
  const get_info = async (): Promise<ApiInfoModel> => {
    const response = await fetch(join_url(base_url, "/api/info"));
    if (!response.ok) {
      throw new Error((await response.json()).message);
    }
    const data = (await response.json()) as ApiInfoModel;
    return data;
  };

  return { get_info };
}
