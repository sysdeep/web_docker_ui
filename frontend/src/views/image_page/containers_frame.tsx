import { ApiFullImageModel } from '../../services/images_service';
import React from 'react';

interface ContainersFrameProps {
  image: ApiFullImageModel;
}
export default function ContainersFrame({ image }: ContainersFrameProps) {
  const rows_view = image.containers.map((container, idx) => {
    return (
      <tr key={idx}>
        <td>
          {/* TODO */}
          {/* <a href='/containers/ .ID '> .Name </a> */}
          {container.name}
        </td>
      </tr>
    );
  });
  return (
    <div className='box'>
      <h2>Containers</h2>
      <table className='table table-small striped'>
        <thead>
          <tr>
            <th></th>
          </tr>
        </thead>
        <tbody>{rows_view}</tbody>
      </table>
    </div>
  );
}
