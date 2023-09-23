import { createStore } from 'vuex'

export const store = createStore({
    state () {
      return {
        theme:"light",
        menuList:[]
      }
    },
    mutations: {
      setTheme(state,theme)
      {
        state.theme = theme
      },
      setMenuList(state,list)
      {
        state.menuList = list
      }
    }
  })