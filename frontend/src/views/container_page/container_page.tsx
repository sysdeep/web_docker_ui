import React, { useEffect, useMemo, useState } from 'react';
import { useParams } from 'react-router-dom';

import ActionsFrame from './actions_frame';
import StatusFrame from './status_frame';
import PageTitle from '../../components/page_title';
import ContainersService, {
  ApiContainerResponseModel,
  ContainerTopModel,
} from '../../services/containers_service';
import DetailsFrame from './details_frame';
import VolumesFrame from './volumes_frame';
import NetworksFrame from './networks_frame';
import IconContainers from '../../components/icon_containers';
import { useConfiguration } from '@src/store/configuration';
import ButtonRefresh from '@src/components/button_refresh';
import TopFrame from './top_frame';

export default function ContainerPage() {
  const { id } = useParams();
  const { configuration } = useConfiguration();

  const containers_service = useMemo(() => {
    return new ContainersService(configuration.base_url);
  }, []);

  const [container, setContainer] = useState<ApiContainerResponseModel | null>(
    null,
  );
  const [loading, setLoading] = useState<boolean>(false);
  const [containerTop, setContainerTop] = useState<ContainerTopModel | null>(
    null,
  );

  const refresh_top = () => {
    if (container && container.state.status === 'running') {
      containers_service
        .get_container_top(id)
        .then(setContainerTop)
        .catch((err) => console.log(err));
    } else {
      setContainerTop(null);
    }
  };

  const refresh = () => {
    setLoading(true);
    containers_service
      .get_container(id)
      .then((data) => {
        setContainer(data);
        refresh_top();
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

  useEffect(() => {
    refresh_top();
  }, [container]);

  const on_action = (action: string) => {
    console.log(action);
    if (id) {
      containers_service.container_action(id, action).then(() => {
        refresh();
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

        <TopFrame top_model={containerTop} />

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
      <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div>
      {main_body()}
    </div>
  );
}
