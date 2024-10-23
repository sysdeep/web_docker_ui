import { ApiContainerListModel } from '@src/models/api_container_list_model';
import { join_url, route } from '@src/routes';
import React from 'react';
import { Link } from 'react-router-dom';
import ContainerStatusIcon from './container_status_icon';
import { format_date } from '@src/utils/humanize';
import { strip_container_name } from '@src/utils/containers';

interface ContainersFrameProps {
  containers: ApiContainerListModel[];
}

export default function ContainersTable({ containers }: ContainersFrameProps) {
  const rows_view = containers.map((container, idx) => {
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
        <td>{container.ip_addresses.join(', ')}</td>
      </tr>
    );
  });

  return (
    <table className='table table-sm table-striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Created</th>
          <th>Image</th>
          <th>IP</th>
        </tr>
      </thead>
      <tbody>{rows_view}</tbody>
    </table>
  );
}
