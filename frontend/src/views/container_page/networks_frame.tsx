import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';
import { route, join_url } from '../../routes';
import { Link } from 'react-router-dom';

interface NetworksFrameProps {
  container: ApiContainerResponseModel;
}

export default function NetworksFrame({ container }: NetworksFrameProps) {
  const networks_view = Object.keys(container.network.networks).map(
    (endpoint, idx) => {
      const net = container.network.networks[endpoint];
      return (
        <tr key={idx}>
          <td>
            <Link to={join_url(route.network, net.network_id)}>{endpoint}</Link>
          </td>
          <td>{net.ip_address}</td>
          <td>{net.gateway}</td>
          <td>{net.mac_address}</td>
          {/* <!-- TODO --> */}
          {/* <!-- <td> --> */}
          {/* <!--   <a href="TODO">Leave TODO</a> --> */}
          {/* <!-- </td> --> */}
        </tr>
      );
    },
  );
  return (
    <div className='box'>
      <h2>Networks</h2>
      {/* <!-- TODO --> */}
      {/* <!-- <div>Connect to: TODO</div> --> */}
      <div>
        <table className='table is-striped is-fullwidth'>
          <thead>
            <tr>
              <th>Network</th>
              <th>IP Address</th>
              <th>Gateway</th>
              <th>MAC</th>
              {/* <!-- <th>Actions</th> --> */}
            </tr>
          </thead>
          <tbody>{networks_view}</tbody>
        </table>
      </div>
    </div>
  );
}
