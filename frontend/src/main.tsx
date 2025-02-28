import React from 'react';
import { createRoot } from 'react-dom/client';

import { RouterProvider, createHashRouter } from 'react-router-dom';

// import 'chota';
// import 'bulma/css/bulma.css';
import 'bootstrap/dist/css/bootstrap.css';
import './style.css';
import 'bootstrap-icons/font/bootstrap-icons.css';
// import '@fontsource/roboto'; // Defaults to weight 400
import '@fontsource/ubuntu'; // Defaults to weight 400

import { routes } from './routes';
import { useConfiguration } from './store/configuration';

// setup configuration
const { setConfiguration } = useConfiguration();
const application_configuration = (window as any).application_configuration || {};

setConfiguration({
  base_url: application_configuration.base_url,
  version: application_configuration.version,
  use_registry: application_configuration.registry,
});

// setup router
// const router = createBrowserRouter([
const router = createHashRouter(routes);

// Render application in DOM
// createRoot(document.getElementById('app')).render(app);
createRoot(document.getElementById('app')).render(<RouterProvider router={router} />);
