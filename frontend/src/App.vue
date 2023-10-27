<script setup>
import { ref, computed } from 'vue'
import { getMenuList, addMenu, updateMenu, deleleMenu } from '@/api/menu'
import MenuList from './components/MenuList.vue'
import MyEditor from './components/MyEditor.vue'
import { ElMessage } from 'element-plus'
import { useBaseInfo } from './store'

const baseInfo = useBaseInfo()
const theme = computed(() => baseInfo.theme)
window.oncontextmenu = function (e) {
  //取消默认的浏览器自带右键 很重要！！
  e.preventDefault()
}
const menuData = ref([])
const getMenuListFn = async () => {
  const res = await getMenuList()
  menuData.value = res.data
  console.log(menuData.value)
}
getMenuListFn()
// 新增
const addMenuClick = async (params) => {
  const res = await addMenu(params)
  ElMessage.success('新增成功')
  console.log(res)
  noteInfo.value = res.data
  getMenuListFn()
  isEdit.value = false
}
// 编辑
const updateMenuClick = async (params) => {
  const res = await updateMenu(params)
  console.log(res)
  ElMessage.success('编辑成功')
  getMenuListFn()
  isEdit.value = false
}
// 删除
const deleteMenuClick = async (params) => {
  await deleleMenu(params)
  ElMessage.success('删除成功')
  getMenuListFn()
}
const noteInfo = ref({})
const isEdit = ref(false)
const noteClick = (menu) => {
  isEdit.value = false
  noteInfo.value = menu
}
const toolClick = (params) => {
  const { type, p_id, name, menuInfo } = params
  switch (type) {
    case 'create-document':
      noteInfo.value = {
        content: '',
        name,
        p_id,
        type,
      }
      isEdit.value = true
      break
    case 'create-catalog':
      const params = {
        name: name,
        p_id: p_id,
        is_dir: 1,
        content: '',
      }
      addMenuClick(params)
      break
    case 'edit':
      if (menuInfo.is_dir === 1) {
        updateMenuClick(menuInfo)
        return
      }
      noteInfo.value = menuInfo
      noteInfo.value.type = 'edit'
      isEdit.value = true
      break
    case 'delete':
      if (menuInfo.menu_id === noteInfo.value.menu_id && isEdit.value) {
        return ElMessage.warning('文档编辑中，无法删除！')
      } else {
        deleteMenuClick(menuInfo.menu_id)
        if (menuInfo.menu_id === noteInfo.value.menu_id) {
          noteInfo.value = {}
        }
      }

      break
  }
}

const saveClick = (val) => {
  const type = noteInfo.value.type
  switch (type) {
    case 'create-document':
      const params = {
        name: val.name,
        p_id: noteInfo.value.p_id,
        is_dir: 0,
        content: val.content,
      }
      addMenuClick(params)
      break
    case 'edit':
      const info = {
        name: val.name,
        p_id: noteInfo.value.p_id,
        is_dir: 0,
        content: val.content,
        menu_id: noteInfo.value.menu_id,
        note_id: noteInfo.value.note_id || 0,
      }
      updateMenuClick(info)
      break
  }
}
const themeType = ref(false)
const themeRes = localStorage.getItem('theme') || 'light'
themeType.value = themeRes === 'light'
const themeChange = () => {
  const type = themeType.value ? 'light' : 'dark'
  baseInfo.changeTheme(type)
}
const showMenu = computed(() => baseInfo.showMenu)
const showAside = computed(() => baseInfo.showAside)
const showMenuClick = () => {
  baseInfo.changeShowMenu()
}
const showAsideClick = () => {
  baseInfo.changeShowAside()
}
</script>

<template>
  <div class="container" :class="theme">
    <div class="head">
      <div class="tool">
        <div class="item" @click="showAsideClick">
          {{ showAside ? '隐藏侧边栏' : '显示侧边栏' }}
        </div>
        <div class="item" @click="showMenuClick">
          {{ showMenu ? '隐藏菜单' : '显示菜单' }}
        </div>
      </div>
      <el-switch
        v-model="themeType"
        class="ml-2"
        @change="themeChange"
        style="--el-switch-on-color: #13ce66; --el-switch-off-color: #73767a"
      />
    </div>
    <div class="content">
      <div class="aside" v-if="showAside">
        <MenuList
          ref="menuRef"
          v-if="menuData.length"
          :menu-list="menuData"
          :menuInfo="noteInfo"
          @noteClick="noteClick"
          @toolClick="toolClick"
        ></MenuList>
        <el-empty v-if="menuData.length === 0"> </el-empty>
      </div>
      <div class="main">
        <MyEditor
          :content="noteInfo.content"
          :isEdit="isEdit"
          :title="noteInfo.name"
          @saveClick="saveClick"
        ></MyEditor>
      </div>
    </div>
  </div>
</template>

<style scoped lang="less">
.container {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--mainColor);

  .head {
    width: 100%;
    height: 60px;
    border-bottom: 1px solid var(--borderColor);
    display: flex;
    align-items: center;
    justify-content: flex-end;
    box-sizing: border-box;
    padding: 0 20px;
    .tool {
      display: flex;
      align-items: center;
      font-size: 14px;
      margin-right: 20px;
      .item {
        padding: 0 20px;
        border-right: 1px solid var(--borderColor);
        cursor: pointer;
        color: var(--textColor);
      }
    }
  }
  .content {
    width: 100%;
    flex: 1;
    overflow: hidden;
    display: flex;
    .aside {
      width: 250px;
      height: 100%;
      overflow-y: auto;
      border-right: 1px solid var(--borderColor);
      // background-color: pink;
    }
    .main {
      flex: 1;
      overflow: hidden;
      // background-color: yellow;
    }
  }
}
</style>
