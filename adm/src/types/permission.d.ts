interface Item {
  id: number
  pid: number
  name: string
  title: string
  icon?: string
  type: 'menu' | 'action'

  weight?: number
  hidden?: boolean
  default?: boolean
  external?: boolean

  children?: Item[]
}

type Lazy<T> = () => Promise<T>

type RawRouteComponent = RouteComponent | Lazy<RouteComponent>
