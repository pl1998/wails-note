import { defineConfig, loadEnv } from 'vite'
import path from 'path';
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({ mode, command }) => {
  const env = loadEnv(mode, './');
  return {
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
      }
    },
    server: {
      https: false,
      port: env.VITE_PORT,
      host: '0.0.0.0',
      proxy: {
        '/dev': {
          target: env.VITE_BASE_URL,
          changeOrigin: true,
          rewrite: path => path.replace(/^\/dev/, '')
        }
      }
    },
    plugins: [
      vue()
    ]
  };
})
