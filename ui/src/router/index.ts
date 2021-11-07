import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Home from '../views/Home.vue';
import Difficulty from '../views/Difficulty.vue';
import Topic from '../views/Topic.vue';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/difficulty',
        name: 'Difficulty',
        component: Difficulty,
    },
    {
        path: '/topic',
        name: 'Topic',
        component: Topic,
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
