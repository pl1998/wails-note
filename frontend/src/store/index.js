import { defineStore } from 'pinia'

export const useBaseInfo = defineStore('baseInfo', {
  state: () => ({
    theme: localStorage.getItem('theme') || 'light', //dark
  }),
  actions: {
    changeTheme(type){
      localStorage.setItem('theme', type)
      this.theme = type
    }
  },
  getters: {},
})
