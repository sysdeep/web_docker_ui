import { Link } from 'react-router-dom';
import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';
import { route, join_url } from '../../routes';

interface VolumesFrameProps {
  container: ApiContainerResponseModel;
}

export default function VolumesFrame({ container }: VolumesFrameProps) {
  const rows_view = container.mounts.map((volume, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.volume, volume.name)}>{volume.name}</Link>
        </td>
        <td>{volume.destination}</td>
      </tr>
    );
  });
  return (
    <div className='box'>
      <h2>Volumes</h2>
      <div>
        <table className='table is-striped is-fullwidth'>
          <thead>
            <tr>
              <th>Host/volume</th>
              <th>Path in container</th>
            </tr>
          </thead>
          <tbody>{rows_view}</tbody>
        </table>
      </div>
    </div>
  );
}
