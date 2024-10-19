import React from 'react';
import { Link } from 'react-router-dom';
import { ApiInfoModel } from '../../services/info_service';
import { format_size } from '../../utils/humanize';
import IconContainers from '../../components/icon_containers';
import { route } from '../../routes';

interface ServerFrameProps {
  info: ApiInfoModel;
}
export default function ServerFrame({ info }: ServerFrameProps) {
  const networks_view = info.system.default_addresses_pools.map((pool, idx) => {
    return (
      <span key={idx}>
        [{pool.base} size: {pool.size}]
      </span>
    );
  });
  return (
    <div className='card mb-2'>
      <div className='card-body'>
        <h2>Server</h2>

        <table className='table table-sm'>
          <tbody>
            <tr>
              <td>Hostname</td>
              <td className='text-end'>{info.system.name}</td>
            </tr>
            <tr>
              <td>Server version</td>
              <td className='text-end'>{info.system.server_version}</td>
            </tr>
            <tr>
              <td>
                <Link to={route.containers}>Containers</Link>
              </td>
              <td className='text-end'>
                <span className='ml-2' title='total'>
                  <IconContainers />
                  &nbsp;
                  {info.system.containers}
                </span>
                <span className='ml-2' title='running'>
                  <i className='bi bi-play-fill text-success'></i>
                  &nbsp;
                  {info.system.containers_running}
                </span>
                <span className='ml-2' title='stopped'>
                  <i className='bi bi-stop-fill text-error'></i>
                  &nbsp;
                  {info.system.containers_stopped}
                </span>
                <span className='ml-2' title='paused'>
                  <i className='bi bi-pause-fill text-grey'></i>
                  &nbsp;
                  {info.system.containers_paused}
                </span>
              </td>
            </tr>
            <tr>
              <td>
                <Link to={route.images}>Images</Link>
              </td>
              <td className='text-end'>{info.system.images}</td>
            </tr>
            <tr>
              <td>OperatingSystem</td>
              <td className='text-end'>
                {info.system.operating_system}({info.system.kernel_version})
              </td>
            </tr>
            <tr>
              <td>Address pool</td>
              <td className='text-end'>{networks_view}</td>
            </tr>
            <tr>
              <td>Default runtime</td>
              <td className='text-end'>{info.system.default_runtime}</td>
            </tr>
            <tr>
              <td>HW</td>
              <td className='text-end'>
                <span className='ml-2'>
                  <i className='bi bi-cpu'></i>
                  &nbsp;
                  {info.system.ncpu}
                </span>
                <span className='ml-2'>
                  <i className='bi bi-sd-card'></i>
                  &nbsp;
                  {format_size(info.system.mem_total)}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
