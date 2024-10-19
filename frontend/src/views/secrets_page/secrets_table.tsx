import { Link } from 'react-router-dom';
import { ApiSecretListModel } from '../../services/secrets_service';
import React from 'react';
import { route, join_url } from '../../routes';
import IconRemove from '@src/components/icon_remove';

interface SecretsTableProps {
  secrets: ApiSecretListModel[];
  on_remove(id: string): void;
}

export default function SecretsTable({
  secrets,
  on_remove,
}: SecretsTableProps) {
  const on_remove_click = (e: any, id: string) => {
    e.preventDefault();
    on_remove(id);
  };

  const secrets_view = secrets.map((secret, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.secret, secret.id)}>{secret.name}</Link>
        </td>
        <td> {secret.created} </td>
        <td> {secret.updated} </td>
        <td>
          <a
            href='#'
            className='error'
            onClick={(e) => on_remove_click(e, secret.id)}
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
          <th>Created</th>
          <th>Updated</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{secrets_view}</tbody>
    </table>
  );
}
