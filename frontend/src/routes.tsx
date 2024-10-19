import React from 'react';

import Application from './views/application';
import ConfigsPage from './views/configs_page/configs_page';
import ContainerPage from './views/container_page/container_page';
import ContainersPage from './views/containers_page/containers_page';
import ErrorPage from './views/error-page';
import HomePage from './views/home_page/home_page';
import ImagePage from './views/image_page/image_page';
import ImagesPage from './views/images_page/images_page';
import NetworksPage from './views/networks_page/networks_page';
import SecretsPage from './views/secrets_page/secrets_page';
import VolumesPage from './views/volumes_page/volumes_page';
import VolumePage from './views/volume_page/volume_page';
import SecretPage from './views/secret_page/secret_page';
import ConfigPage from './views/config_page/config_page';
import NetworkPage from './views/network_page/network_page';
import { join_url } from './utils/url';
import RepositoriesPage from './views/registry/repositories_page/repositories_page';
import RepositoryPage from './views/registry/repository_page/repository_page';
import RepositoryTagPage from './views/registry/repository_tag_page/repository_tag_page';

export const route = {
  image: '/image',
  images: '/images',
  container: '/container',
  containers: '/containers',
  volume: '/volume',
  volumes: '/volumes',
  network: '/network',
  networks: '/networks',
  config: '/config',
  configs: '/configs',
  secret: '/secret',
  secrets: '/secrets',

  registry_repositories: '/registry/repositories',
  registry_repository: '/registry/repository',
  registry_repository_tag: '/registry/repository_tag',
};

export { join_url };

export const routes = [
  // {
  //   path: '/demo',
  //   element: <div>Hello world!</div>,
  // },
  {
    path: '/',
    // element: <div>Hello world!</div>,
    // element: <Application />,
    // element: <Root />,
    element: <Application />,
    errorElement: <ErrorPage />,

    children: [
      {
        path: '/',
        element: <HomePage />,
      },
      {
        path: route.containers,
        element: <ContainersPage />,
      },
      {
        path: join_url(route.container, ':id'),
        element: <ContainerPage />,
      },
      {
        path: route.images,
        element: <ImagesPage />,
      },
      {
        path: join_url(route.image, ':id'),
        element: <ImagePage />,
      },
      {
        path: route.volumes,
        element: <VolumesPage />,
      },
      {
        path: join_url(route.volume, ':id'),
        element: <VolumePage />,
      },
      {
        path: route.networks,
        element: <NetworksPage />,
      },
      {
        path: join_url(route.network, ':id'),
        element: <NetworkPage />,
      },
      {
        path: route.configs,
        element: <ConfigsPage />,
      },
      {
        path: join_url(route.config, ':id'),
        element: <ConfigPage />,
      },
      {
        path: route.secrets,
        element: <SecretsPage />,
      },
      {
        path: join_url(route.secret, ':id'),
        element: <SecretPage />,
      },

      {
        path: route.registry_repositories,
        element: <RepositoriesPage />,
      },
      {
        path: join_url(route.registry_repository, ':id'),
        element: <RepositoryPage />,
      },
      {
        path: join_url(route.registry_repository_tag, ':id/:tag'),
        element: <RepositoryTagPage />,
      },
    ],
  },
];
