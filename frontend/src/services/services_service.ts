import { Service } from "@src/models/service";
import { join_url } from "@src/routes";

export function useServicesService(base_url: string) {
  const get_services = async (): Promise<Service[]> => {
    const response = await fetch(join_url(base_url, "/api/services"));
    const data = (await response.json()) as ServicesListModel;
    return data.services || [];
  };

  const get_service = async (id: string): Promise<Service> => {
    const response = await fetch(join_url(base_url, "/api/service/" + id));
    const data = (await response.json()) as Service;
    return data;
  };

  // async remove_volume(name: string): Promise<void> {
  //   await fetch(join_url(base_url, '/api/volumes/' + name), {
  //     method: 'DELETE',
  //   });
  //   return;
  // }

  return {
    get_service,
    get_services,
  };
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
