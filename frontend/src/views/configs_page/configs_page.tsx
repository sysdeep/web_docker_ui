import PageTitle from "../../components/page_title";
import { useEffect, useMemo, useState } from "react";
import TotalReport from "./total_report";
import { ApiConfigListModel, ConfigsServices } from "../../services/configs_service";
import IconConfigs from "../../components/icon_configs";
import ConfigsTable from "./configs_table";
import { useConfiguration } from "@src/store/configurationContext";

export default function ConfigsPage() {
  const { base_url } = useConfiguration();

  const configs_service = useMemo(() => {
    return new ConfigsServices(base_url);
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
    console.log("page configs mounted");
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
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconConfigs /> Configs
      </PageTitle>

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
