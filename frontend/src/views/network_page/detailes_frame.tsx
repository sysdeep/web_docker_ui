import React from 'react';
import { format_size } from '../../utils/humanize';
import { ApiFullNetworkModel } from '../../services/networks_service';
import ButtonRemove from '@src/components/button_remove';

interface DetailsFrameProps {
  network: ApiFullNetworkModel;
  on_remove(): void;
}

export default function DetailsFrame({
  network,
  on_remove,
}: DetailsFrameProps) {
  return (
    <div>
      <h2>Details</h2>
      <div>
        <table className='table table-small'>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{network.network.id}</td>
            </tr>
            <tr>
              <td>Name</td>
              <td>{network.network.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{network.network.created}</td>
            </tr>
            <tr>
              <td>Driver</td>
              <td>{network.network.driver}</td>
            </tr>
            <tr>
              <td>Scope</td>
              <td>{network.network.scope}</td>
            </tr>
            <tr>
              <td>Internal</td>
              <td>{network.network.internal ? 'yes' : 'no'}</td>
            </tr>
            <tr>
              <td>Attachable</td>
              <td>{network.network.attachable ? 'yes' : 'no'}</td>
            </tr>
            <tr>
              <td>Ingress</td>
              <td>{network.network.ingress ? 'yes' : 'no'}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div>
        <ButtonRemove on_remove={on_remove} />
      </div>
    </div>
  );
}
