import { store } from './../store/index'
import { apiGet, apiPost } from './../api/api'
export function getMenuList(data) {
    apiGet('/api/menu', data).then((res) => {
        if (res.code == 200) {
            store.commit('setMenuList', res.data)
        }
    })
}
export function addNote(data) {
    apiPost('/api/note', data)
        .then((res) => {
            if (res.code == 200) {
                getMenuList({})
            }
        })
}