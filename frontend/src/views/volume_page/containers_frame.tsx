import React from 'react';
import IconContainers from '@src/components/icon_containers';
import { ApiContainerListModel } from '@src/models/api_container_list_model';
import ContainersTable from '@src/components/containers_table';

interface ContainersFrameProps {
  containers: ApiContainerListModel[];
}

export default function ContainersFrame({ containers }: ContainersFrameProps) {
  return (
    <div className='box'>
      <h3>
        <IconContainers />
        &nbsp; Containers
      </h3>
      <ContainersTable containers={containers} />
    </div>
  );
}
