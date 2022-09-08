<script lang="ts" setup>
import { Brand } from '~/types'
import { getBrands } from '~/services/api'

useHead(() => ({
  title: 'Home',
}))

definePageMeta({
  layout: 'food',
})

const brands = reactive<Brand[]>([])
const isLoading = ref(true)

const fetchBrands = async () => {
  const _brands = await getBrands()
  brands.push(..._brands)
  isLoading.value = false
}

onMounted(() => fetchBrands())
</script>

<template>
  <div class="px-4 my-4 mx-auto max-w-screen-2xl">
    <section class="banners mb-8">
      <div
        class="h-96 p-10 relative flex items-center rounded-xl bg-primary-400 shadow-lg"
      >
        <div class="absolute bottom-0 right-0 h-full pr-20 pt-20">
          <img
            src="/assets/images/illustrations/business-3d-close-up-of-businessman-in-dark-blue-suit-with-coffee-mug.png"
            alt="illustrations"
            class="h-full"
          />
        </div>
        <div class="w-1/3">
          <div class="text-5xl font-bold mb-4">
            Selamat datang dan kami perkenalkan kami dari foody
          </div>
          <div class="mb-4 font-light">
            Kami punya banyak menu dari banyak restaurant disekitar mu, mari
            coba in makanan dan minumannya ! dijamin bakalan enak dan bikin kamu
            ketagihan.
          </div>
          <div class="flex space-x-2">
            <!-- <Button text="Lihat Semua Menu" size="sm" /> -->
            <Button text="Masuk" size="sm" />
          </div>
        </div>
      </div>
    </section>
    <section class="mb-8">
      <div class="text-2xl font-bold">Restaurants Near You</div>
      <div class="mt-4 grid grid-cols-6 gap-4">
        <template v-if="!isLoading">
          <BrandCard
            v-for="(brand, i) in brands"
            :key="i"
            :brand="brand"
            class="h-42"
          />
        </template>
        <template v-else>
          <BrandCardPlaceholder v-for="i in 6" :key="i" class="h-42" />
        </template>
      </div>
    </section>
  </div>
</template>
