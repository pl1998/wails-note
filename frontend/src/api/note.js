import { apiGet, apiPost, apiDel, apiPut } from './request'

export function getNote(id) {
  return apiGet('/api/note/' + id)
}
export function addNote(data) {
  return apiPost('/api/note', data)
}
