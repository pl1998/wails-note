import { createApp } from 'vue'
import './styles/reset.less'
import './styles/theme.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// 导入自定义指令
import { vContextmenu } from './directive/contextMenu';

const app = createApp(App)
app.directive('contextmenu', vContextmenu);
app.use(ElementPlus, {
  locale: zhCn,
})
app.use(createPinia())
app.mount('#app')
