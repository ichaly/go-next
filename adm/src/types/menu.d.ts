interface Menu {
  name: string
  pid?: number
  icon?: string
  title?: string
  children?: Menu[]
}

type Lazy<T> = () => Promise<T>

type RawRouteComponent = RouteComponent | Lazy<RouteComponent>