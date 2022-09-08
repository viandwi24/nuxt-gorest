<script lang="ts" setup>
import { getBrand } from '~/services/api'
import { useCart } from '~/stores/cart'
import { Menu } from '~/types'

definePageMeta({
  layout: 'food',
})

const { id } = useRoute().params

const brand = await getBrand(Number(id))

if (!brand || typeof brand === 'undefined') {
  createError({
    statusCode: 404,
    fatal: true,
  })
}

const isLoading = ref(true)
useHead(() => ({
  title: `${brand?.name} - Restaurant`,
}))

// carts
const cart = useCart()
const addMenu = (menu: Menu) => {
  cart.add(menu)
}

onMounted(() => {
  setTimeout(() => (isLoading.value = false), 1000)
})
</script>

<template>
  <div class="px-4 my-4 mx-auto max-w-screen-2xl">
    <section class="banners mb-8">
      <div
        class="h-96 p-10 relative flex overflow-hidden items-end rounded-xl bg-blue-400 shadow-lg"
      >
        <div class="w-1/3 z-10">
          <div class="text-5xl font-bold">
            {{ brand?.name }}
          </div>
          <!-- <div class="flex space-x-2">
            <Button text="Masuk" size="sm" />
          </div> -->
        </div>
        <div class="absolute z-0 top-0 left-0 w-full h-full flex">
          <div class="relative flex-1">
            <div
              class="transition-all flex justify-center items-center transform duration-300 absolute z-10 bg-slate-900/50 group-hover:bg-slate-900/50 w-full h-full"
            />
            <div
              class="transition-all transform duration-300 group-hover:scale-110 w-full h-full"
              :style="{
                backgroundImage: `url(${brand?.image})`,
                backgroundSize: 'cover',
                backgroundPosition: 'center',
                backgroundRepeat: 'no-repeat',
              }"
            />
          </div>
        </div>
      </div>
    </section>
    <section class="mb-8">
      <div class="text-2xl font-bold">Food menu</div>
      <div class="mt-4 grid grid-cols-6 gap-4">
        <template v-if="!isLoading">
          <FoodCard
            v-for="(menu, i) in brand?.menus || []"
            :key="i"
            :menu="menu"
            class="h-42"
            @click.native="addMenu(menu)"
          />
        </template>
        <template v-else>
          <BrandCardPlaceholder v-for="i in 6" :key="i" class="h-42" />
        </template>
      </div>
      <div v-if="(brand?.menus || []).length === 0">
        <div class="text-center text-gray-500">
          Tidak ada menu yang tersedia
        </div>
      </div>
    </section>
  </div>
</template>
