import { createStore } from 'vuex'

export const store = createStore({
  state() {
    return {
      theme: "light",
      menuList: [],
      notes: {
        note_id: 0,
        name: 'name',
        content: '',
        menu_id: 1
      }
    }
  },
  mutations: {
    setTheme(state, theme) {
      state.theme = theme
    },
    setMenuList(state, list) {
      state.menuList = list
    },
    setNote(state, notes) {
      state.notes = notes
    }
  }
})