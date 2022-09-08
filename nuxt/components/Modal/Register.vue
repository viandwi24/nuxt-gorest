<script lang="ts" setup>
import {
  TransitionChild,
  Dialog,
  DialogPanel,
  DialogTitle,
} from '@headlessui/vue'
import { AxiosError } from 'axios'
import { useAuthStore } from '~/stores/auth'
const auth = useAuthStore()
const email = ref('test@mail.com')
const password = ref('password')
const firstName = ref('vian')
const lastName = ref('dwi')
const emit = defineEmits(['close'])
const register = async () => {
  try {
    await auth.register(email.value, password.value, {
      firstName: firstName.value,
      lastName: lastName.value,
    })
    emit('close')
  } catch (error) {
    if (error instanceof AxiosError) {
      return alert('Check your email and password again')
    }
  }
  close()
}
</script>
<template>
  <Dialog as="div" class="relative z-10" @close="emit('close')">
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
      <div class="flex min-h-full items-center justify-center p-4 text-center">
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
                REGISTER
              </DialogTitle>
              <button
                class="text-red-500 font-bold text-4xl"
                @click="emit('close')"
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
                    <div>
                      <label>First Name</label>
                      <FormTextInput v-model="firstName" class="w-full" />
                    </div>
                    <div>
                      <label>Last Name</label>
                      <FormTextInput v-model="lastName" class="w-full" />
                    </div>
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
              <Button text="Login" @click="register" />
            </div>
          </DialogPanel>
        </TransitionChild>
      </div>
    </div>
  </Dialog>
</template>
