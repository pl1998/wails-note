import { store } from './../store/index'
import { apiGet, apiPost, apiDel, apiPut } from './../api/api'
import { ElMessage } from 'element-plus'
export function getMenuList(data) {
    apiGet('/api/menu', data).then((res) => {
        if (res.code == 200) {
            store.commit('setMenuList', res.data)
        }
    })
}
export function addMenu(data) {
    return apiPost('/api/menu', data)
}


export function updateMenu(data) {
    apiPut('/api/menu',data.menu_id, data)
        .then((res) => {
            if (res.code == 200) {
                getMenuList({})
                ElMessage({
                    message: '保存成功~',
                    type: 'success'
                })
            } else {
                ElMessage({
                    message: '保存失败~',
                    type: 'error'
                })
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