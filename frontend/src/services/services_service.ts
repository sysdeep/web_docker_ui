import { Service } from '@src/models/service';
import { join_url } from '@src/routes';

export default class ServicesService {
  private base_url: string;

  constructor(base_url: string) {
    this.base_url = base_url;
    console.log('services_service created');
  }

  async get_services(): Promise<Service[]> {
    const response = await fetch(join_url(this.base_url, '/api/services'));
    const data = (await response.json()) as ServicesListModel;
    return data.services || [];
  }

  async get_service(id: string): Promise<Service> {
    const response = await fetch(join_url(this.base_url, '/api/service/' + id));
    const data = (await response.json()) as Service;
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

interface ServicesListModel {
  services: Service[];
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
