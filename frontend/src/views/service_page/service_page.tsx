import { useNavigate, useParams } from 'react-router-dom';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import VolumesService, { ApiFullVolumeModel } from '../../services/volumes_service';
import IconVolumes from '../../components/icon_volumes';
import { route } from '@src/routes';
import { useConfiguration } from '@src/store/configuration';
import ContainersFrame from './containers_frame';
import ButtonRemove from '@src/components/button_remove';
import ButtonRefresh from '@src/components/button_refresh';
import IconServices from '@src/components/icon_services';
import ServicesService, { ServiceListModel } from '@src/services/services_service';

export default function ServicePage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { configuration } = useConfiguration();

  const services_service = useMemo(() => {
    return new ServicesService(configuration.base_url);
  }, []);

  const [service, setService] = useState<ServiceListModel | null>(null);

  const refresh = () => {
    services_service.get_service(id).then(setService).catch(console.log);
  };

  useEffect(() => {
    console.log('page service mounted');
    refresh();
  }, []);

  const on_remove = () => {
    console.log('TODO: remove service: ', id);
    // services_service
    //   .remove_volume(id)
    //   .then(() => {
    //     navigate(route.volumes);
    //   })
    //   .catch((err) => {
    //     console.log(err);
    //   });
  };

  const body = () => {
    if (service) {
      return (
        <div>
          {/* actions */}
          <div className='pull-right'>
            {/* TODO */}
            {/* {service.containers.length === 0 && <ButtonRemove on_remove={on_remove} />} */}
            <ButtonRefresh on_refresh={refresh} />
          </div>

          {/* TODO */}
          {/* <DetailsFrame volume={service} /> */}
          {/* <ContainersFrame containers={service.containers} /> */}
        </div>
      );
    }

    return <p>no service</p>;
  };

  return (
    <div>
      <PageTitle>
        <IconServices />
        &nbsp; Service: {id}
      </PageTitle>

      {body()}
    </div>
  );
}
