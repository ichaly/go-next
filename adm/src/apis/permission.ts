import http from '@/utils/http'

export const getPermission = () => {
  return http.get<Item[]>('/permission/list')
}
