import axios from 'axios'

export function createAxiosInstanace() {
  const instance = axios.create({
    baseURL: 'http://localhost:8999/api/v1',
    headers: {
      'Content-Type': 'application/json',
      Accept: '*/*',
      // 'Accept-Encoding': 'gzip, deflate, br',
      // Connection: 'keep-alive',
    },
    // withCredentials: true,
  })

  return instance
}

export default defineNuxtPlugin((nuxtApp) => {
  const http = createAxiosInstanace()
  nuxtApp.provide('http', http)
})
