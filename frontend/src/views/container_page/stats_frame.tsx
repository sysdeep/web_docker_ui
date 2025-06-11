import ButtonRefresh from "@src/components/button_refresh";
import { ApiContainerResponseModel, ContainerTopModel, useContainersService } from "@src/services/containers_service";
import { useConfiguration } from "@src/store/configurationContext";
import { useEffect, useState } from "react";

type StatsFrameProps = {
  container: ApiContainerResponseModel | null;
};

export default function StatsFrame({ container }: StatsFrameProps) {
  const { base_url } = useConfiguration();
  const { get_container_stats } = useContainersService(base_url);

  // state --------------------------------------------------------------------
  const [loading, setLoading] = useState<boolean>(false);
  const [containerTop, setContainerTop] = useState<ContainerTopModel | null>(null);

  const refresh_stats = () => {
    if (container && container.state.status === "running") {
      setLoading(true);
      get_container_stats(container.container.id)
        .then(setContainerTop)
        .catch((err) => console.log(err))
        .finally(() => setLoading(false));
    } else {
      setContainerTop(null);
    }
  };

  useEffect(() => {
    refresh_stats();
  }, [container]);

  // view ---------------------------------------------------------------------
  return (
    <div className='box'>
      <h2>Top</h2>

      {containerTop && <ProcessesTable containerTop={containerTop} />}
      {containerTop && <ButtonRefresh on_refresh={refresh_stats} loading={loading} />}
      {!containerTop && <p>no data</p>}
    </div>
  );
}

// table view -----------------------------------------------------------------
type ProcessesViewProps = {
  containerTop: ContainerTopModel;
};
function ProcessesTable({ containerTop }: ProcessesViewProps) {
  const headers_view = containerTop.titles.map((title, idx) => {
    return <th key={idx}>{title}</th>;
  });

  const body_row = (row: string[]) => {
    return row.map((data, idx) => {
      return <td key={idx}>{data}</td>;
    });
  };

  const body_view = containerTop.processes.map((procs, idx) => {
    return <tr key={idx}>{body_row(procs)}</tr>;
  });

  return (
    <table className='table table-sm'>
      <thead>
        <tr>{headers_view}</tr>
      </thead>
      <tbody>{body_view}</tbody>
    </table>
  );
}
