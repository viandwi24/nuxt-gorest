import { AxiosError } from 'axios'
import { useAuthStore } from '~/stores/auth'

export interface Auth {
  isLogged: boolean
  login: (email: string, password: string) => Promise<boolean>
}

export default defineNuxtPlugin(async (nuxtApp) => {
  const store = useAuthStore()
  const http = useHttp()

  // define interceptor
  http.interceptors.request.use(
    (config) => {
      return config
    },
    (err) => {
      return Promise.reject(err)
    }
  )
  http.interceptors.response.use(
    (config) => {
      return config
    },
    async (err) => {
      if (err instanceof AxiosError) {
        // if errpr in login
        if (err.request instanceof XMLHttpRequest) {
          if (err.request.responseURL.includes('login')) {
            return Promise.reject(err)
          }
        }

        // errs
        if (err.code === 'ERR_BAD_REQUEST' && err.response?.status === 401) {
          await store.handle401(err)
        }
      }
      return Promise.reject(err)
    }
  )

  // provide auth
  nuxtApp.provide('auth', {
    store,
    isLogged: store.isLogged,
    login: store.login,
  })

  // setup
  // await store.setupToken()
  if (store.firstSetupToken) await store.firstSetupToken()

  // middleware
  // addRouteMiddleware('auth', )
})
