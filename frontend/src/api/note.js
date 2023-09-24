import { store } from './../store/index'
import { apiGet } from './../api/api'

export function getNote(id) {
    apiGet('/api/note/'+id).then((res) => {
        if (res.code == 200) {
            store.commit('setNote', res.data)
        }
    })
}