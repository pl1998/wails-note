<script setup>
import { ref, reactive, onMounted, onBeforeUnmount, watchEffect } from 'vue'
import {
  FolderOpened,
  DocumentCopy,
  Document,
  FolderAdd,
  Edit,
  FolderDelete,
  DocumentAdd,
} from '@element-plus/icons-vue'
import { ElMessageBox } from 'element-plus'

const props = defineProps({
  menuList: {
    type: Array,
    default: () => [],
  },
  menuInfo: {
    type: Object,
    default: () => {},
  },
})
const noteInfo = ref({})
watchEffect(() => {
  noteInfo.value = props.menuInfo
})
const emit = defineEmits(['noteClick', 'toolClick'])
const noteClick = (data) => {
  noteInfo.value = data
  if (data.is_dir === 1) {
    return
  }
  emit('noteClick', data)
}
const isShowTool = ref(false)
const position = reactive({})
// 菜单详情
const menuInfo = ref({})

const contextmenuClick = (event, data) => {
  position.x = event.clientX
  position.y = event.clientY
  isShowTool.value = true
  menuInfo.value = data
}
const toolClick = (type) => {
  switch (type) {
    case 'create-document':
      ElMessageBox.prompt('请输入文档名称', '温馨提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /^[\s\S]*.*[^\s][\s\S]*$/,
        inputErrorMessage: '名称不能为空',
      })
        .then(({ value }) => {
          emit('toolClick', {
            type: 'create-document',
            p_id:
              menuInfo.value.is_dir === 1
                ? menuInfo.value.menu_id
                : menuInfo.value.p_id,
            name: value,
          })
        })
        .catch(() => {})
      break
    case 'create-catalog':
      ElMessageBox.prompt('请输入目录名称', '温馨提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /^[\s\S]*.*[^\s][\s\S]*$/,
        inputErrorMessage: '名称不能为空',
      })
        .then(({ value }) => {
          emit('toolClick', {
            type: 'create-catalog',
            p_id:
              menuInfo.value.is_dir === 1
                ? menuInfo.value.menu_id
                : menuInfo.value.p_id,
            name: value,
          })
        })
        .catch(() => {})
      break
    case 'edit':
      if (menuInfo.value.is_dir === 1) {
        ElMessageBox.prompt('请输入目录名称', '温馨提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          inputPattern: /^[\s\S]*.*[^\s][\s\S]*$/,
          inputErrorMessage: '名称不能为空',
        })
          .then(({ value }) => {
            emit('toolClick', {
              type: 'edit',
              menuInfo: {
                name: value,
                p_id: menuInfo.value.p_id,
                is_dir: 1,
                content: '',
                menu_id: menuInfo.value.menu_id
              }
            })
          })
          .catch(() => {})
      } else {
        emit('toolClick', {
          type: 'edit',
          menuInfo: menuInfo.value,
        })
      }

      break
    case 'delete':
      emit('toolClick', {
        type: 'delete',
        menuInfo: menuInfo.value,
      })
      break
  }
}
const addClick = (val) => {
  if (val === 'document') {
    ElMessageBox.prompt('请输入文档名称', '温馨提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /^[\s\S]*.*[^\s][\s\S]*$/,
      inputErrorMessage: '名称不能为空',
    })
      .then(({ value }) => {
        emit('toolClick', {
          type: 'create-document',
          p_id: 0,
          name: value,
        })
      })
      .catch(() => {})
  } else {
    ElMessageBox.prompt('请输入目录名称', '温馨提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /^[\s\S]*.*[^\s][\s\S]*$/,
      inputErrorMessage: '名称不能为空',
    })
      .then(({ value }) => {
        emit('toolClick', {
          type: 'create-catalog',
          p_id: 0,
          name: value,
        })
      })
      .catch(() => {})
  }
}
const clearInfo = () => {
  noteInfo.value = {}
}
defineExpose({
  clearInfo,
})
const closeShow = () => {
  isShowTool.value = false
}
onMounted(() => {
  document.addEventListener('click', closeShow)
})
onBeforeUnmount(() => {
  document.removeEventListener('click', closeShow)
})
</script>

<template>
  <div class="operation">
    <div class="icon" title="新增顶级文档" @click="addClick('document')">
      <el-icon><DocumentAdd /></el-icon>
    </div>
    <div class="icon" title="新增顶级目录" @click="addClick('dir')">
      <el-icon><FolderAdd /></el-icon>
    </div>
  </div>
  <div class="tree-wrap">
    <el-tree
      class="tree"
      :data="menuList"
      @node-click="noteClick"
      @node-contextmenu="contextmenuClick"
      default-expand-all
      :props="{
        class: (data, node) => {
          if (data.is_dir === 1) {
            return ''
          }
          return data.menu_id === noteInfo.menu_id ? 'active' : ''
        },
        children: 'children',
        label: 'name',
      }"
    >
      <template #default="{ data }">
        <div class="option">
          <div class="icon" v-if="data && data.is_dir === 1">
            <el-icon>
              <FolderOpened />
            </el-icon>
          </div>
          <div class="icon" v-if="data && data.is_dir !== 1">
            <el-icon>
              <DocumentCopy />
            </el-icon>
          </div>
          <div class="name">{{ data && data.name }}</div>
        </div>
      </template>
    </el-tree>
  </div>
  <div
    class="contextmenu"
    v-if="isShowTool"
    :style="{
      top: position.y + 'px',
      left: position.x + 'px',
    }"
  >
    <ul>
      <li @click="toolClick('create-document')">
        <div class="icon">
          <el-icon><Document /></el-icon>
        </div>
        <div class="name">新建文档</div>
      </li>
      <li @click="toolClick('create-catalog')">
        <div class="icon">
          <el-icon><FolderAdd /></el-icon>
        </div>
        <div class="name">新建目录</div>
      </li>
      <li @click="toolClick('edit')">
        <div class="icon">
          <el-icon><Edit /></el-icon>
        </div>
        <div class="name">编辑</div>
      </li>
      <li @click="toolClick('delete')">
        <div class="icon">
          <el-icon><FolderDelete /></el-icon>
        </div>
        <div class="name">删除</div>
      </li>
    </ul>
  </div>
</template>

<style scoped lang="less">
.operation {
  display: flex;
  width: 100%;
  height: 30px;
  align-items: center;
  justify-content: flex-end;
  box-sizing: border-box;
  padding: 0 10px;
  border-bottom: 1px solid var(--borderColor);
  .icon {
    margin-right: 10px;
    cursor: pointer;
    color: #555;
    &:hover {
      color: #0c64eb;
    }
  }
}
.tree {
  --el-tree-node-content-height: 36px !important;
  .option {
    display: flex;
    align-items: center;
    .icon {
      width: 20px;
      height: 20px;
      display: flex;
      align-items: center;
      margin-right: 5px;
    }
  }
  .active {
    .icon {
      color: #0c64eb;
    }
    .name {
      color: #0c64eb;
    }
  }
}
.contextmenu {
  position: fixed;
  top: 100px;
  left: 50px;
  width: 160px;
  border-radius: 4px 4px 0px 0px;
  box-sizing: border-box;
  box-shadow: 0px 0px 9px 0px rgba(25, 36, 33, 0.28);
  background-color: var(--rightColor);
  box-sizing: border-box;
  padding: 5px 0;
  z-index: 1999;
  ul {
    width: 100%;
    // border: 1px solid #eee;
    // border-bottom: none;
    color: var(--textColor);
    font-size: 14px;
    li {
      width: 100%;
      height: 34px;
      display: flex;
      align-items: center;
      // border-bottom: 1px solid #eee;
      box-sizing: border-box;
      padding: 0 10px;
      cursor: pointer;
      .icon {
        display: flex;
        align-items: center;
        margin-right: 5px;
        font-size: 16px;
      }
      &:hover {
        background-color: var(--rightHoverColor);
      }
    }
  }
}
</style>
