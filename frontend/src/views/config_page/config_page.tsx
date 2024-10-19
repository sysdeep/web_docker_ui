import { useParams, useNavigate } from 'react-router-dom';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import {
  ApiFullConfigModel,
  ConfigsServices,
} from '../../services/configs_service';
import IconConfigs from '../../components/icon_configs';
import { route } from '../../routes';
import { useConfiguration } from '@src/store/configuration';

export default function ConfigPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { configuration } = useConfiguration();

  const config_service = useMemo(() => {
    return new ConfigsServices(configuration.base_url);
  }, []);

  const [config, setConfig] = useState<ApiFullConfigModel | null>(null);

  const refresh = () => {
    config_service
      .get_config(id)
      .then((config) => {
        setConfig(config);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page config mounted');
    refresh();
  }, []);

  const on_remove = () => {
    if (config) {
      config_service
        .remove_config(id)
        .then(() => {
          console.log('remove ok');
          navigate(route.configs);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  const body = () => {
    if (config) {
      return (
        <div>
          <DetailsFrame config={config} on_remove={on_remove} />
        </div>
      );
    }

    return <p>no config</p>;
  };

  const page_title = config ? config.config.name : id;

  return (
    <div>
      <PageTitle>
        <IconConfigs />
        &nbsp; Config: {page_title}
      </PageTitle>

      {body()}
    </div>
  );
}
