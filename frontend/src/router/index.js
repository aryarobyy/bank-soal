import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'
// 1. Impor komponen ForgotPassword di sini
import ForgotPassword from '../views/Auth/ForgotPassword.vue'

const routes = [
  {
    path: '/',
    redirect: '/register'
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  // 2. Tambahkan rute baru untuk halaman forgot password
  {
    path: '/forgot-password',
    name: 'ForgotPassword',
    component: ForgotPassword
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router