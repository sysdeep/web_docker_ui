import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';
import { route, join_url } from '../../routes';
import { Link } from 'react-router-dom';

interface DetailsFrameProps {
  container: ApiContainerResponseModel;
}

export default function DetailsFrame({ container }: DetailsFrameProps) {
  const ports_view = Object.keys(container.network.ports).map(
    (port_name, idx) => {
      const values = container.network.ports[port_name];

      if (!values) {
        return <li key={idx * 100000}>{port_name} - not defined</li>;
      }

      return values.map((segment, idi) => {
        return (
          <li key={(idx + 1) * (idi + 2)}>
            {port_name} - {segment.host_ip}:{segment.host_port}
          </li>
        );
      });
    },
  );

  return (
    <div className='box'>
      <h2>Details</h2>
      <div>
        <table className='table table-small'>
          <tbody>
            <tr>
              <td>Image</td>
              <td>
                <Link to={join_url(route.image, container.container.image)}>
                  {container.config.image}
                </Link>
              </td>
            </tr>
            <tr>
              <td>Ports</td>
              <td>
                <ul>{ports_view}</ul>
              </td>
            </tr>
            <tr>
              <td>CMD</td>
              <td>
                <code>{container.config.cmd}</code>
              </td>
            </tr>
            <tr>
              <td>Entrypoint</td>
              <td>
                <code>{container.config.entrypoint}</code>
              </td>
            </tr>
            <tr>
              <td>ENV</td>
              <td>
                <EnvTable env={container.config.env} />
              </td>
            </tr>
            {/* TODO */}
            {/* <tr>
              <td>Restart policies</td>
              <td>TODO</td>
            </tr> */}
            <tr>
              <td></td>
              <td></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}

interface EnvTableProps {
  env: string[];
}
function EnvTable({ env }: EnvTableProps) {
  const rows = env.map((row, idx) => {
    return (
      <li key={idx}>
        <code>{row}</code>
      </li>
    );
  });

  return <ul>{rows}</ul>;
}
