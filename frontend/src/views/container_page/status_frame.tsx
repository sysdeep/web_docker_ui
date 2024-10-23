import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';
import { format_date } from '@src/utils/humanize';
import ContainerStatusIcon from '@src/components/container_status_icon';
import { strip_container_name } from '@src/utils/containers';
import ActionsFrame from './actions_frame';

interface StatusFrameProps {
  container: ApiContainerResponseModel;
  on_action(action: string): void;
}

export default function StatusFrame({ container, on_action }: StatusFrameProps) {
  const ip_adresses: string[] = Object.keys(container.network.networks).map((network_name) => {
    const net = container.network.networks[network_name];
    return net.ip_address;
  });
  return (
    <div className='card my-2'>
      <div className='card-body'>
        <h2>Status</h2>
        <div>
          <table className='table table-small'>
            <tbody>
              <tr>
                <td>ID</td>
                <td>{container.container.id}</td>
              </tr>
              <tr>
                <td>Name</td>
                <td>{strip_container_name(container.container.name)}</td>
              </tr>
              <tr>
                <td>Ip address</td>
                <td>{ip_adresses.join(', ')}</td>
              </tr>
              <tr>
                <td>Status</td>
                <td>
                  <ContainerStatusIcon status={container.state.status} />
                  &nbsp;
                  {container.state.status}
                </td>
              </tr>
              <tr>
                <td>Created</td>
                <td>{format_date(container.container.created)}</td>
              </tr>
              <tr>
                <td>Start time</td>
                <td>{format_date(container.state.started)}</td>
              </tr>
              <tr>
                <td>RestartCount</td>
                <td>{container.container.restart_count}</td>
              </tr>
            </tbody>
          </table>

          <div className='clearfix'>
            <div className='float-end'>
              <ActionsFrame status={container.state.status} on_action={on_action} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
