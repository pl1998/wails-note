<script setup>
import { reactive, computed, onMounted, h, defineProps } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { store } from './../store/index'
import { apiPost, apiGet, apiPut } from './../api/api'
import { ElMessage } from 'element-plus'
const data = reactive({
  name: "",
  text: "Please enter your name below 👇",
  content: "",

})

const props = defineProps({
  name: "",
  menu_id: 1,
  content: "",
  note_id: 0,
  is_show: false,
})

const form = reactive({
  name: "",
  menu_id: 1,
  content: "",
  note_id: 0
})
const menuList = computed(() => store.state.menuList)
function greet() {
  Greet(data.name).then(result => {
    data.text = result
  })
}

function submitForm() {
  if (form.note_id == 0) {
    apiPost('/api/note', form).then((res) => {
      if (res.code == 200) {
        ElMessage({
          message: h('p', null, [
            h('span', null, '保存成功 '),
            h('i', { style: 'color: teal' }, 'VNode'),
          ]),
        })
        form.note_id = res.data.note_id
        apiGet('/api/menu', {})

      } else {

      }
    })
  } else {
    apiPut('/api/note', form.note_id, form).then((res) => {
      if (res.code == 200) {
        ElMessage({
          message: h('p', null, [
            h('span', null, '保存成功 '),
            h('i', { style: 'color: teal' }, 'VNode'),
          ]),
        })
        form.note_id = res.data.note_id
        apiGet('/api/menu', {})

      } else {

      }
    })
  }

}

function handleCopyCodeSuccess(code) {
  console.log(code);
}



const rules = reactive({
  name: [
    { required: true, message: '请输入文章名称', trigger: 'blur' },
    { min: 1, max: 40, message: '文章长度1-40个字符串', trigger: 'blur' },
  ],
  content: [
    { required: true, message: '请输入文章名称', trigger: 'blur' },
  ]
})

</script>
<template>
  <div class="edit-content">
    <v-md-editor v-if="!is_show" :model-value="props.content" mode="preview"
      @copy-code-success="handleCopyCodeSuccess"></v-md-editor>
    <el-form v-else :inline="true" :model="form" class="form-inline" :rules="rules">
      <el-form-item label="文件夹">
        <el-select v-model="form.menu_id" placeholder="选择文件夹" clearable>
          <el-option v-for="menu in menuList" :label="menu.name" :value="menu.menu_id" />
        </el-select>
      </el-form-item>
      <el-form-item label="文章名称">
        <el-input v-model="form.name" placeholder="文章名称" clearable />
      </el-form-item>
      <v-md-editor v-model="form.content" height="550px" @copy-code-success="handleCopyCodeSuccess"></v-md-editor>
      <el-form-item style="margin-top: 20px;">
        <el-button type="primary" @click="submitForm">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<style>
.edit-content {
  padding: 10px;
}

.edit-content .form-inline .el-input {
  --el-input-width: 220px;
}
</style>