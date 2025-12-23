import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import ContainersFrame from "./containers_frame";
import { useContainersService } from "../../services/containers_service";
import TotalReport from "./total_report";
import IconContainers from "../../components/icon_containers";
import { ApiContainerListModel } from "@src/models/api_container_list_model";
import { useConfiguration } from "@src/store/configurationContext";
import ContainersFilter, { ContainersFilterModel } from "./containers_filter";

const defaultFilter: ContainersFilterModel = {
  status: "all",
};

export default function ContainersPage() {
  const { base_url } = useConfiguration();
  const { get_containers } = useContainersService(base_url);

  const [containers, setContainers] = useState<ApiContainerListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const [filterData, setFilterData] = useState<ContainersFilterModel>(defaultFilter);

  const refresh = () => {
    setLoading(true);
    get_containers()
      .then((containers) => {
        setContainers(containers);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log("containers page mounted");
    refresh();
  }, []);

  const onFilter = (model: ContainersFilterModel) => {
    setFilterData(model);
  };

  // filter records
  const filteredContainers = containers.filter((c) => {
    if (filterData.status == "all") return true;
    return c.state === filterData.status;
  });

  // draw
  return (
    <div>
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconContainers /> Containers
      </PageTitle>

      <ContainersFilter initial_filter={filterData} onChange={onFilter} />
      <ContainersFrame containers={filteredContainers} />

      <TotalReport filtered={filteredContainers.length} total={containers.length} />
    </div>
  );
}
