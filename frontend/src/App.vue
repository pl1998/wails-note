<script setup>
import EditMarkdown from './components/EditMarkdown.vue'
import MenuTree from './components/MenuTree.vue'
import RightMenu from './components/menus/RightMenu.vue'
import { ref, computed, reactive, onMounted, watch, nextTick } from 'vue'
import { Menu as IconMenu, Message, Setting } from '@element-plus/icons-vue'
import { Moon, Sunny, FolderOpened, DocumentCopy, MoreFilled, Plus } from '@element-plus/icons-vue'
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
  }
})
const notes = computed(() => store.state.notes)
const isShowDocs = ref(false)
watch(isShowDocs, (value, old) => {
  console.log('watch更新' + value)

})


const openSetting = ref(false)


const menuList = computed({
  // getter
  get() {
    return store.state.menuList
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

watch(notes, (value, old) => {
  console.log('notes watch更新' + value)
  nextTick()
})


const rightMenu = ref({
  x: 100,
  y: 100,
  isShow: false,
  menu_id: 0
})

onMounted(() => {
  getMenuList({})

})
function openMenu(event, menu_id) {
  event.preventDefault()
  this.rightMenu.x = event.clientX
  this.rightMenu.y = event.clientY
  this.rightMenu.isShow = !this.rightMenu.isShow
  this.rightMenu.menu_id = menu_id
  console.log(this.rightMenu.isShow)
}
function viewDocs(menus) {
  isShowDocs.value = true
  store.commit('setNote', menus)
  console.log('测试查看文章', menus)
  nextTick()
}
const editDocs = (e) => {
  isShowDocs.value = false
  isShowDocs.value = true
  store.commit('setNote', {
    name: "",
    menu_id: 0,
    content: "",
    note_id: 0,
  })
}
function delDocs(id) {
  store.commit('setNote', {})
  store.commit('setIsShowDocs', false)
}

function downMenu(event) {
  event.preventDefault()
  this.rightMenu.isShow = false
}


const onUpdateMenus = (menus) => {
  isShowDocs.value = true
  store.commit('setNote', menus)
  console.log('测试查看文章', menus, isShowDocs)
}

</script>

<template>
  <el-container class="layout-container" @click.left.native="downMenu($event)">
    <!-- 表头 -->
    <!-- <el-header style="text-align: right; font-size: 12px">
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
    </el-header> -->
    <!-- 菜单栏 -->
    <el-container>
      <div>
        <el-aside width="200px">
          <el-scrollbar>
            <el-menu default-active="1" :index="`menu_key` + key" v-for="(menus, key) in menuList">
              <el-sub-menu v-if="menus.is_dir == 1" @click.right.native="openMenu($event, menus.menu_id)"
                @click.left.native="downMenu($event, menus.menu_id)">
                <template #title>
                  <el-icon>
                    <FolderOpened />
                  </el-icon>{{ menus.name }}
                </template>
                <MenuTree v-if="menus.children != null" :menuList="menus.children" @changeMenus="onUpdateMenus" />
              </el-sub-menu>
              <el-menu-item v-else :index="key" @click="viewDocs(menus)"
                @click.right.native="openMenu($event, menus.menu_id)">
                <el-icon>
                  <DocumentCopy />
                </el-icon>{{ menus.name }}
              </el-menu-item>
            </el-menu>
          </el-scrollbar>
        </el-aside>
      </div>
      <el-main>
        <el-scrollbar v-if="isShowDocs">
          <EditMarkdown v-bind="notes" />
        </el-scrollbar>
      </el-main>
    </el-container>


    <!-- 设置弹窗 -->
    <el-dialog v-model="openSetting" title="设置" width="70%" align-center>
      <div style="height:600px;display:block"></div>
    </el-dialog>
    <RightMenu v-if="rightMenu.isShow" :x="rightMenu.x" :y="rightMenu.y" :menuId="rightMenu.menu_id"
      :is-show="rightMenu.isShow" @addDocs="editDocs" />

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

