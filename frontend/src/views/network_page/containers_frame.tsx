import React from 'react';
import { ApiNetworkContainerModel } from '../../services/networks_service';
import { Link } from 'react-router-dom';

interface ContainersFrameProps {
  containers: ApiNetworkContainerModel[];
}
export default function ContainersFrame({ containers }: ContainersFrameProps) {
  const rows_view = containers.map((container, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={'/container/' + container.id}> {container.name} </Link>
        </td>
        <td>{container.mac_address}</td>
        <td>{container.ip_v4_address}</td>
      </tr>
    );
  });
  return (
    <div>
      <h2>Containers</h2>
      <table className='table table-small striped'>
        <thead>
          <tr>
            <th>Name</th>
            <th>Mac</th>
            <th>Ip</th>
          </tr>
        </thead>
        <tbody>{rows_view}</tbody>
      </table>
    </div>
  );
}
