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
        <v-md-editor v-model="text"  height="400px"></v-md-editor>
        <!-- left-toolbar="undo redo clear item  | emoji | tip" -->
    </div>

</template>
  