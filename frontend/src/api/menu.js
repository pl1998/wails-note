import { apiGet, apiPost, apiDel, apiPut } from './request'
export function getMenuList(data) {
  return apiGet('/api/menu', data)
}
export function addMenu(data) {
  return apiPost('/api/menu', data)
}

export function updateMenu(data) {
  return apiPut('/api/menu', data.menu_id, data)
}

export function deleleMenu(data) {
  return apiDel('/api/menu', data)
}
