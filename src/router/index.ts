import Home from '@/components/Home.vue'
import Login from '@/components/Login.vue'
import Register from '@/components/Register.vue'
import Timeline from '@/components/Timeline.vue'
import Messages from '@/components/Messages.vue'


import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      // Redirection vers la route '/Login' par défaut
      path: '/',
      redirect: '/Login'
    },
    {
      path: '/Login',
      name: 'Login',
      component: Login
    },
    {
      path: '/Register',
      name: 'Register',
      component: Register
    },
    {
      path: '/Home',
      name: 'Home',
      component:Home
    },
    {
      path: '/Timeline',
      name: 'Timeline',
      component:Timeline
    },
    {
      path: '/Messages',
      name: 'Messages',
      component: Messages
    },

  ]
})

export default router
