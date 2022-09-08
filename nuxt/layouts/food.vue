<script lang="ts" setup>
import { TransitionRoot } from '@headlessui/vue'
import { useCart } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'

const cart = useCart()
const auth = useAuthStore()

// is login
const loginModalIsOpen = ref(false)
const closeLoginModal = () => {
  loginModalIsOpen.value = false
}
const openLoginModal = () => {
  loginModalIsOpen.value = true
}
const logout = async () => {
  await auth.logout()
}

// register
const registerModalIsOpen = ref(false)
const closeRegisterModal = () => {
  registerModalIsOpen.value = false
}
const openRegisterModal = () => {
  registerModalIsOpen.value = true
}

// cart
const cartModalIsOpen = ref(false)
watch(
  () => cart.totalPrice,
  () => {
    if (cart.totalPrice > 0) {
      cartModalIsOpen.value = true
    }
  }
)
const toggleCart = () => {
  cartModalIsOpen.value = !cartModalIsOpen.value
}
</script>

<template>
  <div>
    <div class="navbar">
      <div
        class="relative overflow-hidden border-b border-gray-900/10 dark:border-gray-50/[0.2]"
      >
        <div
          class="mx-auto max-w-screen-2xl px-4 flex justify-between items-center h-16"
        >
          <Anchor
            class="font-black text-xl hover:no-underline"
            text="FOODY"
            to="/"
          />
          <div class="flex space-x-4">
            <template v-if="!auth.isLogged">
              <Button text="Register" size="sm" @click="openRegisterModal" />
              <Button text="Login" size="sm" @click="openLoginModal" />
            </template>
            <template v-else>
              <div class="flex items-center">
                <icon-ph:user-bold class="mr-2" />
                <span>{{ auth.fullName }}</span>
              </div>
              <Button text="Logout" size="sm" type="danger" @click="logout" />
            </template>
          </div>
        </div>
      </div>
    </div>
    <slot />
    <!-- carts -->
    <div class="fixed bottom-0 right-0">
      <div class="relative shadow">
        <button
          class="w-full block px-6 py-2 rounded-t w-96 bg-slate-700 flex justify-between items-center"
          @click.native="toggleCart"
        >
          <div class="font-bold">Cart</div>
          <div class="flex items-center">
            <icon-ic:outline-minus />
          </div>
        </button>
        <div
          v-if="cartModalIsOpen"
          class="flex flex-col space-y-2 px-4 py-4 w-96 max-w-96 bg-slate-800"
        >
          <div v-for="(item, i) in cart.eachItems" :key="i">
            <div class="flex justify-between items-center space-x-4">
              <div class="flex-1">{{ item.menu.name }}</div>
              <div class="flex items-center justify-end">
                ${{ (item.menu.price * item.quantity).toFixed(2) }}
              </div>
              <div class="flex space-x-2">
                <button
                  class="rounded-full w-6 h-6 flex justify-center items-center bg-slate-700"
                  @click.native="cart.decrease(item.menu)"
                >
                  -
                </button>
                <div>{{ item.quantity }}</div>
                <button
                  class="rounded-full w-6 h-6 flex justify-center items-center bg-slate-700"
                  @click.native="cart.increase(item.menu)"
                >
                  +
                </button>
              </div>
            </div>
          </div>
          <div class="flex justify-between items-center">
            <div>Total :</div>
            <div>${{ cart.totalPrice.toFixed(2) }}</div>
          </div>
          <div class="flex justify-end mt-2 space-x-2">
            <Button
              size="sm"
              type="danger"
              text="Reset"
              @click="cart.reset()"
            />
            <Button
              :disabled="cart.carts.length === 0"
              size="sm"
              text="Checkout"
            />
          </div>
        </div>
      </div>
    </div>
    <!-- modals -->
    <TransitionRoot
      v-if="loginModalIsOpen"
      :show="loginModalIsOpen"
      appear
      as="template"
    >
      <ModalLogin @close="closeLoginModal" />
    </TransitionRoot>
    <TransitionRoot
      v-if="registerModalIsOpen"
      :show="registerModalIsOpen"
      appear
      as="template"
    >
      <ModalRegister @close="closeRegisterModal" />
    </TransitionRoot>
  </div>
</template>
