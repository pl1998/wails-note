<script setup>
import { defineProps, defineEmits } from 'vue'
import { FolderAdd, Edit, Document, FolderDelete } from '@element-plus/icons-vue'
import { store } from './../../store/index'
const props = defineProps({
  x: 100,
  y: 100,
  isShow: false,
  menuId:0,
  menus:{}
})

const emits = defineEmits(['addDocs', 'addDir','delMenu','editDocs'])
const addDocs = () => {
  console.log("addDocs",props.menuId)
  emits('addDocs',props.menus)
}
const addDir = () => {
  console.log("addDir",props.menuId)
  emits('addDir',true)
}
const delMenu = () => {
  console.log("delMenu",props.menuId)
  emits('delMenu',props.menuId)
}
const editDocs = () => {
  console.log("editDocs",props.menuId)
  emits('editDocs',props.menus)
}
</script>
<template>
  <el-card v-if="props.isShow" class="box-card" :style="{
    left: props.x + 'px',
    top: props.y + 'px',
  }">
    <div class="text item" @click="addDocs"> <el-icon>
        <Document />
      </el-icon> 添加文档</div>
    <div class="text item" @click="addDir"> <el-icon>
        <FolderAdd />
      </el-icon> 添加子目录</div>
      <div class="text item" @click="editDocs"> <el-icon>
        <el-icon><Edit /></el-icon>
      </el-icon> 编辑</div>
    <div class="text item" @click="delMenu"> <el-icon>
        <FolderDelete />
      </el-icon> 删除</div>
  </el-card>
</template>
  
<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 13px;
  font-weight: 400;
}

.el-card__body {
  padding: 10px !important;
  z-index: 1 !important;
}

.item {
  line-height: 19px;
  padding: 10px 2px 10px 2px;
  display: block;
}

.text:hover {
  background: var(--el-color-info-light-7);
}

.box-card {
  width: 180px;
  position: absolute;
}
</style>
  