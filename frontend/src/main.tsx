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
setConfiguration({
  // TODO: use global variable
  base_url: 'http://localhost:1313',
});

// setup router
// const router = createBrowserRouter([
const router = createHashRouter(routes);

// Render application in DOM
// createRoot(document.getElementById('app')).render(app);
createRoot(document.getElementById('app')).render(
  <RouterProvider router={router} />,
);
