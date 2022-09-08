<script lang="ts" setup>
import {
  TransitionRoot,
  TransitionChild,
  Dialog,
  DialogPanel,
  DialogTitle,
} from '@headlessui/vue'
import { AxiosError } from 'axios'
import { useCart } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'

const cart = useCart()
const auth = useAuthStore()

// is login
const loginModalIsOpen = ref(false)
const email = ref('test@mail.com')
const password = ref('password')
const closeModal = () => {
  loginModalIsOpen.value = false
}
const openModal = () => {
  loginModalIsOpen.value = true
}
const login = async () => {
  try {
    await auth.login(email.value, password.value)
  } catch (error) {
    if (error instanceof AxiosError) {
      return alert('Check your email and password again')
    }
  }
  closeModal()
}
const logout = async () => {
  await auth.logout()
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
              <Button text="Login" size="sm" @click="openModal" />
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
    <TransitionRoot appear :show="loginModalIsOpen" as="template">
      <Dialog as="div" class="relative z-10" @close="closeModal">
        <TransitionChild
          as="template"
          enter="duration-300 ease-out"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div class="fixed inset-0 bg-black bg-opacity-25" />
        </TransitionChild>
        <div class="fixed inset-0 overflow-y-auto">
          <div
            class="flex min-h-full items-center justify-center p-4 text-center"
          >
            <TransitionChild
              as="template"
              enter="duration-300 ease-out"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="duration-200 ease-in"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <DialogPanel
                class="w-full max-w-md transform overflow-hidden rounded-2xl bg-slate-800 p-6 text-left align-middle shadow-xl transition-all"
              >
                <div class="flex justify-between">
                  <DialogTitle
                    as="h3"
                    class="text-2xl font-bold leading-6 text-gray-100"
                  >
                    LOGIN
                  </DialogTitle>
                  <button
                    class="text-red-500 font-bold text-4xl"
                    @click="closeModal"
                  >
                    x
                  </button>
                </div>
                <div class="mt-2">
                  <Card class="mb-4">
                    <CardContent>
                      <div class="flex flex-col space-y-2 mb-4">
                        <label>Email</label>
                        <FormTextInput v-model="email" class="w-full" />
                      </div>
                      <div class="flex flex-col space-y-2 mb-4">
                        <label>Password</label>
                        <FormTextInput
                          v-model="password"
                          type="password"
                          class="w-full"
                        />
                      </div>
                    </CardContent>
                  </Card>
                  <Button text="Login" @click="login" />
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>
