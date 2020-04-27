import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const router = new Router({

    base: process.env.BASE_URL,
    routes: [
        {
            path: '/login',
            name: 'Login',
            component: () => import(/* webpackChunkName: "Login" */ './views/Login.vue')
        },
        {
            path: '/*',
            component: () => import(/* webpackChunkName: "Layout" */ './views/Layout.vue'),
            children: [
                {
                    path: '/',
                    name: 'Tools',
                    component: () => import(/* webpackChunkName: "Tools" */ './views/Tools.vue')
                }
            ]
        },
    ]
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
