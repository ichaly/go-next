interface Permission {
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
