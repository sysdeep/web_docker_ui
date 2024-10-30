import { join_url } from '@src/routes';

export default class ServicesService {
  private base_url: string;

  constructor(base_url: string) {
    this.base_url = base_url;
    console.log('services_service created');
  }

  async get_services(): Promise<ServiceListModel[]> {
    const response = await fetch(join_url(this.base_url, '/api/services'));
    const data = (await response.json()) as ServicesListModel;
    return data.services || [];
  }

  async get_service(id: string): Promise<ServiceListModel> {
    const response = await fetch(join_url(this.base_url, '/api/service/' + id));
    const data = (await response.json()) as ServiceListModel;
    return data;
  }

  // async remove_volume(name: string): Promise<void> {
  //   await fetch(join_url(this.base_url, '/api/volumes/' + name), {
  //     method: 'DELETE',
  //   });
  //   return;
  // }
}

// list models ----------------------------------------------------------------
type ReplicatedService = {
  replicas: number;
};

export type ServiceMode = {
  replicated: null | ReplicatedService;
  global: null | any; // NOTE: приходит только признак - есть или нет
};

export interface ServiceListModel {
  id: string;
  name: string;
  image: string;
  mode: ServiceMode;
  created_at: string;
  updated_at: string;
}

interface ServicesListModel {
  services: ServiceListModel[];
  total: number;
}

// volume model ---------------------------------------------------------------
// export interface ApiFullVolumeModel {
//   volume: ApiVolumeModel;
//   containers: ApiContainerListModel[];
// }

// interface ApiVolumeModel {
//   name: string;
//   created: string;
//   driver: string;
//   mount_point: string;
// }
