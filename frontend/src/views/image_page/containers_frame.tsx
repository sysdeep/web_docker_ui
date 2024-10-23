import ContainersTable from '@src/components/containers_table';
import { ApiFullImageModel } from '../../services/images_service';
import React from 'react';
import { ApiContainerListModel } from '@src/models/api_container_list_model';

interface ContainersFrameProps {
  containers: ApiContainerListModel[];
}

export default function ContainersFrame({ containers }: ContainersFrameProps) {
  return (
    <div className='card my-2'>
      <div className='card-body'>
        <h2>Containers</h2>

        <ContainersTable containers={containers} />
      </div>
    </div>
  );
}
