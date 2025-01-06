import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/*',
      component: () => import('../views/Layout.vue'),
      children: [
        {
          path: '/',
          name: 'Tools',
          component: () => import('../views/Tools.vue')
        }
      ]
    },
  ],
})

router.beforeEach((to, from, next) => {
  if (!localStorage.getItem('tks') && to.name !== 'Login') {
    next({
      name: 'Login'
    })
  } else if (localStorage.getItem('tks') && to.name === 'Login') {
    next({
      name: 'Tools'
    })
  } else {
    next()
  }
});

export default router
