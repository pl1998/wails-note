<template>
    <div class="theme-switch">
        <el-select v-model="theme" placeholder="Select theme" class="theme-select" @change="selectTheme">
            <el-option v-for="themeItem in themes" :key="themeItem" :label="themeItem" :value="themeItem">
            </el-option>
        </el-select>
    </div>
</template>
  
<script setup>
import { watch, computed } from "vue"
import { store } from './../store/index'
const theme = computed(() => store.state.theme)

const themes = ["default", "dark", "light"]
const themePath = computed(() => {
    return `./theme-${store.state.theme}.css`
})
watch(themePath,() => {
    const themeLink = document.querySelector('link[rel=stylesheet][href^="./theme"]');
    if (themeLink) {
        themeLink.href = themePath.value;
    } else {
        const newThemeLink = document.createElement('link');
        console.log(newThemeLink)
        newThemeLink.rel = 'stylesheet';
        newThemeLink.href = themePath.value;
        //document.head.appendChild(newThemeLink);
    }
})
function selectTheme(value){
    store.commit('setTheme',value)
}

</script>
  
<style scoped>
.theme-switch {
    text-align: right;
    margin-bottom: 20px;
}

.theme-select {
    width: 120px;
}
</style>