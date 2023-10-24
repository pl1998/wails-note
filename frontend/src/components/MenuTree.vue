<script setup>
import { FolderOpened, DocumentCopy } from '@element-plus/icons-vue'
const props = defineProps({
  menuList: {
    type: Array,
    default: () => [],
  },
})

const openMenu = (event, menus) => {
  event.preventDefault()
}
const emit = defineEmits(['menuClick', 'isEditDocs'])
const menuClick = (menu) => {
    emit('menuClick', menu)
}
</script>
<template>
  <div class="custom-menu-tree">
    <template v-for="item in menuList" :key="item.menu_id">
      <el-sub-menu
        :index="item.menu_id + ''"
        v-if="item.notes.length > 0"
        @click.right.native="openMenu($event, menus.menu_id)"
      >
        <template #title>
          <el-icon> <FolderOpened /> </el-icon>
          <span>{{ item.name }}</span>
        </template>
        <el-menu-item
          v-for="note in item.notes"
          :key="'note' + note.note_id"
          :index="item.menu_id + '-' + note.note_id"
          @click.left.native="menuClick(note)"
        >
          <el-icon><DocumentCopy /></el-icon>
          <span>{{ note.name }}</span>
        </el-menu-item>
      </el-sub-menu>
      <el-menu-item
        :index="item.menu_id + ''"
        v-else
        @click.left.native="menuClick(note)"
      >
        <el-icon><DocumentCopy /></el-icon>
        <span>{{ item.name }}</span>
      </el-menu-item>
    </template>
  </div>
</template>
