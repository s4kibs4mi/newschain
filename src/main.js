import Vue from 'vue'
import App from './App.vue'
import VueRouter from "vue-router";
import VueSimpleMarkdown from "vue-simple-markdown";
import NewsList from "@/components/NewsList";
import NewsCreate from "@/components/NewsCreate";
import AuthorProfile from "@/components/AuthorProfile";
import NewsView from "@/components/NewsView";

Vue.config.productionTip = false;
Vue.use(VueRouter);
Vue.use(VueSimpleMarkdown);

const routes = [
    {
        path: '/',
        component: NewsList,
    },
    {
        path: '/write',
        component: NewsCreate,
    },
    {
        path: '/view',
        component: NewsView,
    },
    {
        path: '/profile',
        component: AuthorProfile,
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
