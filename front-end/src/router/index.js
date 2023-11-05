import { createRouter, createWebHistory } from 'vue-router'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: () => import('@/views/HomeView.vue') },
    { path: '/collect', component: () => import('@/views/CollectView.vue') },
    { path: '/login', component: () => import('@/views/LoginView.vue') },
    { path: '/user', component: () => import('@/views/UserView.vue') }
  ]
})

export default router
