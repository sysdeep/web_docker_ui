import { join_url } from '@src/routes';

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
  default_addresses_pools: ApiNetworkPool[];
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

export default class InfoService {
  private base_url: string;
  constructor(base_url: string) {
    this.base_url = base_url;
    console.log('volumes_service created');
  }

  async get_info(): Promise<ApiInfoModel> {
    const response = await fetch(join_url(this.base_url, '/api/info'));

    const data = (await response.json()) as ApiInfoModel;

    return data;
  }

  // async remove_image(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/images/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }
}
