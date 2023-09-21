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
    <div class="left-main">
      
    </div>
</template>
<style>
.left-main{
    
}
</style>
  