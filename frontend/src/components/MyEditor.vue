<script setup>
import { ref, watchEffect, computed } from 'vue'
import { MdEditor, MdPreview, MdCatalog } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { ElMessage } from 'element-plus'
import { useBaseInfo } from '@/store'

const baseInfo = useBaseInfo()
const theme = computed(() => baseInfo.theme)
const showMenu = computed(() => baseInfo.showMenu)
const props = defineProps({
  content: {
    type: String,
    default: '',
  },
  isEdit: {
    type: Boolean,
    default: false,
  },
  title: {
    type: String,
    default: '',
  },
})
const noteText = ref('')
const noteTitle = ref('')
watchEffect(() => {
  noteText.value = props.content
  noteTitle.value = props.title
})

const id = 'preview-only'
const scrollElement = document.documentElement
const emit = defineEmits(['saveClick'])
const saveClick = () => {
  console.log(noteText.value)
  if (!noteText.value || !noteTitle.value) {
    return ElMessage.warning('名称和内容不能为空！')
  }
  emit('saveClick', {
    content: noteText.value,
    name: noteTitle.value,
  })
}
</script>

<template>
  <div class="my-editor">
    <div class="edit" v-if="isEdit">
      <div class="head">
        <div class="title">文档名称：</div>
        <div class="input">
          <el-input v-model="noteTitle" clearable placeholder="请输入" />
        </div>
      </div>
      <div class="content">
        <MdEditor v-model="noteText" :theme="theme" />
      </div>
      <div class="operation">
        <div class="btn" @click="saveClick">
          <el-button type="primary" color="#666" style="color: #fff">
            保存
          </el-button>
        </div>
      </div>
    </div>
    <div class="view" v-else>
      <div class="left" v-show="showMenu">
        <MdCatalog :editorId="id" :theme="theme" :modelValue="noteText" />
      </div>
      <div class="right">
        <MdPreview
          :editorId="id"
          :theme="theme"
          :modelValue="noteText"
          :scrollElement="scrollElement"
        />
      </div>
    </div>
  </div>
</template>

<style scoped lang="less">
.my-editor {
  width: 100%;
  height: 100%;
  .edit {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    .head {
      width: 100%;
      height: 40px;
      display: flex;
      align-items: center;
      background-color: #eee;
      color: #333;
      box-sizing: border-box;
      padding: 0 10px;
      .title {
        font-size: 14px;
      }
    }
    .content {
      flex: 1;
      overflow: hidden;
      :deep(.md-editor) {
        height: 100%;
      }
    }
    .operation {
      width: 100%;
      height: 80px;
      align-items: center;
      display: flex;
      justify-content: flex-end;
      box-sizing: border-box;
      padding: 0 20px 20px;
    }
  }
  .view {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: space-between;
    .left {
      width: 260px;
      box-sizing: border-box;
      padding: 10px;
      // overflow-y: auto;
      border-right: 1px solid var(--borderColor);
      font-size: 14px;
      :deep(.md-editor-catalog) {
        height: 100%;
      }
    }
    .right {
      flex: 1;
      box-sizing: border-box;
      padding: 10px 5px;
      :deep(.md-editor) {
        height: 100%;
      }
    }
  }
}
</style>
