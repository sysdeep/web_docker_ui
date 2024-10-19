import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';

interface StatusFrameProps {
  container: ApiContainerResponseModel;
}

export default function StatusFrame({ container }: StatusFrameProps) {
  const ip_adresses: string[] = Object.keys(container.network.networks).map(
    (network_name) => {
      const net = container.network.networks[network_name];
      return net.ip_address;
    },
  );
  return (
    <div className='box'>
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
              <td>{container.container.name}</td>
            </tr>
            <tr>
              <td>Ip address</td>
              <td>{ip_adresses.join(', ')}</td>
            </tr>
            <tr>
              <td>Status</td>
              <td>{container.state.status}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{container.container.created} </td>
            </tr>
            <tr>
              <td>Start time</td>
              <td>{container.state.started}</td>
            </tr>
            <tr>
              <td>RestartCount</td>
              <td>{container.container.restart_count}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div>
        <ul>
          <li>
            <a href='/container/{container.container.id}/logs'>Logs TODO</a>
          </li>
          <li>
            <a href='/container/{container.container.id}/inspect'>
              Inspect TODO
            </a>
          </li>
          <li>
            <a href='/container/{container.container.id}/stats'>Stats TODO</a>
          </li>
          <li>
            <a href='/container/{container.container.id}/console'>
              Console TODO
            </a>
          </li>
          <li>
            <a href='/container/{container.container.id}/attach'>Attach TODO</a>
          </li>
        </ul>
      </div>
    </div>
  );
}
