import { Link } from 'react-router-dom';
import { ApiConfigListModel } from '../../services/configs_service';
import React from 'react';
import { route, join_url } from '../../routes';
import IconRemove from '@src/components/icon_remove';

interface ConfigsTableProps {
  configs: ApiConfigListModel[];
  on_remove(id: string): void;
}

export default function ConfigsTable({
  configs,
  on_remove,
}: ConfigsTableProps) {
  const configs_view = configs.map((config, idx) => {
    const on_remove_click = (e: any, id: string) => {
      e.preventDefault();
      on_remove(id);
    };

    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.config, config.id)}>{config.name}</Link>
        </td>
        <td> {config.created} </td>
        <td> {config.updated} </td>
        <td>
          <a
            href='#'
            className='error'
            onClick={(e) => on_remove_click(e, config.id)}
          >
            <IconRemove />
            Remove
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
          <th>Created</th>
          <th>Updated</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{configs_view}</tbody>
    </table>
  );
}
