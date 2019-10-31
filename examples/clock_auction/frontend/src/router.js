import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from './views/Login'
import AuctionInfo from './views/AuctionInfo'
import Results from './views/MyResults'
import PlaceBid from './views/PlaceBid'

Vue.use(VueRouter)

const routes = [
    { path: '/auction_info', component: AuctionInfo},
    { path: '/my_results', component: Results},
    { path: '/place_bid', component: PlaceBid},
    { path: '/login', component: Login },
    { path: '/', component: Login },
    { path: '*', redirect: '/' }
]

export default new VueRouter({
    routes // short for `routes: routes`
})
