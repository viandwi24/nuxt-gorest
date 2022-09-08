import { defineStore } from 'pinia'
import { Menu } from '~/types'

export interface CartItem {
  id: number
  name: string
  quantity: number
  menu: Menu
}

export interface ICartState {
  restaurantId?: number
  carts: CartItem[]
}

export const useCart = defineStore('cart', {
  state: (): ICartState => ({
    restaurantId: undefined,
    carts: [],
  }),
  actions: {
    update() {
      // console.log('updating cart lists...')
    },
    add(menu: Menu) {
      const cart = this.carts.find((c) => c.id === menu.id)
      if (cart) {
        cart.quantity++
      } else {
        if (this.carts.length === 0) this.restaurantId = menu.restaurantId
        if (menu.restaurantId !== this.restaurantId) {
          alert('You can only order from one restaurant at a time')
          return
        }
        this.carts.push({
          id: menu.id,
          name: menu.name,
          quantity: 1,
          menu,
        })
      }
    },
    remove(menu: Menu) {
      const cart = this.carts.find((c) => c.id === menu.id)
      if (cart) {
        this.carts = this.carts.filter((c) => c.id !== menu.id)
      }
    },
    reset() {
      this.carts = []
      this.restaurantId = undefined
    },
    decrease(menu: Menu) {
      const cart = this.carts.find((c) => c.id === menu.id)
      if (cart) {
        if (cart.quantity === 1) this.remove(menu)
        cart.quantity -= 1
      }
    },
    increase(menu: Menu) {
      const cart = this.carts.find((c) => c.id === menu.id)
      if (cart) cart.quantity += 1
    },
  },
  getters: {
    totalPrice(): number {
      return this.carts.reduce(
        (acc, cart) => acc + cart.menu.price * cart.quantity,
        0
      )
    },
    eachItems(): CartItem[] {
      return this.carts
    },
  },
})
