import { QueryClient } from '@tanstack/react-query'
import axios from 'axios'
import { getToken } from './auth'

const { REACT_APP_API_URL } = process.env

export { REACT_APP_API_URL }

export const queryClient = new QueryClient()

const api = axios.create({
  baseURL: REACT_APP_API_URL,
})

api.interceptors.request.use(async config => {
  const token = getToken(true)
  if (token && config && config.headers) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default api
