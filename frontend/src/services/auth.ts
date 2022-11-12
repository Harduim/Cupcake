import { Buffer } from 'buffer'
import Cookies from 'universal-cookie'

const { REACT_APP_WEBSITE_DOMAIN } = process.env

const TOKEN_KEY = '@TOKEN_KEY_CUPCAKE'

interface IJWT {
  uid: string
  exp: number
}

const parseJWT = (jwt_string: string) => {
  const buffer = Buffer.from(jwt_string.split('.')[1], 'base64')
  const jwt = JSON.parse(buffer.toString()) as IJWT
  return jwt
}

const isExpired = (expires: number) => {
  return Math.floor(new Date().getTime() / 1000) <= expires
}

const isAuthenticated = () => {
  return getToken() !== null
}

const getToken = (raw = false) => {
  const cookies = new Cookies()
  const jwt_string = cookies.get(TOKEN_KEY) as string | undefined
  if (!jwt_string) return null
  const jwt = parseJWT(jwt_string)
  if (isExpired(jwt.exp)) return null
  if (raw) return jwt_string
  return jwt
}

const login = (jwt_string: string) => {
  const jwt = parseJWT(jwt_string)
  const cookies = new Cookies()
  cookies.set(TOKEN_KEY, jwt_string, {
    path: '/',
    maxAge: jwt.exp,
    domain: REACT_APP_WEBSITE_DOMAIN,
  })
}

const logout = () => {
  const token = getToken()
  if (!token) return
  const cookies = new Cookies()
  cookies.remove(TOKEN_KEY, { path: '/' })
}

export { isAuthenticated, getToken, login, logout }
