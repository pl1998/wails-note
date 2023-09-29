import { createApp } from 'vue'
import App from './App.vue'
import './style.css';

import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/display.css'

import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import 'md-editor-v3/lib/preview.css';


createApp(App).
use(ElementPlus).
use(MdEditor).
mount('#app')


