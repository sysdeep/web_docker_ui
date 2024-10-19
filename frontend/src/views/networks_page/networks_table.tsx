import { Link } from 'react-router-dom';
import { ApiNetworkListModel } from '../../services/networks_service';
import React from 'react';
import { route, join_url } from '../../routes';
import IconRemove from '@src/components/icon_remove';

interface NetworksTableProps {
  networks: ApiNetworkListModel[];
  on_remove(id: string): void;
}

export default function NetworksTable({
  networks,
  on_remove,
}: NetworksTableProps) {
  const on_remove_click = (e: any, id: string) => {
    e.preventDefault();
    on_remove(id);
  };

  const networks_view = networks.map((network, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.network, network.id)}>{network.name}</Link>
        </td>
        <td> {network.driver} </td>
        <td> {network.created} </td>
        <td>
          <a
            href='#'
            className='error'
            onClick={(e) => on_remove_click(e, network.id)}
          >
            <IconRemove />
            &nbsp; Remove
          </a>
        </td>
      </tr>
    );
  });
  return (
    <table className='table table-sm table-striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Driver</th>
          <th>Created</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{networks_view}</tbody>
    </table>
  );
}
