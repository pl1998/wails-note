
<script  setup>
import { reactive, defineEmits, defineProps } from 'vue'
import { addMenu,getMenuList } from './../api/menu'
import { ElMessage } from 'element-plus'
const props = defineProps({
    centerDialogVisible: false,
    pId:0
})

const form = reactive({
  name: "",
  p_id: props.pId,
  is_dir: 1,
  content: ""
})

const emits = defineEmits(['onSubmit'])
function onSubmit() {
  console.log('提交')
  console.log(props.pId)
  console.log(form.p_id)
    addMenu(form).then((res) => {
        if (res.code == 200) {
            getMenuList({})
            ElMessage({
                message: '创建成功~',
                type: 'success'
            })
        } else {
            ElMessage({
                message: '创建失败~',
                type: 'error'
            })
        }
    })
    emits('onSubmit')
}

</script>
<template>
    <el-dialog v-model="props.centerDialogVisible" title="创建文件夹" width="30%" center>
        <el-form :model="form" label-width="120px">
            <el-form-item label="文件夹名称">
                <el-input v-model="form.name" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">创建</el-button>
            </el-form-item>
        </el-form>
    </el-dialog>
</template>
<style scoped>
.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
  