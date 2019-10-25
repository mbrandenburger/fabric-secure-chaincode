import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from './views/Login'
import Bidder from './views/BidderHome'

Vue.use(VueRouter)

const routes = [
    { path: '/bidder', component: Bidder},
    { path: '/login', component: Login },
    { path: '/', component: Login },
    { path: '*', redirect: '/' }
]

export default new VueRouter({
    routes // short for `routes: routes`
})
