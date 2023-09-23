<script setup>
import EditMarkdown from './components/EditMarkdown.vue'
import { ref, computed, reactive, onMounted } from 'vue'
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { Moon, Sunny, FolderOpened, DocumentCopy } from '@element-plus/icons-vue'
import { store } from './store/index'
import { getMenuList } from './api/menu'
const item = {
  date: '2016-05-02',
  name: 'Tom',
  address: 'No. 189, Grove St, Los Angeles',
}
const theme = computed({
  // getter
  get() {
    return store.state.theme == 'dark' ? false : true
  },
  // setter
  set(value) {
    store.commit('setTheme', value == false ? 'dark' : 'light')
  }
})

const tableData = ref(Array.from({ length: 20 }).fill(item))

const openSetting = ref(false)
const menuList = computed(() => store.state.menuList)

onMounted(() => {
  getMenuList({})
})

</script>

<template>
  <el-container class="layout-container">

    <!-- 菜单栏 -->
    <el-aside width="200px">
      <el-scrollbar>
        <el-menu :default-openeds="['1']">
          <el-sub-menu index="{{menus.menu_id}}" v-for="menus in menuList">
            <template #title>
              <el-icon>
                <FolderOpened />
              </el-icon>{{ menus.name }}
            </template>
            <el-menu-item v-for="note in menus.notes" index="{{note_id}}">
              <el-icon>
                <DocumentCopy />
              </el-icon>{{ note.name }}
            </el-menu-item>
            <el-menu-item-group v-for="menu in menus.children">
              <!-- <template #title>Group 1</template> -->
              <el-menu-item index="{{menu.menu_id}}">
                <el-icon>
                  <FolderOpened />
                </el-icon>{{ menu.name }}</el-menu-item>
              <el-menu-item index="note_{{note.note_id}}" v-for="note in menu.notes"><el-icon>
                  <DocumentCopy />
                </el-icon> {{ note.name }}</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
        </el-menu>
      </el-scrollbar>
    </el-aside>


    <el-container>
      <!-- 表头 -->
      <el-header style="text-align: right; font-size: 12px">
        <div class="theme">
          <el-tooltip :content="'Switch value: ' + theme" placement="top">
            <el-switch v-model="theme" class="mt-2" style="margin-left: 24px" inline-prompt :active-icon="Sunny"
              :inactive-icon="Moon" />
          </el-tooltip>
        </div>
        <div class="toolbar">
          <el-icon style="margin-right: 8px; margin-top: 1px" @click="openSetting = true">
            <setting />
          </el-icon>
          ✨✨✨
        </div>
      </el-header>
      <!-- 内容 -->
      <el-main>
        <el-scrollbar>
          <EditMarkdown />
        </el-scrollbar>
      </el-main>
    </el-container>


    <!-- 设置弹窗 -->
    <el-dialog v-model="openSetting" title="设置" width="70%" z-index="10000px" align-center>
      <div style="height:600px;display:block"></div>
    </el-dialog>

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

