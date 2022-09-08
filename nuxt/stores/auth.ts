import { AxiosError } from 'axios'
import { defineStore } from 'pinia'
import { User } from '~/types'

export interface IAuthState {
  isLogged: boolean
  accessToken: string
  refreshToken: string
  user?: User
}

export const useAuthStore = defineStore('auth', {
  state: (): IAuthState => ({
    isLogged: false,
    accessToken: '',
    refreshToken: '',
    user: undefined,
  }),
  getters: {
    fullName(): string {
      return (this.user?.firstName || '') + ' ' + (this.user?.lastName || '')
    },
  },
  actions: {
    async firstSetupToken() {
      const accessToken = useCookie('accessToken')
      const refreshToken = useCookie('refreshToken')
      this.accessToken = accessToken.value
      this.refreshToken = refreshToken.value
      if (this.accessToken !== '' && this.refreshToken !== '') {
        try {
          this.updateToken()
          await this.fetchMe()
          this.isLogged = true
        } catch (error) {
          this.removeToken()
          this.isLogged = false
        }
      }
    },
    login(email: string, password: string) {
      const http = useHttp()
      return new Promise((resolve, reject) => {
        http({
          method: 'POST',
          url: 'login',
          data: {
            email,
            password,
          },
        })
          .then((response) => {
            if (response.status === 200) {
              this.isLogged = true
              this.accessToken = response.data.accessJWT
              this.refreshToken = response.data.refreshJWT
              this.updateToken()
              this.fetchMe()
                .then(() => resolve(true))
                .catch(() => reject(new Error('Error fetching user')))
            }
          })
          .catch((error) => {
            reject(error)
          })
      })
    },
    refreshAuthToken() {
      const http = useHttp()
      return new Promise((resolve, reject) => {
        http({
          method: 'POST',
          url: 'refresh',
          data: {
            refreshJWT: this.refreshToken,
          },
        })
          .then((response) => {
            this.isLogged = true
            this.accessToken = response.data.accessJWT
            this.refreshToken = response.data.refreshJWT
            this.updateToken()
            this.fetchMe()
              .then(() => resolve(true))
              .catch(() => reject(new Error('Error fetching user')))
          })
          .catch((error) => {
            this.isLogged = false
            reject(error)
          })
      })
    },
    logout() {
      return new Promise((resolve) => {
        this.isLogged = false
        this.removeToken()
        this.user = undefined
        resolve(true)
      })
    },
    fetchMe() {
      const http = useHttp()
      return new Promise((resolve, reject) => {
        http({
          method: 'POST',
          url: 'me',
        })
          .then((response) => {
            if (response.status === 200) {
              this.user = response.data
              resolve(true)
            }
            return reject(new Error(response.statusText))
          })
          .catch((error) => {
            reject(error)
          })
      })
    },
    updateToken() {
      try {
        // inject token to headers
        const http = useHttp()
        http.defaults.headers.common.Authorization = `Bearer ${this.accessToken}`
        // inject to cookies
        const accessToken = useCookie('accessToken')
        const refreshToken = useCookie('refreshToken')
        accessToken.value = this.accessToken
        refreshToken.value = this.refreshToken
      } catch (error) {}
    },
    removeToken() {
      try {
        const accessToken = useCookie('accessToken')
        const refreshToken = useCookie('refreshToken')
        accessToken.value = ''
        refreshToken.value = ''
        this.accessToken = ''
        this.refreshToken = ''
      } catch (error) {}
    },
    async handle401(_err: AxiosError) {
      if (this.refreshToken !== '') {
        try {
          await this.refreshAuthToken()
        } catch (error) {
          console.log(error)
          await this.logout()
          navigateTo('/')
        }
      } else {
        await this.logout()
        navigateTo('/')
      }
    },
  },
})
