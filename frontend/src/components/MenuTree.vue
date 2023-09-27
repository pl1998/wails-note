<script setup>
import { defineProps,defineEmits } from 'vue';
import { Moon, Sunny, FolderOpened, DocumentCopy } from '@element-plus/icons-vue'
const props = defineProps({
    menuList: []
})
const emits = defineEmits(['changeMenus', 'isdelDocs'])
function changeMenus(menus){
    emits('changeMenus',menus)
}

</script>
<template>
    <div class="custom-menu-tree">
        <template v-for="(menus, key) in props.menuList">

            <el-sub-menu v-if="menus.is_dir == 1" :index="key">
                <template #title>
                    <el-icon>
                        <FolderOpened />
                    </el-icon>{{ menus.name }}
                </template>
                <MenuTree v-if="menus.children != null" :menuList="menus.children" />
            </el-sub-menu>
            <el-menu-item v-else :index="key" @click="changeMenus(menus)">
                <el-icon>
                    <DocumentCopy />
                </el-icon>{{ menus.name }}
            </el-menu-item>

        </template>
    </div>
</template>