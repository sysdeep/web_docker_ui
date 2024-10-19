import React, { useEffect, useMemo, useState } from 'react';
import { useParams } from 'react-router-dom';

import ActionsFrame from './actions_frame';
import StatusFrame from './status_frame';
import PageTitle from '../../components/page_title';
import ContainersService, {
  ApiContainerResponseModel,
} from '../../services/containers_service';
import DetailsFrame from './details_frame';
import VolumesFrame from './volumes_frame';
import NetworksFrame from './networks_frame';
import IconContainers from '../../components/icon_containers';
import { useConfiguration } from '@src/store/configuration';

export default function ContainerPage() {
  const { id } = useParams();
  const { configuration } = useConfiguration();

  const containers_service = useMemo(() => {
    return new ContainersService(configuration.base_url);
  }, []);

  const [container, setContainer] = useState<ApiContainerResponseModel | null>(
    null,
  );

  const refresh = () => {
    // setLoading(true);
    containers_service
      .get_container(id)
      .then((data) => {
        console.log(data);
        setContainer(data);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        // setLoading(false);
      });
  };

  useEffect(() => {
    console.log('container page mounted');
    refresh();
  }, []);

  const main_body = () => {
    if (container === null) {
      return <div>not loaded</div>;
    }

    return (
      <>
        <ActionsFrame id={id} />
        <StatusFrame container={container} />
        <DetailsFrame container={container} />
        <NetworksFrame container={container} />
        <VolumesFrame container={container} />
      </>
    );
  };

  return (
    <div>
      <PageTitle>
        <IconContainers /> Container:{' '}
        {container ? container.container.name : 'loading'}
      </PageTitle>
      {main_body()}
    </div>
  );
}
