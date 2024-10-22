import React, { useEffect, useMemo, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

import ActionsFrame from './actions_frame';
import StatusFrame from './status_frame';
import PageTitle from '../../components/page_title';
import ContainersService, { ApiContainerResponseModel, ContainerTopModel } from '../../services/containers_service';
import DetailsFrame from './details_frame';
import VolumesFrame from './volumes_frame';
import NetworksFrame from './networks_frame';
import IconContainers from '../../components/icon_containers';
import { useConfiguration } from '@src/store/configuration';
import ButtonRefresh from '@src/components/button_refresh';
import TopFrame from './top_frame';
import { route } from '@src/routes';

export default function ContainerPage() {
  const { id } = useParams();
  const { configuration } = useConfiguration();
  const navigate = useNavigate();

  const containers_service = useMemo(() => {
    return new ContainersService(configuration.base_url);
  }, []);

  const [container, setContainer] = useState<ApiContainerResponseModel | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    containers_service
      .get_container(id)
      .then((data) => {
        setContainer(data);
        // refresh_top();
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log('container page mounted');
    refresh();
  }, []);

  const on_action = (action: string) => {
    if (id) {
      containers_service.container_action(id, action).then(() => {
        if (action === 'remove') {
          navigate(route.containers);
        } else {
          refresh();
        }
      });
    }
  };

  const main_body = () => {
    if (container === null) {
      return <div>not loaded</div>;
    }

    return (
      <>
        <ActionsFrame status={container.state.status} on_action={on_action} />
        <StatusFrame container={container} />

        <TopFrame container={container} containers_service={containers_service} />

        <DetailsFrame container={container} />
        <NetworksFrame container={container} />
        <VolumesFrame container={container} />
      </>
    );
  };

  return (
    <div>
      <PageTitle>
        <IconContainers /> Container: {container ? container.container.name : 'loading'}
      </PageTitle>
      <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div>
      {main_body()}
    </div>
  );
}
