<script setup>
import {reactive } from 'vue'
import {Greet} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  text: "Please enter your name below 👇",
  content : "",
})


function greet() {
  Greet(data.name).then(result => {
    data.text = result
  })
}

</script>
<template>
    <div>
        <h1>这是一个富文本</h1>
        <v-md-editor v-model="text" height="400px"></v-md-editor>
        <h1>这是一个富文本</h1>
    </div>
  
</template>
  