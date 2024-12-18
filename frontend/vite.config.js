import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        AntDesignVueResolver({importStyle: false})
      ]
    })
  ],
  server: {
    port: 5432
  },
  resolve: {
    alias: {
      '@wailsjs': path.resolve(__dirname, "wailsjs"),
      '@assets': path.resolve(__dirname, "src/assets"),
      '@utils': path.resolve(__dirname, "src/utils"),
      "@": path.resolve(__dirname, "src"),
    }
  }
})
