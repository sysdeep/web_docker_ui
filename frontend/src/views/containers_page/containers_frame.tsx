import React from 'react';
import { ApiContainerListModel } from '@src/models/api_container_list_model';
import ContainersTable from '@src/components/containers_table';

interface ContainersTableProps {
  containers: ApiContainerListModel[];
}

export default function ContainersFrame({ containers }: ContainersTableProps) {
  return (
    <div>
      <ContainersTable containers={containers} />
    </div>
  );
}
