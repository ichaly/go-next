import http from '@/utils/http'

export interface Permission {
  id: number
  pid: number
  name: string
  title: string
  icon?: string
  weight?: number
  hidden?: boolean
  external?: boolean
  type: 'menu' | 'action'
}

export const getPermission = () => {
  return http.get<Permission[]>('/permission/list')
}
