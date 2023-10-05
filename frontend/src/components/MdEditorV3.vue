  
<script setup>
import { ref, defineProps, reactive, computed, watch } from 'vue';
import { MdEditor, MdPreview, MdCatalog } from 'md-editor-v3';
import { store } from './../store/index'
import { getDirTree } from './../library/func'
import { getMenuList, addMenu, updateMenu } from './../api/menu'
import 'md-editor-v3/lib/preview.css';
import { ElMessage } from 'element-plus'
const props = defineProps({
    name: "",
    menu_id: 0,
    p_id: 0,
    content: "# Hello Editor",
    note_id: 0,
    is_dir: 0,
    is_show: false,
    isEditDocs: false
})

const id = 'preview-only';
const text = ref('# Hello Editor');
const scrollElement = document.documentElement;
const isPreview = ref(false)
const options = ref({
  expandTrigger: 'click' | 'hover',
  multiple: false,
  checkStrictly: true,
  value: 'menu_id',
  label: 'name'
})

const form = reactive({
    name: props.name,
    menu_id: props.menu_id,
    content: props.content ?? "# hello world",
    note_id: props.note_id ?? 0,
    p_id: props.menu_id,
    is_dir: props.is_dir
})

const menuList = computed(() => {
     return getDirTree(store.state.menuList ?? [])
})
import { useDark } from '@vueuse/core'
const isDark = useDark()

const rules = reactive({
    name: [
        { required: true, message: '请输入文章名称', trigger: 'blur' },
        { min: 1, max: 40, message: '文章长度1-40个字符串', trigger: 'blur' },
    ],
    content: [
        { required: true, message: '请输入文章名称', trigger: 'blur' },
    ]
})
function  changeForm(data){
  form.p_id = data.at(data.length-1)
  console.log(form.p_id)
}
function onSubmit() {
    if (form.menu_id == 0 || form.menu_id == undefined || form.menu_id == null) {

        addMenu(form).then((res) => {
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
        updateMenu(form)
    }

}


</script>
<template>
    <!-- 编辑 -->
    <div class="edit-content">
        <el-form v-if="props.isEditDocs || props.menu_id == 0" :inline="true" :model="form" class="form-inline"
            :rules="rules">
            <el-form-item label="文件夹">
              <el-cascader :options="menuList" :show-all-levels="false" v-model="form.p_id" :props="options" @change="changeForm"  />
<!--                <el-select v-model="form.p_id" placeholder="选择文件夹" clearable>-->
<!--                    <el-option v-for="menu in menuList" :label="menu.name" :value="menu.menu_id" :selected="menu.menu_id==form.p_id ? `selected` : ``"   />-->
<!--                </el-select>-->
            </el-form-item>
            <el-form-item label="文章名称">
                <el-input v-model="form.name" placeholder="文章名称" clearable />
            </el-form-item>
            <MdEditor :preview="isPreview" v-model="form.content" :theme="!isDark ? `light` : `dark`" />

            <el-button type="primary" @click="onSubmit">{{ props.isEditDocs ? 'Save' : 'Create' }}</el-button>
        </el-form>

        <!-- 预览 -->
        <div v-else>
            <div class="common-layout">
                <el-container>
                    <el-aside width="200px;">
                        <MdCatalog :editorId="id" :modelValue="props.content" :theme="!isDark ? `light` : `dark`" />
                    </el-aside>
                    <el-main>
                        <MdPreview showCodeRowNumber="true" :editorId="id" :modelValue="props.content"
                            :theme="!isDark ? `light` : `dark`" :scrollElement="scrollElement" />
                    </el-main>
                </el-container>
            </div>
        </div>

    </div>
</template>
<style scoped>
/deep/ .el-aside .md-editor-catalog-link {
    max-height: 1200px;
    font-size: 14px;
    overflow: auto !important;
    color: #000;
}

/deep/ .el-aside .md-editor-catalog-link ::-webkit-scrollbar {
    width: 5px;
    height: 5px;
    background-color: transparent;
}

.edit-content {
    .md-editor-previewOnly {
        max-height: 1200px;
    }
}

.md-editor-catalog-active>span {
    color: #000 !important;
}

.md-editor {
    max-height: 1200px;

}

#md-editor-v3 {
    height: 600px;
}

.md-editor-previewOnly {
    max-height: 1200px;

}

.el-main {
    --el-main-padding: 0px !important;
}
</style>
