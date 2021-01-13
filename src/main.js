import Vue from 'vue'
import App from './App.vue'
import VueRouter from "vue-router";
import NewsList from "@/components/NewsList";
import NewsNew from "@/components/NewsNew";
import AuthorLogin from "@/components/AuthorLogin";

Vue.config.productionTip = false;

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        component: NewsList,
    },
    {
        path: '/write',
        component: NewsNew,
    },
    {
        path: '/login',
        component: AuthorLogin,
    }
]

const router = new VueRouter({
    mode: 'hash',
    routes,
});

new Vue({
    render: h => h(App),
    router,
}).$mount('#app');
