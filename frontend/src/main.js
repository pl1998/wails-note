import { createApp } from 'vue'
import App from './App.vue'
import './style.css';


import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';



// highlightjs
import Prism from 'prismjs';

// codemirror 编辑器的相关资源
import Codemirror from 'codemirror';
// mode
import 'codemirror/mode/markdown/markdown';
import 'codemirror/mode/javascript/javascript';
import 'codemirror/mode/css/css';
import 'codemirror/mode/htmlmixed/htmlmixed';
import 'codemirror/mode/vue/vue';
// edit
import 'codemirror/addon/edit/closebrackets';
import 'codemirror/addon/edit/closetag';
import 'codemirror/addon/edit/matchbrackets';
// placeholder
import 'codemirror/addon/display/placeholder';
// active-line
import 'codemirror/addon/selection/active-line';
// scrollbar
import 'codemirror/addon/scroll/simplescrollbars';
import 'codemirror/addon/scroll/simplescrollbars.css';
// style
import 'codemirror/lib/codemirror.css';

// 表情包
import createEmojiPlugin from '@kangc/v-md-editor/lib/plugins/emoji/index';
import '@kangc/v-md-editor/lib/plugins/emoji/emoji.css';

// 代码复制
import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index';
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css';

// tip标签
// import createTipPlugin from '@kangc/v-md-editor/lib/plugins/tip/index';
// import '@kangc/v-md-editor/lib/plugins/tip/tip.css';


import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// import {store} from "./store/index";

//数学公式
// import createKatexPlugin from '@kangc/v-md-editor/lib/plugins/katex/cdn';

// //流程图
// import createMermaidPlugin from '@kangc/v-md-editor/lib/plugins/mermaid/cdn';
// import '@kangc/v-md-editor/lib/plugins/mermaid/mermaid.css';

// //任务列表
// import createTodoListPlugin from '@kangc/v-md-editor/lib/plugins/todo-list/index';
// import '@kangc/v-md-editor/lib/plugins/todo-list/todo-list.css';

//代码行号
import createLineNumbertPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index';

//高亮代码行
import createHighlightLinesPlugin from '@kangc/v-md-editor/lib/plugins/highlight-lines/index';
import '@kangc/v-md-editor/lib/plugins/highlight-lines/highlight-lines.css';

//内容定位
import createAlignPlugin from '@kangc/v-md-editor/lib/plugins/align';

VueMarkdownEditor.Codemirror = Codemirror;
VueMarkdownEditor.use(vuepressTheme, {
  Prism,
  extend(md) {
    // md为 markdown-it 实例，可以在此处进行修改配置,并使用 plugin 进行语法扩展
    // md.set(option).use(plugin);
  },
}).use(createEmojiPlugin()).
// use(createTipPlugin()).
use(createCopyCodePlugin()).
// use(createKatexPlugin()).
// use(createMermaidPlugin()).
// use(createTodoListPlugin()).
use(createLineNumbertPlugin()).
use(createHighlightLinesPlugin()).
use(createAlignPlugin());

createApp(App).use(ElementPlus).use(VueMarkdownEditor).mount('#app')


