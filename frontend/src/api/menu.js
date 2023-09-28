import { store } from './../store/index'
import { apiGet, apiPost, apiDel } from './../api/api'
import { ElMessage } from 'element-plus'
export function getMenuList(data) {
    apiGet('/api/menu', data).then((res) => {
        if (res.code == 200) {
            store.commit('setMenuList', res.data)
        }
    })
}
export function addMenu(data) {
    apiPost('/api/menu', data)
        .then((res) => {
            if (res.code == 200) {
                ElMessage({
                    showClose: true,
                    message: '创建成功~',
                    type: 'success',
                })
                getMenuList({})
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

export function deleleMenu(data) {
    apiDel('/api/menu', data)
        .then((res) => {
            if (res.code == 200) {
                ElMessage({
                    showClose: true,
                    message: '删除成功~',
                    type: 'success',
                })
                getMenuList({})
            }
        })
}