import { createRouter, createWebHistory } from 'vue-router';

import Home from '../components/Home';
import Room from '../components/Room';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/c/:id',
        name: 'Room',
        component: Room
    }
    //{
        //path: '/about',
        //name: 'About',
        //// route level code-splitting
        //// this generates a separate chunk (about.[hash].js) for this route
        //// which is lazy-loaded when the route is visited.
        //component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    //}
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
