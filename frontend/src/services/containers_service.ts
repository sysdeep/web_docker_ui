import { join_url } from '@src/routes';
import { ApiContainerListModel } from '@src/models/api_container_list_model';

interface ApiContainersListModel {
  containers: ApiContainerListModel[];
}

export default class ContainersService {
  private base_url: string;

  constructor(base_url: string) {
    this.base_url = base_url;
    console.log('containers_service created');
  }

  async get_containers(): Promise<ApiContainerListModel[]> {
    const response = await fetch(join_url(this.base_url, '/api/containers'));

    const data = (await response.json()) as ApiContainersListModel;

    const containers = data.containers || [];

    return containers;
  }

  // async remove_container(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/containers/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }

  async get_container(id: string): Promise<ApiContainerResponseModel> {
    const response = await fetch(
      join_url(this.base_url, '/api/containers/' + id),
    );
    const data = (await response.json()) as ApiContainerResponseModel;
    return data;
  }
}

interface ApiContainerModel {
  id: string;
  created: string;
  name: string;
  restart_count: number;
  image: string;
}
interface ApiContainerStateModel {
  status: string;
  started: string;
}
interface ApiContainerMountsModel {
  name: string;
  destination: string;
}
interface ApiContainerConfigModel {
  env: string[];
  cmd: string;
  image: string;
  entrypoint: string;
}
interface ApiNetworkSegment {
  gateway: string;
  ip_address: string;
  mac_address: string;
  network_id: string;
}
interface ApiPortSegment {
  host_ip: string;
  host_port: string;
}
type ApiNetworkMap = { [id: string]: ApiNetworkSegment };
type ApiPortMap = { [id: string]: ApiPortSegment[] | null };
interface ApiContainerNetworkModel {
  networks: ApiNetworkMap;
  ports: ApiPortMap;
}

export interface ApiContainerResponseModel {
  container: ApiContainerModel;
  state: ApiContainerStateModel;
  mounts: ApiContainerMountsModel[];
  config: ApiContainerConfigModel;
  network: ApiContainerNetworkModel;
}
