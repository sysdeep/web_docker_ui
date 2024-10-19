import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import TotalReport from './total_report';
import {
  ApiConfigListModel,
  ConfigsServices,
} from '../../services/configs_service';
import IconConfigs from '../../components/icon_configs';
import ConfigsTable from './configs_table';
import { useConfiguration } from '@src/store/configuration';
import ButtonRefresh from '@src/components/button_refresh';

export default function ConfigsPage() {
  const { configuration } = useConfiguration();

  const configs_service = useMemo(() => {
    return new ConfigsServices(configuration.base_url);
  }, []);

  const [configs, setConfigs] = useState<ApiConfigListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    configs_service
      .get_configs()
      .then((configs: ApiConfigListModel[]) => {
        setConfigs(configs);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log('page configs mounted');
    refresh();
  }, []);

  const on_remove = (id: string) => {
    configs_service
      .remove_config(id)
      .then(() => {
        refresh();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <div>
      <PageTitle>
        <IconConfigs /> Configs
      </PageTitle>

      <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div>

      {/* // TODO //{' '} */}
      {/* <div>
        //{' '}
        <a href='/volumes/actions/prune' class='button error'>
          // <i class='fa fa-trash-o' aria-hidden='true'></i>
          // Prune //{' '}
        </a>
        //{' '}
      </div> */}

      <ConfigsTable configs={configs} on_remove={on_remove} />
      <TotalReport total={configs.length} />
    </div>
  );
}
