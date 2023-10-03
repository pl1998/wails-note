<script setup>
import MdEditorV3 from './components/MdEditorV3.vue'
import MenuTree from './components/MenuTree.vue'
import CreateFile from './components/CreateFile.vue'
import RightMenu from './components/menus/RightMenu.vue'
import { ref, onMounted, watch, nextTick ,computed} from 'vue'
import {  Setting } from '@element-plus/icons-vue'
import { Moon, Sunny, FolderOpened, DocumentCopy,MoreFilled } from '@element-plus/icons-vue'
import { store } from './store/index'
import { getMenuList, deleleMenu } from './api/menu'
import { useDark, useToggle } from '@vueuse/core'
const isDark = useDark()
const toggleDark = useToggle(isDark)
const notes = computed(() => store.state.notes ?? {
  name: "",
  menu_id: 0,
  content: "# Hello Editor",
  note_id: 0,
  is_dir: 0,
})
const isShowDocs = ref(false)
const isEditDocs = ref(false)
const openSetting = ref(false)
const centerDialogVisible = ref(false)
const isLeftMenu = ref(false)

const menuList = computed({
  // getter
  get() {
    return store.state.menuList ?? []
  },
  // setter
  set(value) {
    nextTick()
    store.commit('setMenuList', value)
  }
})
watch(menuList, (value, old) => {
  console.log('menuList watch更新' + value)
  nextTick()
})

const rightMenu = ref({
  x: 100,
  y: 100,
  isShow: false,
  menu_id: 0,
  menus: {}
})

onMounted(() => {
  getMenuList({})
})
function createDocs() {
  centerDialogVisible.value = true
}

function clickMenu(event,menu) {
  // console.log(event)
  event.preventDefault()
  this.rightMenu.x = event.clientX
  this.rightMenu.y = event.clientY
  isLeftMenu.value = true
  this.rightMenu.menu_id = menu.menu_id
  this.rightMenu.menus = menu
  this.rightMenu.isShow = true
  console.log(this.rightMenu)
}

function openMenu(event, menu) {
  event.preventDefault()
  this.rightMenu.x = event.clientX
  this.rightMenu.y = event.clientY
  isLeftMenu.value = true
  this.rightMenu.menu_id = menu.menu_id
  this.rightMenu.menus = menu
}
function viewDocs(menus) {
  console.log('viewDocs')
  isEditDocs.value = false
  isShowDocs.value = false
  isShowDocs.value = true
  openIsLeftMenu()
  store.commit('setNote', menus)
}
const addDocs = (e) => {

  console.log(e.menu_id)
  isShowDocs.value = false
  isShowDocs.value = true
  openIsLeftMenu()
  store.commit('setNote', {
    name: "",
    menu_id: 0,
    content: "",
    note_id: 0,
    p_id:e.menu_id,
    is_dir:0,
    menus:e
  })
  isEditDocs.value = true
}
function delDocs(id) {
  store.commit('setNote', {})
  store.commit('setIsShowDocs', false)
}

function downMenu(event) {
  isLeftMenu.value = false

}

function openIsLeftMenu() {
  event.preventDefault()
  isLeftMenu.value = !isLeftMenu.value
  console.log(isLeftMenu.value)
}

const downCenterDialogVisible = () => {
  centerDialogVisible.value = false
}

const onUpdateMenus = (menus) => {
  isShowDocs.value = true
  store.commit('setNote', menus)
}

const delMenu = (value) => {
  openIsLeftMenu()
  deleleMenu(value)
}
const addDir = () => {
  centerDialogVisible.value = true
  openIsLeftMenu()
}
const editDocs = (menus) => {
  isShowDocs.value = false
  isShowDocs.value = true
  isEditDocs.value = true
  rightMenu.menus = menus
  console.log(menus)
  store.commit('setNote', menus)
  openIsLeftMenu()
}

</script>

<template>
  <el-container class="layout-container" @click.left.native="downMenu($event)">
    <!-- 表头 -->
    <el-header style="text-align: right; font-size: 12px">
      <div class="theme">
        <el-switch v-model="isDark" :active-icon="Sunny" :inactive-icon="Moon" class="mt-2"
          inline-prompt size="large" style="--el-switch-off-color: #73767a;--el-switch-on-color: #73767a"
          @click="!isDark" />
      </div>
      <div class="toolbar">
        <el-icon style="margin-right: 8px; margin-top: 1px" @click="openSetting = true">
          <setting />
        </el-icon>
      </div>
    </el-header>
    <!-- 菜单栏 -->
    <el-container>
      <div>
        <el-aside v-if="menuList != null" width="200px">
          <el-scrollbar>
            <el-menu v-for="(menus, key) in menuList">
              <el-sub-menu v-if="menus.is_dir == 1" :index="`menu` + key" @click.right.native="openMenu($event, menus)"
                @click.left.native="downMenu($event, menus.menu_id)">
                <template #title>
                  <el-icon>
                    <FolderOpened />
                  </el-icon>{{ menus.name }}
                 <div style="display: flex;justify-content: flex-end;height: 100%;width: 100%">
                   <el-icon @click="clickMenu($event,menus)" style="height: 100%">
                     <MoreFilled />
                   </el-icon>
                 </div>
                </template>
                <MenuTree v-if="menus.children != null" :menuList="menus.children" @changeMenus="onUpdateMenus" />
              </el-sub-menu>
              <el-menu-item v-else :index="`menu` + key" @click="viewDocs(menus)"
                @click.right.native="openMenu($event, menus)">
                <el-icon>
                  <DocumentCopy />
                </el-icon>{{ menus.name }}
              </el-menu-item>
            </el-menu>
          </el-scrollbar>
        </el-aside>
        <el-empty v-if="menuList.length == 0">
          <el-button type="primary" @click="createDocs">创建文档</el-button>
        </el-empty>
      </div>
      <el-main>

        <el-scrollbar v-if="isShowDocs">
          <!-- <EditMarkdown v-bind="notes" /> -->
          <MdEditorV3 :isEditDocs="isEditDocs" v-bind="notes" />

        </el-scrollbar>
      </el-main>
    </el-container>
    <!-- 设置弹窗 -->
    <el-dialog v-model="openSetting" align-center title="设置" width="70%">
      <div style="height:600px;display:block"></div>
    </el-dialog>
    <RightMenu v-if="isLeftMenu" :is-show="isLeftMenu" :menuId="rightMenu.menu_id" :menus="rightMenu.menus" :x="rightMenu.x"
      :y="rightMenu.y" @addDir="addDir" @addDocs="addDocs" @delMenu="delMenu" @editDocs="editDocs" />
    <CreateFile :centerDialogVisible="centerDialogVisible" @onSubmit="downCenterDialogVisible" />
  </el-container>
</template>


<style scoped>
.layout-container {
  height: 100%;
}

.layout-container .el-header {
  position: relative;
}

.layout-container .el-aside {}

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

.layout-container .theme .el-switch.is-checked .el-switch__core {}

.el-aside .el-scrollbar__view .el-scrollbar__view ui .el-sub-menu__title {}
</style>

