import { ApiContainerListModel } from "@src/models/api_container_list_model";
import { join_url, route } from "@src/routes";
import { Link } from "react-router-dom";
import ContainerStatusIcon from "./container_status_icon";
import { format_date } from "@src/utils/humanize";
import { strip_container_name } from "@src/utils/containers";
import IconSort from "./icon_sort";
import { SortState } from "@src/data/sort_state";
import { useState } from "react";

interface ContainersFrameProps {
  containers: ApiContainerListModel[];
}

const HEADER_NAME = "name";
const HEADER_CREATED = "created";

export default function ContainersTable({ containers }: ContainersFrameProps) {
  const [sortHeader, setSortHeader] = useState<string>(HEADER_CREATED);
  const [sortAsc, setSortAsc] = useState<boolean>(false);

  const onHeaderClick = (header: string) => {
    if (sortHeader === header) {
      return setSortAsc(!sortAsc);
    }

    setSortHeader(header);
  };

  const name_icon_state = makeSortState(HEADER_NAME, sortHeader, sortAsc);
  const created_icon_state = makeSortState(HEADER_CREATED, sortHeader, sortAsc);

  const sort_k = sortAsc ? 1 : -1;
  const sortedRecords = containers.sort((a, b) => {
    if (sortHeader === HEADER_NAME) {
      return (a.name > b.name ? 1 : -1) * sort_k;
    } else {
      return (a.created > b.created ? 1 : -1) * sort_k;
    }
  });

  const rows_view = sortedRecords.map((container, idx) => {
    return (
      <tr key={idx}>
        <td>
          <ContainerStatusIcon status={container.state} />
          &nbsp;
          <Link to={join_url(route.container, container.id)}>{strip_container_name(container.name)}</Link>
        </td>
        <td>{format_date(container.created)}</td>
        <td>
          <Link to={join_url(route.image, container.image_id)}>{container.image}</Link>
        </td>
        <td>{container.ip_addresses.join(", ")}</td>
      </tr>
    );
  });

  return (
    <table className='table table-sm table-striped'>
      <thead>
        <tr>
          <th role='button' onClick={() => onHeaderClick(HEADER_NAME)}>
            Name
            <span style={{ float: "right" }}>
              <IconSort state={name_icon_state} />
            </span>
          </th>
          <th role='button' onClick={() => onHeaderClick(HEADER_CREATED)}>
            Created
            <span style={{ float: "right" }}>
              <IconSort state={created_icon_state} />
            </span>
          </th>
          <th>Image</th>
          <th>IP</th>
        </tr>
      </thead>
      <tbody>{rows_view}</tbody>
    </table>
  );
}

function makeSortState(header_name: string, cur_header: string, cur_asc: boolean): SortState {
  return cur_header === header_name ? (cur_asc ? SortState.asc : SortState.desc) : SortState.none;
}
