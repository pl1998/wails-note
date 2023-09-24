<script setup>
import {defineProps,defineEmits} from 'vue'
import { FolderAdd, Sunny, Document, FolderDelete } from '@element-plus/icons-vue'
const props = defineProps({
    x:100,
    y:100,
    isShow:false
})
const emit = defineEmits(['isEditDocs','isdelDocs'])

const addDcos = ()=>{
    emit('isEditDocs',false)
}
const delDir = ()=>{
    emit('isdelDocs',false)
}
</script>
<template>
    <el-card v-if="props.isShow" class="box-card" style="transform" :style="{
        left: props.x + 'px',
        top: props.y + 'px',
      }">
      <!-- <template #header>
        <div class="card-header">
          <span>Card name</span>
        
        </div>
      </template> -->
      <div class="text item" @click="addDcos"> <el-icon><Document /></el-icon> 添加文档</div>
      <div class="text item"> <el-icon><FolderAdd /></el-icon> 添加子目录</div>
      <div class="text item" @click="delDir"> <el-icon><FolderDelete /></el-icon> Delete</div>
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
  .el-card__body{
    padding: 10px  !important;
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
  