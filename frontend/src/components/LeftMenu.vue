
<script setup>
import { ref } from 'vue'
import Node from 'element-plus/es/components/tree/src/model/node'

let id = 1000
function append(tree) {
  const newChild = { id: id++, label: 'testtest', children: [] }
  if (!data.children) {
    data.children = []
  }
  data.children.push(newChild)
  dataSource.value = [...dataSource.value]
}
function remove(node) {
  const parent = node.parent
  const children =  parent.data.children || parent.data
  const index = children.findIndex((d) => d.id === data.id)
  children.splice(index, 1)
  dataSource.value = [...dataSource.value]
}
</script>
<template>
  <div class="custom-tree-container">
    <p>Using render-content</p>
    <el-tree :data="dataSource" show-checkbox node-key="id" default-expand-all :expand-on-click-node="false"
      :render-content="renderContent" />
    <p>Using scoped slot</p>
    <el-tree :data="dataSource" show-checkbox node-key="id" default-expand-all :expand-on-click-node="false">
      <template #default="{ node, data }">
        <span class="custom-tree-node">
          <span>{{ node.label }}</span>
          <span>
            <a @click="append(data)"> Append </a>
            <a style="margin-left: 8px" @click="remove(node, data)"> Delete </a>
          </span>
        </span>
      </template>
    </el-tree>
  </div>
</template>