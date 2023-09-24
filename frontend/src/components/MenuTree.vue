<script setup>
import { defineProps } from 'vue';
import { Moon, Sunny, FolderOpened, DocumentCopy } from '@element-plus/icons-vue'
const menuList = defineProps([
])
</script>
<template>
    <div class="custom-menu-tree">
        <template v-for="(menus,key) in menuList">
            <el-menu-item>
                <el-sub-menu :index="`menu_key`+key">
                    <template #title>
                        <el-icon>
                            <FolderOpened />
                        </el-icon>{{ menus.name }}
                    </template>
                    <el-menu-item v-for="(note,key) in menus.notes" :index="`menu_key`+key">
                        <el-icon>
                            <DocumentCopy />
                        </el-icon>{{ note.name }}
                    </el-menu-item>
                    <el-menu-item-group v-for="(menu,key) in menus.children">
                        <el-menu-item index="{{menu.menu_id}}">
                            <el-icon>
                                <FolderOpened />
                            </el-icon>{{ menu.name }}</el-menu-item>
                        <el-menu-item :index="`note_key`+key" v-for="note in menu.notes"><el-icon>
                                <DocumentCopy />
                            </el-icon> {{ note.name }}</el-menu-item>
                    </el-menu-item-group>
                </el-sub-menu>
            </el-menu-item>
        </template>
    </div>
</template>