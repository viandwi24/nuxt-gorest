<script lang="ts" setup>
import Card from '../Card/index.vue'
import { Brand } from '~/types'

defineProps({
  brand: {
    type: Object as () => Brand,
    required: true,
  },
})
</script>

<template>
  <Anchor
    class="group hover:no-underline flex h-full"
    :to="`/restaurant/${brand.id}`"
  >
    <Card
      class="h-full flex flex-col relative overflow-hidden bg-gray-200 shadow-lg dark:group-hover:border-primary-200"
    >
      <div class="w-full h-full relative overflow-hidden">
        <div
          class="transition-all flex justify-center items-center transform duration-300 absolute z-10 bg-slate-900/30 group-hover:bg-slate-900/50 w-full h-full"
        >
          <div
            class="text-xl font-extrabold text-gray-100 drop-shadow-xl group-hover:animate-pulse text-center"
          >
            {{ brand.name }}
          </div>
        </div>
        <div
          class="transition-all transform duration-300 group-hover:scale-110"
          :style="{
            backgroundImage: `url(${brand.image})`,
            backgroundSize: 'cover',
            backgroundPosition: 'center',
            backgroundRepeat: 'no-repeat',
            width: '100%',
            height: '100%',
          }"
        />
      </div>
      <div
        class="text-sm p-2 flex justify-between items-center bg-slate-800 text-gray-200 font-semibold"
      >
        <div class="font-bold flex items-center">
          <icon-ic:baseline-star class="text-yellow-500" />
          <span class="ml-1 mt-0.5">{{ brand.rating }}</span>
        </div>
        <div class="font-light">{{ brand.avearageTime }} min</div>
      </div>
    </Card>
  </Anchor>
</template>
