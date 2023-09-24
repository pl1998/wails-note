<script setup>
import EditMarkdown from './components/EditMarkdown.vue'
import MenuTree from './components/MenuTree.vue'
import RightMenu from './components/menus/RightMenu.vue'
import { contextMenu } from './components/menus/index'
import { ref, computed, reactive, onMounted } from 'vue'
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { Moon, Sunny, FolderOpened, DocumentCopy, MoreFilled, Plus } from '@element-plus/icons-vue'
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

const notes = computed(() => store.state.notes)
const isShowDocs =ref(false)

const openSetting = ref(false)

const menuList = computed({
  // getter
  get() {
    console.log('memnlist')
    return  store.state.menuList
  },
  // setter
  set(value) {
    store.commit('setMenuList',value)
  }
})

const rightMenu = ref({
  x: 100,
  y: 100,
  isShow: false
})

onMounted(() => {
  getMenuList({})
})
function openMenu(event, item) {
  this.rightMenu.x = event.clientX
  this.rightMenu.y = event.clientY
  event.preventDefault()
  this.rightMenu.isShow = !this.rightMenu.isShow
}
function viewDocs(item) {
  this.isShowDocs = true
  store.commit('setNote', item)
  console.log(this.isShowDocs)
  console.log(item)
}

function downMenu(event) {
  event.preventDefault()
  this.isShowDocs = false

}

</script>

<template>
  <el-container class="layout-container">
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
    <!-- 菜单栏 -->
    <el-container>
      <div>
        <el-aside width="200px">
          <el-scrollbar>
            <el-menu default-active="1">
              <el-sub-menu :index="`menu_key`+key" v-for="(menus,key) in menuList"
                @click.right.native="openMenu($event, menus.menu_id)"
                @click.left.native="downMenu($event, menus.menu_id)">
                <template #title>
                  <el-icon>
                    <FolderOpened />
                  </el-icon>{{ menus.name }}
                </template>
                <el-menu-item v-for="(note,key) in menus.notes" :index="`note_key`+key" @click="viewDocs(note)">
                  <el-icon>
                    <DocumentCopy />
                  </el-icon>{{ note.name }}
                </el-menu-item>
                <MenuTree :menuList="menus.children == null ?? []" />
              </el-sub-menu>
            </el-menu>
          </el-scrollbar>
        </el-aside>
      </div>
      <el-main>
        <el-scrollbar>
          <EditMarkdown v-bind="notes"  :is-show="isShowDocs" />
        </el-scrollbar>
      </el-main>
    </el-container>


    <!-- 设置弹窗 -->
    <el-dialog v-model="openSetting" title="设置" width="70%"  align-center>
      <div style="height:600px;display:block"></div>
    </el-dialog>
    <RightMenu :x="rightMenu.x" :y="rightMenu.y" :is-show="rightMenu.isShow" @isEditDocs="!isShowDocs" />

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

