import { Link } from 'react-router-dom';
import { ApiVolumeListModel } from '../../services/volumes_service';
import React from 'react';
import { route, join_url } from '../../routes';
import IconRemove from '@src/components/icon_remove';

interface VolumesTableProps {
  volumes: ApiVolumeListModel[];
  on_remove(name: string): void;
}

export default function VolumesTable({
  volumes,
  on_remove,
}: VolumesTableProps) {
  const on_remove_click = (e: any, name: string) => {
    e.preventDefault();
    on_remove(name);
  };

  const volumes_view = volumes.map((volume, idx) => {
    const options_view = () => {
      if (volume.used) {
        return <td></td>;
      }

      return (
        <td>
          <a
            href='#'
            className='error'
            onClick={(e) => on_remove_click(e, volume.name)}
          >
            <IconRemove />
            &nbsp; Remove
          </a>
        </td>
      );
    };

    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.volume, volume.name)}>{volume.name}</Link>
        </td>
        <td> {volume.used ? 'yes' : 'no'} </td>
        <td> {volume.stack_name} </td>
        <td> {volume.driver} </td>
        {/* <!-- <td> .Mountpoint </td> --> */}
        <td> {volume.created} </td>
        {options_view()}
      </tr>
    );
  });
  return (
    <table className='table table-sm table-striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Used</th>
          <th>Stack</th>
          <th>Driver</th>
          {/* <!-- <th>Mount point</th> --> */}
          <th>Created</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{volumes_view}</tbody>
    </table>
  );
}
