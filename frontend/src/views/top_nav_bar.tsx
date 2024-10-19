import React from 'react';
import { Link } from 'react-router-dom';
import IconContainers from '../components/icon_containers';
import IconImages from '../components/icon_images';
import IconVolumes from '../components/icon_volumes';
import IconNetworks from '../components/icon_networks';
import IconConfigs from '../components/icon_configs';
import IconSecrets from '../components/icon_secrets';
import IconHome from '../components/icon_home';
import { route } from '../routes';

export default function TopNavBar() {
  return (
    <nav className='navbar navbar-expand-lg bg-body-tertiary'>
      <div className='container-fluid'>
        <a className='navbar-brand' href='/'>
          Go hdu
        </a>
        <button
          className='navbar-toggler'
          type='button'
          data-bs-toggle='collapse'
          data-bs-target='#navbarSupportedContent'
          aria-controls='navbarSupportedContent'
          aria-expanded='false'
          aria-label='Toggle navigation'
        >
          <span className='navbar-toggler-icon'></span>
        </button>
        <div className='collapse navbar-collapse' id='navbarSupportedContent'>
          <ul className='navbar-nav me-auto mb-2 mb-lg-0'>
            <li className='nav-item'>
              <Link to={'/'} className='nav-link'>
                <IconHome /> Main
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.containers} className='nav-link'>
                <IconContainers /> Containers
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.images} className='nav-link'>
                <IconImages /> Images
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.volumes} className='nav-link'>
                <IconVolumes /> Volumes
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.networks} className='nav-link'>
                <IconNetworks /> Networks
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.configs} className='nav-link'>
                <IconConfigs /> Configs
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.secrets} className='nav-link'>
                <IconSecrets /> Secrets
              </Link>
            </li>
            <li className='nav-item'>
              <Link to={route.registry_repositories} className='nav-link'>
                Registry
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
}
