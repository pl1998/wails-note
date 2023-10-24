<script setup>
import EditMarkdown from './components/EditMarkdown.vue'
import MenuTree from './components/MenuTree.vue'
import RightMenu from './components/menus/RightMenu.vue'
import { ref, computed, reactive, onMounted } from 'vue'
import { Setting } from '@element-plus/icons-vue'
import {
  Moon,
  Sunny,
} from '@element-plus/icons-vue'
import { store } from './store/index'
import { getMenuList } from './api/menu'

const theme = computed({
  // getter
  get() {
    return store.state.theme == 'dark' ? false : true
  },
  // setter
  set(value) {
    store.commit('setTheme', value == false ? 'dark' : 'light')
  },
})

const notes = computed(() => store.state.notes)
const isShowDocs = ref(false)

const openSetting = ref(false)

const menuList = computed({
  get() {
    console.log('memnlist', store.state.menuList)
    return store.state.menuList
  },
  set(value) {
    store.commit('setMenuList', value)
  },
})

const rightMenu = reactive({
  x: 100,
  y: 100,
  isShow: false,
})

onMounted(() => {
  getMenuList({})
  document.addEventListener('click', (e) => {
    rightMenu.isShow = false
  })
})
const openMenu = (val) => {
  rightMenu.x = val.x
  rightMenu.y = val.y
  rightMenu.isShow = true
}
const menuClick = (menu) => {
  console.log(menu)
  isShowDocs.value = false
  store.commit('setNote', menu)
}
</script>

<template>
  <el-container class="layout-container">
    <!-- 表头 -->
    <el-header style="text-align: right; font-size: 12px">
      <div class="theme">
        <el-tooltip :content="'Switch value: ' + theme" placement="top">
          <el-switch
            v-model="theme"
            class="mt-2"
            style="margin-left: 24px"
            inline-prompt
            :active-icon="Sunny"
            :inactive-icon="Moon"
          />
        </el-tooltip>
      </div>
      <div class="toolbar">
        <el-icon
          style="margin-right: 8px; margin-top: 1px"
          @click="openSetting = true"
        >
          <setting />
        </el-icon>
        ✨✨✨
      </div>
    </el-header>
    <!-- 菜单栏 -->
    <el-container>
      <div>
        <el-aside width="200px">
          <el-scrollbar>
            <el-menu default-active="1">
              <MenuTree
                :menuList="menuList"
                @menuClick="menuClick"
                @openMenu="openMenu"
              />
            </el-menu>
          </el-scrollbar>
        </el-aside>
      </div>
      <el-main>
        <el-scrollbar>
          <EditMarkdown v-bind="notes" :is-show="isShowDocs" />
        </el-scrollbar>
      </el-main>
    </el-container>

    <!-- 设置弹窗 -->
    <el-dialog v-model="openSetting" title="设置" width="70%" align-center>
      <div style="height: 600px; display: block"></div>
    </el-dialog>
    <RightMenu
      :x="rightMenu.x"
      :y="rightMenu.y"
      :is-show="rightMenu.isShow"
      @isEditDocs="!isShowDocs"
    />
  </el-container>
</template>

<style scoped>
.layout-container {
  height: 90%;
}

.layout-container .el-header {
  position: relative;
  background-color: var(--el-color-info-light-9);
  color: var(--el-text-color-primary);
}

.layout-container .el-aside {
  color: var(--el-text-color-primary);
  background: var(--el-color-white);
}

.layout-container .el-menu {
  border-right: none;
}

.layout-container .el-main {
  padding: 0;
}

.layout-container .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
  margin-left: 20px;
}

.layout-container .theme {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}

.layout-container .theme .el-switch.is-checked .el-switch__core {
  background: var(--el-bg-color) !important;
  border-color: none !important;
}

.el-aside .el-scrollbar__view .el-scrollbar__view ui .el-sub-menu__title {
  font-size: 12px !important;
  font-weight: bold !important;
}
</style>
