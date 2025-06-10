// import path from "path";
import { resolve } from 'path';
// import { fileURLToPath } from "url";
// import { resolve } from "path";
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// const __filename = fileURLToPath(import.meta.url);

// const __dirname = path.dirname(__filename);

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@src': resolve(__dirname, 'src'),
    },
  },
  build: {
    rollupOptions: {
      input: {
        appMain: resolve(__dirname, 'index.html'),
        // appDev: resolve(__dirname, 'index_dev.html'),
      },
    },
  },
});
