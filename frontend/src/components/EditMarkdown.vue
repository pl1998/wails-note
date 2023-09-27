<script setup>
import { reactive, computed, onMounted, h, defineProps, watch } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { store } from './../store/index'
import { apiPost, apiGet, apiPut } from './../api/api'
import { getMenuList } from './../api/menu'

import { ElMessage } from 'element-plus'
const data = reactive({
  name: "",
  text: "Please enter your name below 👇",
  content: "",

})


const props = defineProps({
  name: "",
  menu_id: 0,
  content: "",
  note_id: 0,
  is_dir: 0,
})
const form = reactive({
  name: "",
  menu_id: 1,
  content: "",
  note_id: 0,
  p_id: props.menu_id,
  is_dir: props.is_dir
})
const menuList = computed(() => store.state.menuList)
function greet() {
  Greet(data.name).then(result => {
    data.text = result
  })
}
watch(props.note_id, (val, old) => {
  console.log("子组件" + val, old)
})


function submitForm() {
  if (form.note_id == 0) {
    if (form.is_dir == 0) {
      form.p_id = 0
    }
    apiPost('/api/menu', form).then((res) => {
      if (res.code == 200) {
        ElMessage({
          message: '添加成功~',
          type: 'success',
        })
        form.note_id = res.data.note_id
        form.menu_id = res.data.menu_id
        form.content = res.data.content
        form.name = res.data.name
        getMenuList({}).then((res) => {
          if (res.code == 200) {
            store.commit('setMenuList', res.data)
            console.log("更新数据成功")
          }
        })
      } else {
        ElMessage({
          message: '添加失败',
          type: 'error',
        })
      }
    })
  } else {
    apiPut('/api/menu', form.note_id, form).then((res) => {
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

function openEdit() {
  form.content = props.content
  form.note_id = props.note_id
  form.name = props.name
  form.content = props.content
  form.menu_id = props.menu_id
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
    <el-button type="primary" v-show="props.menu_id == 0" @click="openEdit">编辑</el-button>
    <div>
      <v-md-editor v-if="props.menu_id != 0" :include-level="[3, 4]" :model-value="props.content"
        mode="preview"></v-md-editor>
    </div>
    <el-form v-if="props.menu_id == 0" :inline="true" :model="form" class="form-inline" :rules="rules">
      <el-form-item label="文件夹">
        <el-select v-model="form.p_id" placeholder="选择文件夹" clearable>
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