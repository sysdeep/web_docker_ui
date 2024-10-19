import { ApiInfoModel } from '../../services/info_service';
import React from 'react';

interface ClientFrameProps {
  info: ApiInfoModel;
}

export default function ClientFrame({ info }: ClientFrameProps) {
  return (
    <div className='card mb-2'>
      <div className='card-body'>
        <h2>Client</h2>
        <table className='table table-sm'>
          <tbody>
            <tr>
              <td>DaemonHost</td>
              <td className='text-end'> {info.daemon_host}</td>
            </tr>
            <tr>
              <td>ClientVersion</td>
              <td className='text-end'>{info.client_version}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
