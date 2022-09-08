export interface Menu {
  id: number
  restaurantId: number
  name: string
  image: string
  price: number
  description: string
  category: string
}

export interface Brand {
  id: number
  name: string
  image: string
  rating: number
  avearageTime: number
  menus: Menu[]
}

export interface User {
  userID: number
  firstName: string
  lastName: string
}
