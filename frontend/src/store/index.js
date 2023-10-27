import { defineStore } from 'pinia'

export const useBaseInfo = defineStore('baseInfo', {
  state: () => ({
    theme: localStorage.getItem('theme') || 'light', //dark
    showMenu: localStorage.getItem('showMenu') || true,
    showAside: localStorage.getItem('showAside') || true,
  }),
  actions: {
    changeTheme(type){
      localStorage.setItem('theme', type)
      this.theme = type
    },
    changeShowMenu(){
      this.showMenu = !this.showMenu
      localStorage.setItem('showMenu', this.showMenu)
    },
    changeShowAside(){
      this.showAside = !this.showAside
      localStorage.setItem('showAside', this.showAside)
    }
  },
  getters: {},
})
