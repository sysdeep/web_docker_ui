import ButtonRefresh from "@src/components/button_refresh";
import ContainersService, { ApiContainerResponseModel, ContainerTopModel } from "@src/services/containers_service";
import { useEffect, useState } from "react";

type TopFrameProps = {
  container: ApiContainerResponseModel | null;
  containers_service: ContainersService;
};

export default function TopFrame({ container, containers_service }: TopFrameProps) {
  // state --------------------------------------------------------------------
  const [loading, setLoading] = useState<boolean>(false);
  const [containerTop, setContainerTop] = useState<ContainerTopModel | null>(null);

  const refresh_top = () => {
    if (container && container.state.status === "running") {
      setLoading(true);
      containers_service
        .get_container_top(container.container.id)
        .then(setContainerTop)
        .catch((err) => console.log(err))
        .finally(() => setLoading(false));
    } else {
      setContainerTop(null);
    }
  };

  useEffect(() => {
    refresh_top();
  }, [container]);

  // view ---------------------------------------------------------------------
  return (
    <div className='card my-2'>
      <div className='card-body'>
        <h2>Top</h2>

        {containerTop && <ProcessesTable containerTop={containerTop} />}
        {containerTop && <ButtonRefresh on_refresh={refresh_top} loading={loading} />}
        {!containerTop && <p>no data</p>}
      </div>
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
      return (
        <td key={idx}>
          <code>{data}</code>
        </td>
      );
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
