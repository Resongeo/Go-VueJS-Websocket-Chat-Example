import { createRouter, createWebHistory } from 'vue-router'
import ConnectView from '../views/ConnectView.vue'
import ChatView from '../views/ChatView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'main',
      component: ConnectView
    },
    {
      path: '/chat',
      name: 'chat',
      component: ChatView
    }
  ]
})

export default router