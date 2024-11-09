import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  server: {
    port: 5432
  },
  resolve: {
    alias: {
      '@wailsjs': path.resolve(__dirname, "wailsjs"),
      '@assets': path.resolve(__dirname, "src/assets")
    }
  }
})
