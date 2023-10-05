<script setup>
import { defineProps, defineEmits } from 'vue';
import { FolderOpened, DocumentCopy } from '@element-plus/icons-vue'
import {store} from "../store/index.js";
const props = defineProps({
    menuList: [] | undefined,
  isShowDocs:false
})
const emits = defineEmits(['changeMenus', 'isdelDocs'])
function changeMenus(menus) {
    emits('changeMenus', menus)
    console.log('选中')
}
const onUpdateMenus = (menus) => {
  props.isShowDocs = true
  store.commit('setNote', menus)
}
</script>
<template>
    <div class="custom-menu-tree">
        <template v-for="(menus, key) in props.menuList">
            <el-sub-menu v-if="menus.is_dir == 1" :index="`sub_menu`+key">
                <template #title>
                    <el-icon>
                        <FolderOpened />
                    </el-icon>{{ menus.name }}
                </template>
                <MenuTree v-if="menus.children != null" :menuList="menus.children" @changeMenus="onUpdateMenus"  :isShowDocs="props.isShowDocs "/>
            </el-sub-menu>
            <el-menu-item v-else :index="`sub_menu`+key" @click="changeMenus(menus)">
                <el-icon>
                    <DocumentCopy />
                </el-icon>{{ menus.name }}
            </el-menu-item>

        </template>
    </div>
</template>