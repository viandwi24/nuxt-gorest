import { AxiosInstance } from 'axios'
import { Auth } from '~/plugins/auth'

declare module '#app' {
  interface NuxtApp {
    $http: AxiosInstance
    $auth: Auth
  }
}

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $http: AxiosInstance
    $auth: Auth
  }
}

export {}
