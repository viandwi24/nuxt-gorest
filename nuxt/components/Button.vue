<script lang="ts" setup>
const props = defineProps({
  text: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    default: 'primary',
  },
  size: {
    type: String,
    default: 'md',
  },
  to: {
    type: [String, Object],
    default: undefined,
  },
  href: {
    type: String,
    default: undefined,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})

// state:styles
const defaultStyle = `
  z-1
  cursor-pointer
  border-transparent
  rounded-xl relative
  transition-color duration-300
  outline-none font-semibold
  flex items-center justify-center font-semibold
  focus:outline-none focus:ring-1 focus:ring-offset-1 focus:dark:ring-offset-gray-50 focus:dark:ring-gray-400 focus:ring-gray-600/[0.6] focus:ring-offset-gray-800/[0.6]
  disabled:cursor-not-allowed
`
// focus:outline-none focus:ring-1 focus:ring-offset-1 focus:dark:ring-offset-gray-50 focus:dark:ring-gray-400 focus:ring-gray-600/[0.6] focus:ring-offset-gray-800/[0.6]
const styles = reactive<{
  [key: string]: string
}>({
  none: '',
  primary: 'btn-primary',
  danger: 'btn-danger',
  secondary:
    'text-slate-800 bg-gray-200 border-gray-200 hover:bg-gray-300 dark:text-white dark:border-slate-800 dark:bg-slate-800 dark:hover:bg-slate-700',
  opposite:
    'text-white bg-gray-800 hover:bg-white hover:text-gray-800 hover:border-gray-900 dark:text-gray-800 dark:bg-gray-100 dark:hover:bg-gray-800 dark:hover:text-gray-100 dark:border-white',
})
const sizes = reactive<{
  [key: string]: string
}>({
  lg: 'btn-lg h-13 px-8 text-lg rounded-xl',
  md: 'btn-md h-10 px-8 text-base rounded-xl',
  sm: 'btn-sm h-9 px-6 text-sm rounded-xl',
  xs: 'btn-xs h-6 px-4 text-xs rounded-xl',
})

// state
const selectedStyle = computed(() => styles[props.type] || styles.primary)
const selectedSize = computed(() => sizes[props.size] || sizes.lg)

// methods
const onClick = (event: MouseEvent) => {
  const router = useRouter()
  if (props.to) {
    router.push(props.to)
  }
  if (!props.href) {
    event.preventDefault()
  }
}
</script>

<template>
  <NuxtLink
    v-if="to"
    tag="a"
    :to="to"
    :class="`button ${defaultStyle} ${selectedStyle} ${selectedSize}`"
    :disabled="disabled"
  >
    <slot>{{ text }}</slot>
  </NuxtLink>
  <a
    v-else
    :class="`button ${defaultStyle} ${selectedStyle} ${selectedSize}`"
    :href="href"
    :disabled="disabled"
    @click="onClick"
  >
    <slot>{{ text }}</slot>
  </a>
</template>

<style lang="scss">
.button {
  &.btn-primary {
    --bg-color: theme('colors.primary.500');
    --border-color: theme('colors.primary.700');
    --color-hover: theme('colors.primary.600');
  }
  &.btn-danger {
    --bg-color: theme('colors.red.500');
    --border-color: theme('colors.red.700');
    --color-hover: theme('colors.red.600');
  }
  &.btn-success {
    --bg-color: theme('colors.blue.500');
    --border-color: theme('colors.blue.700');
    --color-hover: theme('colors.blue.600');
  }

  background-color: var(--bg-color);
  transition: background-color 0.3s ease;
  &::before {
    border-radius: 0.75rem;
    box-shadow: 0 3px 0 var(--border-color);
    content: '';
    position: absolute;
    left: 0px;
    bottom: 0px;
    right: 0px;
    top: 0px;
    z-index: -1;
  }

  &.btn-sm {
    &::before {
      border-radius: 0.75rem;
      box-shadow: 0 3px 0 var(--border-color);
    }
  }

  &.btn-xs {
    &::before {
      border-radius: 0.75rem;
      box-shadow: 0 2px 0 var(--border-color);
    }
  }

  &.btn-md {
    &::before {
      border-radius: 0.75rem;
      box-shadow: 0 4px 0 var(--border-color);
    }
  }

  &:hover {
    &::before {
      background-color: var(--color-hover);
    }
  }

  &:active {
    &.btn-xs {
      transform: translateY(2px);
    }
    &.btn-sm {
      transform: translateY(3px);
    }
    &.btn-md {
      transform: translateY(4px);
    }

    &::before {
      bottom: 0px;
      box-shadow: none;
    }
  }
}
</style>
