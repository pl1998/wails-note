import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import { prismjsPlugin } from "vite-plugin-prismjs";
import ElementPlus from 'unplugin-element-plus/vite'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    prismjsPlugin({
    languages: ['json','php','javascript','java','go','python'],
  }), 
  ElementPlus({
    // options
  }),
]
})
