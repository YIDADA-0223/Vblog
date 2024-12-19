import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), vueDevTools()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      // string shorthand: http://localhost:5173/vblog/api/v1 -> http://l27.0.0.1:8080/vblog/api/v1
      // 访问URL变成 http://localhost:5173/vblog/api/v1 不能访问BaseURL
      '/vblog/api/v1/': 'http://localhost:8080'
    }
  }
})
