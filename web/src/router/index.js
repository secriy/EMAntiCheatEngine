import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);
// 解决重复点击导航路由报错
const originalPush = Router.prototype.push;
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err);
};

export default new Router({
    routes: [
        {
            path: '/',
            redirect: '/login'
        },
        {
            path: '/',
            component: () => import(/* webpackChunkName: "home" */ '../components/Home.vue'),
            meta: { title: '自述文件' },
            children: [
                {
                    path: '/dashboard',
                    component: () => import(/* webpackChunkName: "dashboard" */ '../views/Dashboard.vue'),
                    meta: { title: '系统首页' }
                },
                {
                    path: '/userTable',
                    component: () => import(/* webpackChunkName: "userTable" */ '../views/UserTable.vue'),
                    meta: { title: '用户管理', permission: true }
                },
                {
                    path: '/hookTable',
                    component: () => import(/* webpackChunkName: "newsTable" */ '../views/HookTable.vue'),
                    meta: { title: '应用日志' }
                },
                {
                    path: '/antiTable',
                    component: () => import(/* webpackChunkName: "newsTable" */ '../views/AntiTable.vue'),
                    meta: { title: '内核日志' }
                },
                {
                    path: '/404',
                    component: () => import(/* webpackChunkName: "404" */ '../views/404.vue'),
                    meta: { title: '404' }
                },
                {
                    path: '/403',
                    component: () => import(/* webpackChunkName: "403" */ '../views/403.vue'),
                    meta: { title: '403' }
                }
            ]
        },
        {
            path: '/login',
            component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue'),
            meta: { title: '登录' }
        },
        {
            path: '*',
            redirect: '/404'
        }
    ]
});
