import { createRouter, createWebHistory } from 'vue-router'
import StartView from "@/views/StartView.vue";
import ConfigView from "@/views/ConfigView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/start',
            name: 'start',
            component: StartView,
        },
        {
            path:'/config',
            name: 'config',
            component:ConfigView
        }
    ]
})
router.beforeEach((to, from, next) => {
    localStorage.setItem('lastPath', from.path)
    localStorage.setItem('nowPath', to.path)
    next()
})
export default router