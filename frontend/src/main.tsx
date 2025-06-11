import { createRoot } from "react-dom/client";

import { RouterProvider, createHashRouter } from "react-router-dom";

import "bootstrap/dist/css/bootstrap.css";
import "./style.css";
import "bootstrap-icons/font/bootstrap-icons.css";
// import '@fontsource/roboto'; // Defaults to weight 400
import "@fontsource/ubuntu/400.css"; // Defaults to weight 400

import { routes } from "./routes";
import { ConfigurationContext } from "./store/configurationContext";

function App() {
  // setup configuration
  const application_configuration = (window as any).application_configuration || {};

  const configuration = {
    base_url: application_configuration.base_url,
    version: application_configuration.version,
    use_registry: application_configuration.registry,
  };

  // setup router
  const router = createHashRouter(routes);

  return (
    <ConfigurationContext value={configuration}>
      <RouterProvider router={router} />
    </ConfigurationContext>
  );
}

// Render application in DOM
createRoot(document.getElementById("app")!).render(<App />);
