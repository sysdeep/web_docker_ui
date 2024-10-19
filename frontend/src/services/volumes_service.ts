import { ApiContainerListModel } from '@src/models/api_container_list_model';
import { join_url } from '@src/routes';

export default class VolumesService {
  private base_url: string;

  constructor(base_url: string) {
    this.base_url = base_url;
    console.log('volumes_service created');
  }

  async get_volumes(): Promise<ApiVolumeListModel[]> {
    const response = await fetch(join_url(this.base_url, '/api/volumes'));

    const data = (await response.json()) as ApiVolumesListModel;

    return data.volumes || [];
  }

  async get_volume(id: string): Promise<ApiFullVolumeModel> {
    const response = await fetch(join_url(this.base_url, '/api/volumes/' + id));

    const data = (await response.json()) as ApiFullVolumeModel;

    return data;
  }

  async remove_volume(name: string): Promise<void> {
    await fetch(join_url(this.base_url, '/api/volumes/' + name), {
      method: 'DELETE',
    });
    return;
  }
}

// list models ----------------------------------------------------------------
export interface ApiVolumeListModel {
  created: string;
  name: string;
  stack_name: string;
  mount_point: string;
  driver: string;
  used: boolean;
}

interface ApiVolumesListModel {
  volumes: ApiVolumeListModel[];
  total: number;
}

// volume model ---------------------------------------------------------------
export interface ApiFullVolumeModel {
  volume: ApiVolumeModel;
  containers: ApiContainerListModel[];
}

interface ApiVolumeModel {
  name: string;
  created: string;
  driver: string;
  mount_point: string;
}
